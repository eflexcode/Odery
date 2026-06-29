package service

import (
	"github.com/cmd/database"
	"github.com/gin-gonic/gin"
)

type Payload struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Intent  string `json:"intent"`
	ImgUrl  string `json:"imgUrl"`
}

type Notification struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type Repo struct{
	db *database.Repo
}

func(r *Repo) GetNotifications(c *gin.Context) {

	
	
}

func (r *Repo)GetNotification(c *gin.Context) {

	
	
}


