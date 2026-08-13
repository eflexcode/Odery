package message

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/cmd/config"
	"github.com/cmd/service"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var exchangeKey = "payment.exchange.key"
var queueName = "payment"
var exchangeName string = "payment.exchange"
var consumer string = "payment.consumer"

func Publish(body []byte, ch *amqp.Channel, cxt context.Context) error {

	return ch.PublishWithContext(cxt,
		exchangeName, // exchange
		exchangeKey,  // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

}

func Consume(ch *amqp.Channel, mongoClient *mongo.Client) {

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
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

	// dbConfig := config.DatabaseConfig{
	// 	ConnUrl:      evn.GetString("DB_URL", "mongodb://localhost/5432"),
	// 	MaxOpenTime:  5,
	// 	MaxIdealConn: 2,
	// 	MaxIdealTime: 5,
	// }

	// mongoClient, err := database.ConnDatabase(dbConfig, context.Background())

	// if err != nil {
	// 	log.Printf("error connecting to mongodb aborting....")
	// 	return
	// }

	// go func() {

	for mg := range msgs {

		var errorCheck bool = false

		// log.Print("RabbitMq Message recived:: " + string(mg.Body))
		var body string = string(mg.Body)
		var splitedBody = strings.Split(body, " rabbitmqIfy ") //  manditory expacteing 2 split. dont use the first one only in notification service to differciate
		if len(splitedBody) != 2 {
			errorCheck = true
			continue
		}

		var orderBody string
		if !errorCheck {
			orderBody = splitedBody[1]
			// log.Print("RabbitMq Message recived:: " + splitedBody[1])
		}

		var order service.Order
		if err := json.Unmarshal([]byte(orderBody), &order); err != nil {
			log.Print("error unwrapping json:: " + err.Error())
			errorCheck = true
			continue
		}

		ordery := mongoClient.Database(config.Dbname).Collection(config.CollectionNameCards)
		orderyPaymentCollection := mongoClient.Database(config.Dbname).Collection(config.CollectionNamePayments)
		log.Println("order:: "+orderBody)
		filter := bson.M{"userid": order.UserId}
		result := ordery.FindOne(context.Background(), filter)
		
		// log.Println(result.Err().Error())
		//check for no card error 404
		// if result.Err().Error() == mongo.ErrNilDocument.Error() {
		// 	log.Print("error getting data from db:: " + result.Err().Error())
		// 	errorCheck = true
		// 	continue
		// }

		var card service.Card
		err := result.Decode(&card)

		if err != nil {
			log.Print("error decoding json:: " + err.Error())
			errorCheck = true
			continue
		}

		var p service.Payment

		// You can fake both network response and card issues here eg.
		// refund,paid:{Type}, insufficient funds, card error, order canceled, network/bank error:{Reason}, done, processing, submitted, failed:{Status}
		// for fields/types
		//	Status    string `json:"status"`
		//	Reason    string `json:"reason"`
		//	Type      string `json:"type"`

		if order.Status == "SUBMITTED" {
			//do something with the else mainly insufficient funds or fake network error.
			if card.Balance > 0 && card.Balance > order.Amount {

				p = service.Payment{
					CardId:      card.Id,
					UserId:      order.UserId,
					Amount:      order.Amount,
					OrderId:     order.Id,
					ProductId:   order.ProductId,
					Status:      "done",
					Type:        "paid",
					Reason:      "-",
					Description: order.Description,
					CreatedAt:   time.Now().String(),
					UpdatedAt:   time.Now().String(),
				}

				update := bson.D{
					{Key: "$set", Value: bson.D{
						{Key: "balance", Value: card.Balance - order.Amount},
					}},
				}

				if !errorCheck {
					_, err := ordery.UpdateOne(context.Background(), filter, update)
					if err != nil {
						log.Println("Payment refund failed")
						errorCheck = true
						continue
					}
				}

			} else {
				p = service.Payment{
					CardId:      card.Id,
					UserId:      order.UserId,
					Amount:      order.Amount,
					OrderId:     order.Id,
					ProductId:   order.ProductId,
					Status:      "failed",
					Type:        "-",
					Reason:      "insufficient funds",
					Description: order.Description,
					CreatedAt:   time.Now().String(),
					UpdatedAt:   time.Now().String(),
				}
			}
		} else if order.Status == "CANCELED" {
			//do refund
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "balance", Value: card.Balance + order.Amount},
				}},
			}

			// id, _ := bson.ObjectIDFromHex(card.Id)
			// _, err := ordery.UpdateByID(context.Background(), id, update)

			if !errorCheck {
				_, err := ordery.UpdateOne(context.Background(), filter, update)
				if err != nil {
					log.Println("Payment refund failed")
					errorCheck = true
					continue
				}
			}

			p = service.Payment{
				CardId:      card.Id,
				UserId:      order.UserId,
				Amount:      order.Amount,
				OrderId:     order.Id,
				ProductId:   order.ProductId,
				Status:      "done",
				Type:        "refund",
				Description: order.Description,
				Reason:      "order canceled",
				CreatedAt:   time.Now().String(),
				UpdatedAt:   time.Now().String(),
			}
		}

		if !errorCheck {

			p.Id = uuid.NewString()
			_, err = orderyPaymentCollection.InsertOne(context.TODO(), p)

			if err != nil {
				log.Print("error inserting payment info")
				errorCheck = true
				continue
			}

		}

		paymentByte, err := json.Marshal(p)
		if err != nil {
			log.Print("error unwrapping bytes")
			continue
		}

		paymentDataForNotificationService := "QueueType:Payment rabbitmqIfy " + string(paymentByte)
		g := bytes.NewBufferString(paymentDataForNotificationService)
		//if error it would just be o bytes
		ch.PublishWithContext(context.Background(),
			exchangeName, // exchange
			exchangeKey,  // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType:  "application/json",
				Body:         g.Bytes(),
				DeliveryMode: amqp.Persistent,
			})
	}

	// }()
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
