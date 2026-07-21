package message

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/cmd/config"
	"github.com/cmd/database"
	"github.com/cmd/evn"
	"github.com/cmd/service"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var key = "order-routing-key"
var queueName = "order"
var exchangeName string = "exchange-order"
var consumer string = "consumer-order"

// func Publish(body any, ch *amqp.Channel, cxt context.Context) error {

// 	return ch.PublishWithContext(cxt,
// 		exchangeName, // exchange
// 		key,          // routing key
// 		false,        // mandatory
// 		false,        // immediate
// 		amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        []byte(body),
// 		})

// }

func Consume(ch *amqp.Channel) {

	msgs, err := ch.Consume(
		queueName, // queue
		consumer,  // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)

	if err != nil {
		log.Panicf("%s", err)
		return
	}

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

	go func() {

		var order service.Order

		for mg := range msgs {

			if err := json.Unmarshal(mg.Body, order); err != nil {
				log.Print("error unwraping json")
				return
			}

			ordery := mongoClient.Database(config.Dbname).Collection(config.CollectionName)
			filter := bson.M{"userId": order.UserId}
			result := ordery.FindOne(context.Background(), filter)

			var card service.Card
			err := result.Decode(&order)

			if err != nil {
				log.Print("error decoding json")
				return
			}
        

			// You can fake both network response and card issues here eg.
			// refund,paid:{Type}, insufficient funds, card error, network error:{Reason}, done, processing, submitted, failed:{Status}
			// for fields/types
			//	Status    string `json:"status"` 
			//	Reason    string `json:"reason"` 
			//	Type      string `json:"type"`

			//do something with the else mainly insufficient funds.
			if card.Balance > 0 && card.Balance > order.Amount {

				p := service.Payment{
					CardId:    card.Id,
					Amount:    order.Amount,
					OrderId:   order.Id,
					ProductId: order.ProductId,
					Status:    "done",
					Type:      "paid",
					Reason:    "-",
					CreatedAt: time.Now().String(),
					UpdatedAt: time.Now().String(),
				}

				_, err := ordery.InsertOne(context.TODO(), p)

				if err != nil {
					log.Print("error inserting payment info")
				}

				paymentByte, err := json.Marshal(p)

				if err != nil {
					log.Print("error unwrapping bytes")
				}

				ch.PublishWithContext(context.Background(),
					exchangeName, // exchange
					key,          // routing key
					false,        // mandatory
					false,        // immediate
					amqp.Publishing{
						ContentType:  "application/json",
						Body:         paymentByte,
						DeliveryMode: amqp.Persistent,
					})

			}else{
				p := service.Payment{
					CardId:    card.Id,
					Amount:    order.Amount,
					OrderId:   order.Id,
					ProductId: order.ProductId,
					Status:    "failed",
					Type:      "-",
					Reason:    "insufficient funds",
					CreatedAt: time.Now().String(),
					UpdatedAt: time.Now().String(),
				}
			} //publish payment failed

		}

	}()
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
