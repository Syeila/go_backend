package rabbitmq

import (
	"encoding/json"

	"auth-service/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type UserCreatedEvent struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func PublishUserCreated(
	id int64,
	name string,
	email string,
) error {

	body, _ := json.Marshal(
		UserCreatedEvent{
			UserID: id,
			Name:   name,
			Email:  email,
		},
	)

	return config.RabbitChannel.Publish(
		"",
		"user.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
