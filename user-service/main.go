package main

import (
	"user-service/app/messagebroker"
	"user-service/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectPostgres()
	config.ConnectRabbitMQ()
	messagebroker.StartConsumer()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "User Service Running",
		})
	})

	r.Run(":8080")
}
