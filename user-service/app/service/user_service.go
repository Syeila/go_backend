package service

import (
	"user-service/app/repository"
)

func CreateUser(
	id int64,
	name string,
	email string,
) error {
	return repository.CreateUser(
		id,
		name,
		email,
	)
}
