package main

import (
	"context"
	"log"

	"github.com/cmd/database"
	"github.com/cmd/service"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	mCli, err := database.ConnectDatabase(context.Background(), "http://localhost:54321")

	if err != nil {
		log.Print("Error connecting to mongodb")
		return
	}
	
	log.Print("MongoDb connection established")

	dbRepo := database.Repo{
		DatabaseMongo: mCli,
	}
	
	s := service.Repo{
		Db: &dbRepo,
	}

	r.GET("/get-notifications/:userId", s.GetNotifications)
	r.GET("/get/:id",s.GetNotification)

}
