package service

import (
	"context"

	"github.com/cmd/database"
	"github.com/env"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)


type Repo struct {
	Db *database.Repo
}

func (r *Repo) GetNotifications(c *gin.Context) {

	var userId = c.Param("userId")
	col := r.Db.DatabaseMongo.Database(env.GetString("DATABASE_NAME", "OrderyNotifications")).Collection(env.GetString("COLLECTION_NAME", "Notifications"))

	

}

func (r *Repo) GetNotification(c *gin.Context) {

}
