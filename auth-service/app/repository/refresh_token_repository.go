package repository

import (
	"auth-service/app/domain/dao"
	"auth-service/config"
	"time"
)

func SaveRefreshToken(
	userID int,
	token string,
	deviceName string,
	ipAddress string,
	userAgent string,
	expiredAt time.Time,
) error {

	query := `
		INSERT INTO refresh_tokens(
			user_id,
			token,
			device_name,
			ip_address,
			user_agent,
			expired_at
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := config.DB.Exec(
		query,
		userID,
		token,
		deviceName,
		ipAddress,
		userAgent,
		expiredAt,
	)

	return err
}

func FindRefreshToken(token string) error {

	query := `
		SELECT id
		FROM refresh_tokens
		WHERE token = ?
	`

	var id int

	err := config.DB.QueryRow(
		query,
		token,
	).Scan(&id)

	return err
}

func DeleteRefreshToken(token string) error {

	query := `
		DELETE FROM refresh_tokens
		WHERE token = ?
	`

	_, err := config.DB.Exec(
		query,
		token,
	)

	return err
}

func FindUserSessions(userID int) ([]dao.RefreshToken, error) {

	query := `
		SELECT
			id,
			user_id,
			token,
			device_name,
			ip_address,
			user_agent,
			expired_at,
			created_at
		FROM refresh_tokens
		WHERE user_id = ?
	`

	rows, err := config.DB.Query(
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sessions []dao.RefreshToken

	for rows.Next() {

		var session dao.RefreshToken

		err := rows.Scan(
			&session.ID,
			&session.UserID,
			&session.Token,
			&session.DeviceName,
			&session.IPAddress,
			&session.UserAgent,
			&session.ExpiredAt,
			&session.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		sessions = append(
			sessions,
			session,
		)
	}

	return sessions, nil
}
