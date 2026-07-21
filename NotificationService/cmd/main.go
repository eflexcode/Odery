package main

import (
	"context"
	"log"

	"github.com/cmd/database"
	"github.com/cmd/service"
	"github.com/env"
	"github.com/gin-gonic/gin"
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

	messages, err := chann.Consume(
		"order", // queue
		"",      // consumer
		true,    // durability
		false,   // delete when unused
		false,   // exclusive
		false,
		nil)

	if err != nil {
		log.Print("Rabbitmq Consume error")
		return
	}
	
	log.Print("Rabbitmq Connection established")

	go func() {
		for m := range messages {
			log.Println("llll"+string(m.Body))
		}
	}()

	dbRepo := database.Repo{
		DatabaseMongo: mCli,
	}

	s := service.Repo{
		Db: &dbRepo,
	}

	r.GET("/get-notifications/:userId", s.GetNotifications)
	r.GET("/get/:id", s.GetNotification)

	r.Run(env.GetString("PORT", ":8084"))
}
