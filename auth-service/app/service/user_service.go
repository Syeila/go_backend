package service

import (
	"auth-service/app/domain/dao"
	"auth-service/app/repository"
)

func GetProfile(userID int) (dao.User, error) {

	return repository.FindUserByID(userID)
}
