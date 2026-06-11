package message

import (
	"context"
	"encoding/json"
	"log"

	"github.com/cmd/config"
	"github.com/cmd/database"
	"github.com/cmd/service"
	"go.mongodb.org/mongo-driver/v2/bson"
	amqp "github.com/rabbitmq/amqp091-go"
)

var key = "order-routing-key"
var queueName = "order"
var exchangeName string = "exchange-order"
var consumer string = "consumer-order"

func Publish(body any, ch *amqp.Channel, cxt context.Context) error {

	return ch.PublishWithContext(cxt,
		exchangeName, // exchange
		key,          // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

}

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
				log.Print("uu")
				return
			}

			ordery := mongoClient.Database(config.Dbname).Collection(config.CollectionName)
			filter := bson.M{"cardId": order.CardId}
			result := ordery.FindOne(context.Background(), filter)
			
			var card service.Card
			err := result.Decode(&order)
						
			if err != nil{
				log.Print("error decoding json")
				return
			}
			
			if card.Balance > 0  && card.Balance > {
				
				
				
			}
			
		}

	}()
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

