package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cmd/config"
	"github.com/cmd/database"
	"github.com/cmd/evn"
	"github.com/cmd/message"
	"github.com/cmd/service"
	"github.com/gin-gonic/gin"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	if err := evn.LoadEvn(); err != nil {
		log.Printf("error loading evn file aborting....")
		return
	}

	// var cxt = context.Background()

	dbConfig := config.DatabaseConfig{
		ConnUrl:      evn.GetString("DB_URL", "mongodb://localhost/27017/OrderyPayment"),
		MaxOpenTime:  5,
		MaxIdealConn: 2,
		MaxIdealTime: 5,
	}

	mongoClient, err := database.ConnDatabase(dbConfig, context.Background())

	if err != nil {
		log.Printf("error connecting to mongodb aborting....")
		return
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		message.FailOnError(err, "Failed to connect to RabbitMQ")
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		message.FailOnError(err, "Failed to open a channel")
		return
	}

	defer ch.Close()

	qu, err := ch.QueueDeclare(
		"payment", // name
		true,      // durability
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)

	if err != nil {
		message.FailOnError(err, "Failed to declare a RabbitMq queue")
		return
	}

	var paymentExchangeName string = "payment.exchange"
	// var paymentExchangeKey string = "payment.exchange.key" //not used in fanout tho
	// var orderQueueName string = "order"
	var orderExchangeName string = "exchange-order"
	var orderExchangeKey string = "order-routing-key"

	err = ch.ExchangeDeclare(paymentExchangeName, "fanout", true, false, false, false, nil)
	if err != nil {
		message.FailOnError(err, "Declare exchange  failed")
		return
	}

	// err = ch.QueueBind(qu.Name, paymentExchangeKey, paymentExchangeName, false, nil) //payement exchange
	// if err != nil {
	// 	message.FailOnError(err, "Failed to bind to payment queue ")
	// } //dont want to see what i published

	err = ch.QueueBind(qu.Name, orderExchangeKey, orderExchangeName, false, nil)
	if err != nil {
		message.FailOnError(err, "Failed to bind to payment queue ")
	}

	message.Consume(ch, mongoClient)
	// msgs, err := ch.Consume(
	// 	q.Name, // queue
	// 	"",     // consumer
	// 	true,   // auto-ack
	// 	false,  // exclusive
	// 	false,  // no-local
	// 	false,  // no-wait
	// 	nil,    // args
	// )

	// go func() {
	// 	for d := range msgs {
	// 		log.Printf("Received a message: %s", d.Body)
	// 	}
	// }()

	dbRep := database.DatabaseRep{
		Mongo: mongoClient,
	}

	s := service.Repo{
		Database: &dbRep,
		MqChan:   ch,
	}

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/add-card", s.AddCard)
	r.PUT("/update-card", s.UpdateCard)
	r.DELETE("/delete-card/{user_id}", s.DeleteCard)
	r.GET("/info/{id}", s.GetCardInfo)
	r.POST("/process-payment", s.MakePayment)
	r.POST("/request-refund", s.RequestRefund)
	r.GET("/get-payment-slipts/:user_id", s.GetPaymentSlipts)
	r.GET("/get-payment-slipt/:id", s.GetPaymentSlipt)

	r.Run(evn.GetString("PORT", ":8089"))

}
