package service

import (
	"context"
	"evn"
	"net/http"

	"github.com/cmd/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Card struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	Pan       string `json:"pan"`
	Cvv       int    `json:"cvv"`
	Exp       string `json:"exp"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type CardUpdate struct {
	UserId    string `json:"userId"`
	Pan       string `json:"pan"`
	Cvv       int    `json:"cvv"`
	Exp       string `json:"exp"`
}

type StandardResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Repo struct {
	Database *database.DatabaseRep
}

var dbname = evn.GetString("DATABASE_NAME")
var collectionName = evn.GetString("COLLECTION_NAME")

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
	_, err := ordery.InsertOne(context.TODO(), card)

	if err != nil {
		s := StandardResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, s)
		return
	}

	s := StandardResponse{
		Message: "Card added succefully",
		Status:  http.StatusOK,
	}

	c.JSON(http.StatusOK, s)
}

func (r *Repo) UpdateCard(c *gin.Context) {
	
	ordery := r.Database.Mongo.Database(dbname).Collection(collectionName)
	
}

func (r *Repo) DeleteCard(c *gin.Context) {

	ordery := r.Database.Mongo.Database(dbname).Collection(collectionName)
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

	ordery := r.Database.Mongo.Database(dbname).Collection(collectionName)
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
