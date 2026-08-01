package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/cmd/config"
	"github.com/cmd/database"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Card struct {
	Id        string  `json:"id"`
	UserId    string  `json:"userId"`
	Pan       string  `json:"pan"`
	Cvv       int     `json:"cvv"`
	Balance   float64 `json:"balance"`
	Exp       string  `json:"exp"`
	Active    bool    `json:"active"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type CardUpdate struct {
	Id      string  `json:"id"`
	UserId  string  `json:"userId"`
	Pan     string  `json:"pan"`
	Cvv     int     `json:"cvv"`
	Exp     string  `json:"exp"`
	Balance float64 `json:"balance"`
}

type StandardResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Payment struct {
	Id        string  `json:"id"`
	UserId    string  `json:"user_id"`
	CardId    string  `json:"card_id"`
	Amount    float64 `json:"amount"`
	ProductId string  `json:"product_id"`
	OrderId   string  `json:"order_id"`
	Status    string  `json:"status"` //done, processing, submitted, failed
	Reason    string  `json:"reason"` //eg: insufficient funds, card error, network error,order canceled
	Type      string  `json:"type"`   //refund,paid,-
	// Description string  `json:"description"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Order struct {
	Id          string  `json:"id"`
	UserId      string  `json:"user_id"`
	Count       int     `json:"count"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	ProductId   string  `json:"product_id"`
	Status      string  `json:"status"` //done, processing, submitted, canceled
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type Refund struct {
	OrderId   string `json:"order_id"`
	PaymentId string `json:"payment_id"`
	UserId    string `json:"user_id"`
}

type Repo struct {
	Database *database.DatabaseRep
	MqChan   *amqp091.Channel
}

var exchangeKey = "payment.exchange.key"
var queueName = "payment"
var exchangeName string = "payment.exchange"
var consumer string = "payment.consumer"

// var dbname = evn.GetString("DATABASE_NAME")
// var collectionName = evn.GetString("COLLECTION_NAME")

func (r *Repo) AddCard(c *gin.Context) {

	var card Card

	if err := c.ShouldBindJSON(&card); err != nil {

		s := StandardResponse{
			Message: "Bad request",
			Status:  http.StatusBadRequest,
		}

		c.JSON(http.StatusBadRequest, s)
		return
	}

	ordery := r.Database.Mongo.Database("OrderyDatabase").Collection("cards")

	queryFilter := bson.M{"userId": card.UserId}
	totalCount, err := ordery.CountDocuments(c.Copy(), queryFilter)

	if err != nil {
		s := StandardResponse{
			Message: "Internal server Error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusBadRequest, s)
		return
	}

	if totalCount > 0 {
		s := StandardResponse{
			Message: "Only one card can be added",
			Status:  http.StatusOK,
		}

		c.JSON(http.StatusBadRequest, s)
		return
	}
	_, err = ordery.InsertOne(context.TODO(), card)

	if err != nil {

		s := StandardResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	s := StandardResponse{
		Message: "Card added successfully",
		Status:  http.StatusOK,
	}

	c.JSON(http.StatusOK, s)
}

func (r *Repo) UpdateCard(c *gin.Context) {

	var cardUp CardUpdate

	if err := c.ShouldBindJSON(cardUp); err != nil {

		s := StandardResponse{
			Message: "Bad request",
			Status:  http.StatusBadRequest,
		}

		c.JSON(http.StatusBadRequest, s)
		return
	}

	ordery := r.Database.Mongo.Database(config.Dbname).Collection(config.CollectionNameCards)

	// filter := bson.M{"userId": cardUp.UserId}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "pan", Value: cardUp.Pan},
			{Key: "exp", Value: cardUp.Exp},
			{Key: "cvv", Value: cardUp.Cvv},
			{Key: "balance", Value: cardUp.Balance},
		}},
	}

	id, _ := bson.ObjectIDFromHex(cardUp.Id)

	_, err := ordery.UpdateByID(c.Copy(), id, update)

	if err != nil {

		s := StandardResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	s := StandardResponse{
		Message: "Card updated sucefuly",
		Status:  http.StatusOK,
	}

	c.JSON(http.StatusOK, s)

}

func (r *Repo) DeleteCard(c *gin.Context) {

	ordery := r.Database.Mongo.Database(config.Dbname).Collection(config.CollectionNameCards)
	userId := c.Param("userId")
	filter := bson.M{"userId": userId}

	_, err := ordery.DeleteOne(c.Copy(), filter)

	if err != nil {

		s := StandardResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	s := StandardResponse{
		Message: "Card deleted successfully",
		Status:  http.StatusOK,
	}

	c.JSON(http.StatusOK, s)

}

func (r *Repo) GetCardInfo(c *gin.Context) {

	ordery := r.Database.Mongo.Database(config.Dbname).Collection(config.CollectionNameCards)
	userId := c.Param("userId")
	filter := bson.M{"userId": userId}

	var card Card

	result := ordery.FindOne(c.Copy(), filter)

	if err := result.Decode(&card); err != nil {

		s := StandardResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	c.JSON(http.StatusInternalServerError, result)

}

func (r *Repo) MakePayment(c *gin.Context) {

	var order Order

	err := c.ShouldBindBodyWithJSON(order)

	if err != nil {
		s := StandardResponse{
			Message: "Bad request",
			Status:  http.StatusBadRequest,
		}

		c.JSON(http.StatusBadRequest, s)
		return
	}

	ordery := r.Database.Mongo.Database(config.Dbname).Collection(config.CollectionNameCards)
	filter := bson.M{"userId": order.UserId}
	result := ordery.FindOne(context.Background(), filter)

	var card Card
	err = result.Decode(&card)

	if err != nil {
		log.Print("error decoding bson")
		s := StandardResponse{
			Message: "Internal server error",
			Status:  500,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	if card.Balance > 0 && card.Balance > order.Amount {

		p := Payment{
			CardId:    card.Id,
			Amount:    order.Amount,
			UserId:    order.UserId,
			OrderId:   order.Id,
			ProductId: order.ProductId,
			Status:    "done",
			Type:      "paid",
			Reason:    "",
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		}

		_, err := ordery.InsertOne(context.TODO(), p)

		if err != nil {
			log.Print("error inserting payment info")
			return
		} //publish on rmQueue
		b, err := json.Marshal(p)
		if err != nil {
			log.Println("failed to wrap struct to json for mq")
		}
		
		r.MqChan.PublishWithContext(c.Copy(),
			exchangeName, // exchange
			exchangeKey,  // routing key
			false,        // mandatory
			false,        // immediate
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        b,
			})
		c.JSON(http.StatusAccepted, p)
	}

}

func (r *Repo) RequestRefund(c *gin.Context) {
	var refund Refund
	if err := c.ShouldBindBodyWithJSON(refund); err != nil {
		s := StandardResponse{
			Message: "invalid json sent",
			Status:  http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, s)
		return
	}

}
