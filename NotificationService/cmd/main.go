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

	queue, err := chann.QueueDeclare("order-notification", true, false, false, false, nil)

	if err != nil {
		log.Print("Failed to declare notification queue")
		return
	}

	err = chann.ExchangeDeclare("order-notification.exchange", "direct", true, false, false, false, nil)

	if err != nil {
		log.Print("Failed to declare exchange queue")
		return
	}

	err = chann.QueueBind(queue.Name, "order-routing-key", "exchange-order", false, nil)
	err = chann.QueueBind(queue.Name, "order-notification.exchange-key", "order-notification.exchange", false, nil)

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

			if t[1] == "Order" { //two type  sofar Order and Payment

				var order entity.Order
				//static queue type has to be
				if err := json.Unmarshal([]byte(payloadAndType[1]), &order); err != nil {
					log.Println("Error parsing rabbitmq json to struct")
					return
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
					payload.Message =  count + " Purchase for " + order.ProductName + " at " + order.ProductCurrency + " " + amount + " with order id: " + order.Id + " canceld"

				}

				payload.Intent = "http://localhost:8092/" + order.Id

				notification.Payload = payload
			} else if t[1] == "Payment" {

			}

			database.InsertNotification(context.Background(), dbRepo.DatabaseMongo, notification)

		}
	}()

	r.GET("/get-notifications/:userId", s.GetNotifications)
	r.GET("/get/:id", s.GetNotification)

	r.Run(env.GetString("PORT", ":8084"))
}
