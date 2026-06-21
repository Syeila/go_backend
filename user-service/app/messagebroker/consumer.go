package messagebroker

import (
	"encoding/json"
	"log"

	"user-service/app/domain/dto"
	"user-service/app/service"
	"user-service/config"
)

func StartConsumer() {

	msgs, err := config.RabbitChannel.Consume(
		"user.created",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	go func() {

		for msg := range msgs {

			var event dto.UserCreatedEvent

			err := json.Unmarshal(
				msg.Body,
				&event,
			)

			if err != nil {
				log.Print(err)
				continue
			}

			err = service.CreateUser(
				event.UserID,
				event.Name,
				event.Email,
			)

			if err != nil {
				log.Println(err)
				continue
			}

			log.Println(
				"user synced:",
				event.UserID,
			)
		}
	}()
}
