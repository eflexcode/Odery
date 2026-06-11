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

	var cxt = context.Background()

	dbConfig := config.DatabaseConfig{
		ConnUrl:      evn.GetString("DB_URL", "mongodb://localhost/5432"),
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
	message.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	message.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durability
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	dbRep := database.DatabaseRep{
		Mongo: mongoClient,
	}

	s := service.Repo{
		Database: &dbRep,
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

	r.Run(evn.GetString("PORT", ":8089"))

}
