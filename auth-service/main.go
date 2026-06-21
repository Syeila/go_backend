package main

import (
	"auth-service/app/router"
	"auth-service/config"

	"github.com/gin-gonic/gin"
)

// @title Auth Service API
// @version 1.0
// @description Auth Microservice API
// @host localhost:8081
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {

	config.LoadEnv()

	config.ConnectMySQL()

	config.ConnectRedis()

	config.ConnectRabbitMQ()

	r := gin.Default()

	router.SetupRouter(r)

	r.Run(":8080")
}
