package main

import (
	"time"

	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/cmd/database"
	"github.com/cmd/entity"
	"github.com/cmd/env"
	"github.com/cmd/service"
	"github.com/gin-gonic/gin"

	_"github.com/lib/pq"
)

func main() {

	r := gin.Default()

	db, err := database.ConnectDatabase(env.GetString("DB_URL","postgres://postgres:12345@localhost/oderyUser?sslmode=disable"))

	if err != nil {
		panic("db connection failded")
	}

	database := service.Postgres{
		Postgresql: db,
	}

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, entity.StandardResp{
			Message:    "pong",
			StatusCode: 200,
		})
	})

	r.GET("/login", database.Auth)

	registerWithEureka()
	r.Run(":8081")
}

func registerWithEureka() {

	client := eureka.NewClient([]string{env.GetString("EUREKA_ADDR", "http://localhost:8085/eureka")})

	instance := eureka.NewInstanceInfo(env.GetString("DISCOVERY_ADDR", "localhost:8084"), env.GetString("SERVER_NAME", "Login-server"), env.GetString("IP", "127.0.0.1"), env.GetInt("PORT", 8081), uint(env.GetInt("ttl", 30)), false)

	client.RegisterInstance(env.GetString("SERVER_NAME", "Login-server"), instance)

	go func() {

		for {

			_ = client.SendHeartbeat(instance.App, instance.HostName)

			// if err != nil {
			// 	log.Print("Error: Eureka heartbeat failed " + err.Error())
			// } else {
			// 	log.Print("Info: Eureka heartbeat success")
			// }

			time.Sleep(time.Second * 100)

		}

	}()

}
