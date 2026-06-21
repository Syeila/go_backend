package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

func ConnectRabbitMQ() {

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal("RabbitMQ connection error:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	RabbitConn = conn
	RabbitChannel = ch

	log.Println("RabbitMQ Connected (user-service)")
}
