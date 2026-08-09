package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cmd/config"
	"github.com/cmd/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Card struct {
	Id        string  `bson:"_id"`
	UserId    string  `json:"userid"`
	Pan       string  `json:"pan"`
	Cvv       string  `json:"cvv"`
	Balance   float64 `json:"balance"`
	Exp       string  `json:"exp"`
	Active    bool    `json:"active"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type CardAdd struct {
	UserId  string  `json:"userid"`
	Pan     string  `json:"pan"`
	Cvv     string  `json:"cvv"`
	Balance float64 `json:"balance"`
	Exp     string  `json:"exp"`
	Active  bool    `json:"active"`
}

type CardUpdate struct {
	UserId  string  `json:"userId"`
	Pan     string  `json:"pan"`
	Cvv     string     `json:"cvv"`
	Exp     string  `json:"exp"`
	Balance float64 `json:"balance"`
}

type StandardResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Payment struct {
	Id          string  `json:"id"`
	UserId      string  `json:"user_id"`
	CardId      string  `json:"card_id"`
	Amount      float64 `json:"amount"`
	ProductId   string  `json:"product_id"`
	OrderId     string  `json:"order_id"`
	Status      string  `json:"status"` //done, processing, submitted, failed
	Reason      string  `json:"reason"` //eg: insufficient funds, card error, network error,order canceled
	Type        string  `json:"type"`   //refund,paid,-
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type PaymentResult struct {
	Payment []Payment `json:"payments"`
	Total   int64     `json:"total"`
	Page    int64     `json:"page"`
	Limit   int64     `json:"limit"`
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

type User struct {
	Id        string `json:"id"`
	UserName  string `json:"username"`
	Role      string `json:"role"`
	Name      string `json:"Name"`
	ImgUrl    string `json:"img_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
var orderServerBaseUrl string = "http://localhost:8091/"
var userServerBaseUrl string = "http://localhost:8095/"

// var dbname = evn.GetString("DATABASE_NAME")
// var collectionName = evn.GetString("COLLECTION_NAME")

func (r *Repo) AddCard(c *gin.Context) {

	var card CardAdd

	if err := c.ShouldBindJSON(&card); err != nil {

		s := StandardResponse{
			Message: "Bad request",
			Status:  http.StatusBadRequest,
		}

		c.JSON(http.StatusBadRequest, s)
		return
	}

	ordery := r.Database.Mongo.Database("OrderyPayment").Collection("cards")

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

	resp, err := http.Get(userServerBaseUrl + card.UserId)
	if err != nil {

		s := StandardResponse{
			Message: "Invalid user id sent",
			Status:  http.StatusBadRequest,
		}

		c.JSON(http.StatusBadRequest, s)
		return
	}

	defer resp.Body.Close()

	bodyByte, err := io.ReadAll(resp.Body)

	if err != nil {

		s := StandardResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	var user User

	err = json.Unmarshal(bodyByte, &user)

	if err != nil {
		s := StandardResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	if user.Id != card.UserId {
		s := StandardResponse{
			Message: "Invalid user id sent",
			Status:  http.StatusUnauthorized,
		}

		c.JSON(http.StatusUnauthorized, s)
		return
	}

	var cardR Card
	cardR.Id = uuid.New().String()
	cardR.Active = true
	cardR.Balance = card.Balance
	cardR.Cvv = card.Cvv
	cardR.Pan = card.Pan
	cardR.Exp = card.Exp
	cardR.UserId = card.UserId
	cardR.UpdatedAt = time.Now().String()
	cardR.CreatedAt = time.Now().String()

	_, err = ordery.InsertOne(context.TODO(), cardR)

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

	if err := c.ShouldBindJSON(&cardUp); err != nil {

		s := StandardResponse{
			Message: "Bad request",
			Status:  http.StatusBadRequest,
		}
		log.Println("uoououououououppipipipipipi")
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
	filter := bson.M{"userId": cardUp.UserId}
	// id, _ := bson.ObjectIDFromHex(cardUp.Id)
	_, err := ordery.UpdateOne(c.Copy(), filter, update)
	// _, err := ordery.UpdateByID(c.Copy(), id, update)

	if err != nil {

		s := StandardResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	s := StandardResponse{
		Message: "Card updated successfully",
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

	var p Payment

	if card.Balance > 0 && card.Balance > order.Amount {

		var deducte float64 = card.Balance - order.Amount
		va := bson.D{{Key: "$set", Value: bson.D{
			{Key: "balance", Value: deducte},
		}}}

		id, _ := bson.ObjectIDFromHex(card.Id)
		_, err := ordery.UpdateByID(c.Copy(), id, va)

		if err != nil {
			p = Payment{
				CardId:      card.Id,
				Amount:      order.Amount,
				UserId:      order.UserId,
				OrderId:     order.Id,
				ProductId:   order.ProductId,
				Status:      "failed",
				Type:        "_",
				Reason:      "insufficient funds",
				Description: order.Description,
				CreatedAt:   time.Now().String(),
				UpdatedAt:   time.Now().String(),
			}
		} else {
			p = Payment{
				CardId:      card.Id,
				Amount:      order.Amount,
				UserId:      order.UserId,
				OrderId:     order.Id,
				ProductId:   order.ProductId,
				Status:      "done",
				Type:        "paid",
				Reason:      "",
				Description: order.Description,
				CreatedAt:   time.Now().String(),
				UpdatedAt:   time.Now().String(),
			}
		}
	} else {
		p = Payment{
			CardId:      card.Id,
			Amount:      order.Amount,
			UserId:      order.UserId,
			OrderId:     order.Id,
			ProductId:   order.ProductId,
			Status:      "failed",
			Type:        "_",
			Reason:      "insufficient funds",
			Description: order.Description,
			CreatedAt:   time.Now().String(),
			UpdatedAt:   time.Now().String(),
		}

	}
	orderyP := r.Database.Mongo.Database(config.Dbname).Collection(config.CollectionNamePayments)

	_, err = orderyP.InsertOne(context.TODO(), p)

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

func (r *Repo) GetPaymentSlipts(c *gin.Context) {

	var userId string = c.Param("user_id")

	var p = c.Query("page")
	var l = c.Query("limit")

	page, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, StandardResponse{Message: "NAN Page", Status: http.StatusBadRequest})
		return
	}
	limit, err := strconv.Atoi(l)
	if err != nil {
		c.JSON(http.StatusBadRequest, StandardResponse{Message: "NAN Limit", Status: http.StatusBadRequest})
		return
	}

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 1
	}

	offset := (page - 1) * limit
	ordery := r.Database.Mongo.Database(config.Dbname).Collection(config.CollectionNamePayments)
	filter := bson.M{"user_id": userId}
	ops := bson.D{{Key: "created_at", Value: -1}}

	total, err := ordery.CountDocuments(c.Copy(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StandardResponse{Message: "Internal server error", Status: http.StatusInternalServerError})
		return
	}
	option := options.Find().SetSort(ops).SetSkip(int64(offset)).SetLimit(int64(limit))
	results, err := ordery.Find(context.Background(), filter, option)

	if err != nil {
		c.JSON(http.StatusInternalServerError, StandardResponse{Message: "Internal server error", Status: http.StatusInternalServerError})
		return
	}

	defer results.Close(c.Copy())

	var payments []Payment
	err = results.All(c.Copy(), &payments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StandardResponse{Message: "Internal server error", Status: http.StatusInternalServerError})
		return
	}

	paymentResult := PaymentResult{
		Payment: payments,
		Total:   total,
		Limit:   int64(limit),
		Page:    int64(page),
	}

	c.JSON(http.StatusOK, paymentResult)

}

func (r *Repo) GetPaymentSlipt(c *gin.Context) {
	var id string = c.Param("id")

	ordery := r.Database.Mongo.Database(config.Dbname).Collection(config.CollectionNamePayments)
	filter := bson.M{"id": id}
	result := ordery.FindOne(context.Background(), filter)

	var p Payment
	err := result.Decode(&p)
	if err != nil {
		s := StandardResponse{
			Message: "internal server error",
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, s)
		return
	}
	c.JSON(http.StatusOK, p)
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

	//make api call to order server to verify order id
	//make db call for payment id
	//verify user id in both payment and order or make api request

	resp, err := http.Get(orderServerBaseUrl + "getById/" + refund.OrderId)

	if err != nil {
		s := StandardResponse{
			Message: "invalid order id sent",
			Status:  http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, s)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s := StandardResponse{
			Message: "invalid order id sent",
			Status:  http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, s)
		return
	}

	var order Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		s := StandardResponse{
			Message: "invalid order id sent",
			Status:  http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, s)
		return
	}

	coll := r.Database.Mongo.Database(config.Dbname).Collection(config.CollectionNamePayments)
	filter := bson.M{"id": refund.PaymentId}
	sResult := coll.FindOne(c.Copy(), filter)

	var payment Payment
	err = sResult.Decode(&payment)

	if err != nil {
		s := StandardResponse{
			Message: "invalid payment id sent",
			Status:  http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, s)
		return
	}

	if payment.UserId != order.UserId {
		s := StandardResponse{
			Message: "invalid user id sent",
			Status:  http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, s)
		return
	}

	if payment.UserId != refund.UserId {
		s := StandardResponse{
			Message: "invalid user id sent",
			Status:  http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, s)
		return
	}

	if payment.Status == "refund" {
		s := StandardResponse{
			Message: "already refunded",
			Status:  http.StatusOK,
		}
		c.JSON(http.StatusOK, s)
		return
	}

	if payment.OrderId != order.Id {
		s := StandardResponse{
			Message: "Order id in payment does not match",
			Status:  http.StatusOK,
		}
		c.JSON(http.StatusOK, s)
		return
	}

	cardColl := r.Database.Mongo.Database(config.Dbname).Collection(config.CollectionNameCards)
	cardFilter := bson.M{"userId": refund.UserId}
	sRe := cardColl.FindOne(c.Copy(), cardFilter)

	var card Card
	err = sRe.Decode(&card)

	if err != nil {
		s := StandardResponse{
			Message: "internal server error",
			Status:  http.StatusInternalServerError,
		}
		c.JSON(http.StatusInternalServerError, s)
		return
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "balance", Value: card.Balance + payment.Amount},
		}},
	}

	id, _ := bson.ObjectIDFromHex(card.Id)

	_, err = cardColl.UpdateByID(c.Copy(), id, update)

	if err != nil {

		s := StandardResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	//update payment
	updateP := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: "done"},
			{Key: "reason", Value: "order canceled"},
		}},
	}

	idP, _ := bson.ObjectIDFromHex(payment.Id)

	_, err = coll.UpdateByID(c.Copy(), idP, updateP)

	if err != nil {

		s := StandardResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	payment.Status = "done"
	payment.Reason = "order canceled"

	pBytes, err := json.Marshal(payment)

	r.MqChan.PublishWithContext(context.Background(),
		exchangeName, // exchange
		exchangeKey,  // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         pBytes,
			DeliveryMode: amqp.Persistent,
		})

	s := StandardResponse{
		Message: "refund made successfully",
		Status:  http.StatusOK,
	}
	c.JSON(http.StatusOK, s)

}
