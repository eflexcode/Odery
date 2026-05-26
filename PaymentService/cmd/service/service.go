package service

import (
	
	"github.com/gin-gonic/gin"
)

type card struct{
	Id string `json:"id"`
	Pan string `json:"pan"`
	Cvv int `json:"cvv"`
	Exp string `json:"exp"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func AddCard (c *gin.Context)  {
	
}

func UpdateCard(c *gin.Context){
	
}

func DeleteCard(c *gin.Context){
	
}

func GetCardInfo(c *gin.Context){
	
}



