package service

import (
	"auth-service/app/domain/dao"
	"auth-service/app/domain/dto"
	rabbitmq "auth-service/app/messagebroker"
	"auth-service/app/repository"
	"errors"
	"math"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Register(req dto.RegisterRequest) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)

	// 1. simpan ke mysql
	userId, err := repository.CreatedUser(req)
	if err != nil {
		return err
	}

	// 2. publish event ke rabbitmq
	err = rabbitmq.PublishUserCreated(
		userId,
		req.Name,
		req.Email,
	)

	if err != nil {
		return err
	}

	return nil
}

func Login(
	req dto.LoginRequest,
	ipAddress string,
	userAgent string,
) (map[string]string, error) {

	user, err := repository.FindUserByEmail(req.Email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := GenerateToken(user.ID)

	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateRefreshToken(user.ID)

	if err != nil {
		return nil, err
	}

	err = repository.SaveRefreshToken(
		user.ID,
		refreshToken,
		req.DeviceName,
		ipAddress,
		userAgent,
		time.Now().Add(time.Hour*24*7),
	)

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func RefreshToken(
	req dto.RefreshTokenRequest,
) (map[string]string, error) {

	err := repository.FindRefreshToken(
		req.RefreshToken,
	)

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	claims, err := ParseToken(
		req.RefreshToken,
	)

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	userID := int(claims["user_id"].(float64))

	err = repository.DeleteRefreshToken(
		req.RefreshToken,
	)

	if err != nil {
		return nil, err
	}

	newAccessToken, err := GenerateToken(userID)

	if err != nil {
		return nil, err
	}

	newRefreshToken, err := GenerateRefreshToken(
		userID,
	)

	if err != nil {
		return nil, err
	}

	err = repository.SaveRefreshToken(
		userID,
		newRefreshToken,
		"refresh-device",
		"",
		"",
		time.Now().Add(time.Hour*24*7),
	)

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	}, nil
}

func Logout(req dto.LogoutRequest) error {

	return repository.DeleteRefreshToken(
		req.RefreshToken,
	)
}

func GetSessions(
	userID int,
) ([]dao.RefreshToken, error) {

	return repository.FindUserSessions(
		userID,
	)
}

func GetUsers(
	page int,
	limit int,
	search string,
) (*dto.PaginationResponse, error) {

	offset := (page - 1) * limit

	users, err := repository.GetUsers(
		limit,
		offset,
		search,
	)

	if err != nil {
		return nil, err
	}

	total, err := repository.CountUsers(search)

	if err != nil {
		return nil, err
	}

	totalPages := int(
		math.Ceil(
			float64(total) / float64(limit),
		),
	)

	return &dto.PaginationResponse{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		Data:       users,
	}, nil
}

func UpdateUser(
	id int,
	req dto.RegisterRequest,
) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)

	return repository.UpdatedUser(
		id,
		req,
	)
}

func DeleteUser(id int) error {

	return repository.DeletedUser(
		id,
	)
}
