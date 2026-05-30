package service

import (
	"context"
	"net/http"

	"github.com/cmd/database"
	"github.com/gin-gonic/gin"
)

type Card struct {
	Id        string `json:"id"`
	Pan       string `json:"pan"`
	Cvv       int    `json:"cvv"`
	Exp       string `json:"exp"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type StandardResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Repo struct {
	Database *database.DatabaseRep
}

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
	
	if err != nil{
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

}

func (r *Repo) DeleteCard(c *gin.Context) {

}

func (r *Repo) GetCardInfo(c *gin.Context) {

}
