package main

import (
	"log"
	"net/http"

	"github.com/cmd/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func  main(){
	
 	if err:=godotenv.Load(); err != nil{
		log.Printf("error loading evn file aborting....")
		return
  	}

	r:= gin.Default()
			
	r.GET("/ping",func (ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})
	} )
	
	r.POST("/add-card",service.AddCard)
	r.PUT("/update-card",service.UpdateCard)
	r.DELETE("/delete-card")
	
	r.Run(":8089")
	
}