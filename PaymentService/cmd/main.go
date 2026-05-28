package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cmd/config"
	"github.com/cmd/database"
	"github.com/cmd/evn"
	"github.com/cmd/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("error loading evn file aborting....")
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
	dbRep := database.DatabaseRep{
		mongoC
	}
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/add-card", service.AddCard)
	r.PUT("/update-card", service.UpdateCard)
	r.DELETE("/delete-card", service.DeleteCard)
	r.GET("/info/{id}", service.GetCardInfo)

	r.Run(evn.GetString("PORT", ":8089"))

}
