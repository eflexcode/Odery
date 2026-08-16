package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/cmd/database"
	"github.com/cmd/entity"
	"github.com/cmd/service"
	"github.com/env"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
)

func main() {

	r := gin.Default()

	mCli, err := database.ConnectDatabase(context.Background(), env.GetString("DB_URL", "mongodb://localhost/27017"))

	if err != nil {
		log.Print("Error connecting to mongodb")
		return
	}

	log.Print("MongoDb connection established")

	// conn, err := ampq.Dial("amqp://guest:guest@localhost:5672/")
	conn, err := amqp091.Dial(env.GetString("RABBITMQ", "amqp://guest:guest@localhost:5672/"))

	if err != nil {
		log.Print("Error connecting to rabbitmq")
		return
	}

	defer conn.Close()

	chann, err := conn.Channel()

	if err != nil {
		log.Print("Error connecting to rabbitmq channel")
		return
	}

	defer chann.Close()

	queue, err := chann.QueueDeclare("order-notification", true, false, false, false, nil) //notification queue

	if err != nil {
		log.Print("Failed to declare notification queue")
		return
	}

	err = chann.ExchangeDeclare("order-notification.exchange", "fanout", true, false, false, false, nil)

	if err != nil {
		log.Print("Failed to declare exchange queue error::  " + err.Error())
		return
	}

	err = chann.ExchangeDeclare("payment.exchange", "fanout", true, false, false, false, nil)

	if err != nil {
		log.Print("Failed to declare exchange queue error::  " + err.Error())
		return
	}

	err = chann.QueueBind(queue.Name, "order-routing-key", "exchange-order", false, nil)
	err = chann.QueueBind(queue.Name, "", "payment.exchange", false, nil)
	// err = chann.QueueBind(queue.Name, "order-notification.exchange-key", "order-notification.exchange", false, nil) //dont want to see my own queue onles testing

	if err != nil {
		log.Print("Failed to bind exchange with  queue")
		return
	}

	messages, err := chann.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // durability
		false,      // delete when unused
		false,      // exclusive
		false,
		nil)

	if err != nil {
		log.Print("Rabbitmq Consume error")
		return
	}

	log.Print("Rabbitmq Connection established")

	dbRepo := database.Repo{
		DatabaseMongo: mCli,
	}

	s := service.Repo{
		Db: &dbRepo,
	}

	go func() {
		for m := range messages {
			log.Println("RabbitMq received payload:: ")

			var mqBody string = string(m.Body)
			var payloadAndType = strings.Split(mqBody, " rabbitmqIfy ")

			var typeQ = payloadAndType[0]
			var t = strings.Split(typeQ, ":")

			var notification entity.Notification
			notification.Id = uuid.New().String()
			log.Println("RabbitMq received payload:: " + payloadAndType[0])

			if t[1] == "Order" { //two type  sofar Order and Payment

				var order entity.Order
				//static queue type has to be
				if err := json.Unmarshal([]byte(payloadAndType[1]), &order); err != nil {
					log.Println("Error parsing rabbitmq json to struct")
					continue
				}

				notification.UserId = order.UserId
				notification.CreatedAt = time.Now().String()
				notification.Type = "order"

				// var currencySymbol string

				// if order.ProductCurrency == "NGN" {
				// 	currencySymbol = "₦"
				// } // do the rest letter or leave it like this
				//
				//
				// acFormatter :=	accounting.Accounting{
				// 	Symbol	:order.ProductCurrency,
				// 	Precision: 2,
				// 	Thousand: ",",
				// 	Decimal: ".",
				// 	}

				// formattedAmount :=	acFormatter.FormatMoney(order.Amount)
				var count string = strconv.Itoa(int(order.Count))
				var amount string = strconv.Itoa(int(order.Amount))

				var payload entity.Payload
				payload.ImgUrl = order.ProductImgUrl
				payload.Title = "Order placed!!"
				payload.Message = "Purchase for " + count + " " + order.ProductName + " at " + order.ProductCurrency + " " + amount + " with order id: " + order.Id + " placed"

				if order.Status == "CANCELED" {
					payload.Title = "Order canceld"
					payload.Message = count + " Purchase for " + order.ProductName + " at " + order.ProductCurrency + " " + amount + " with order id: " + order.Id + " canceld"
				}

				payload.Intent = "http://localhost:8092/" + order.Id

				notification.Payload = payload

				database.InsertNotification(context.Background(), dbRepo.DatabaseMongo, notification)

			} else if t[1] == "Payment" {
				var payment entity.Payment

				if err := json.Unmarshal([]byte(payloadAndType[1]), &payment); err != nil {
					log.Println("Error parsing rabbitmq json to struct")
					continue
				}

				notification.UserId = payment.UserId
				notification.CreatedAt = time.Now().String()
				notification.Type = "payment"

				var payload entity.Payload
				if payment.Status == "done" {
					payload.Title = "Payment made!"
				} else if payment.Status == "failed" {
					payload.Title = "Payment failed!"
				} else {
					payload.Title = "Issue with payment"
				}

				payload.Message = payment.Description + " id's payment id: " + payment.Id + "\n order id: " + payment.OrderId + "\n product id: " + payment.ProductId + "\n"
				payload.Intent = "http://localhost:8089/get-payment-slipt/" + payment.Id

				notification.Payload = payload
				if payment.CardId != "" {
					database.InsertNotification(context.Background(), dbRepo.DatabaseMongo, notification)
				}
			}

		}
	}()

	r.GET("/get-notifications/:userId", s.GetNotifications)
	r.GET("/get/:id", s.GetNotification)

	r.Run(env.GetString("PORT", ":8084"))
}
