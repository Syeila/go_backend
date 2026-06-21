package repository

import "user-service/config"

func CreateUser(
	id int64,
	name string,
	email string,
) error {

	query := `
		INSERT INTO users (
			id,
			name,
			email
		)
		VALUES ($1, $2, $3)
	`
	_, err := config.DB.Exec(
		query,
		id,
		name,
		email,
	)

	return err
}
