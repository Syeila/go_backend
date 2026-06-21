package repository

import (
	"auth-service/app/domain/dao"
	"auth-service/app/domain/dto"
	"auth-service/config"
	"log"
)

func CreatedUser(req dto.RegisterRequest) (int64, error) {

	query := `
		INSERT INTO users(name, email, password)
		VALUES (?, ?, ?)
	`
	result, err := config.DB.Exec(
		query,
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, err
}

func FindUserByEmail(email string) (dao.User, error) {

	var user dao.User

	query := `
		SELECT id, name, email, password
		FROM users
		WHERE email = ?
	`

	err := config.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	return user, err
}

func FindUserByID(id int) (dao.User, error) {

	var user dao.User

	query := `
		SELECT id, name, email
		FROM users
		WHERE id = ?
	`

	err := config.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	return user, err
}

func GetUsers(
	limit int,
	offset int,
	search string,
) ([]dao.User, error) {

	query := `
		SELECT
			id,
			name,
			email
		FROM users
		WHERE
			name LIKE ?
			OR email LIKE ?
		ORDER BY id
		LIMIT ?
		OFFSET ?
	`

	searchKeyword := "%" + search + "%"

	rows, err := config.DB.Query(
		query,
		searchKeyword,
		searchKeyword,
		limit,
		offset,
	)

	log.Println("SEARCH KEYWORD =", searchKeyword)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []dao.User

	for rows.Next() {

		var user dao.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdatedUser(id int, req dto.RegisterRequest) error {
	query := `
		UPDATE users
		SET name = ?, email= ?, password = ?
		WHERE id = ?
	`

	_, err := config.DB.Exec(
		query,
		req.Name,
		req.Email,
		req.Password,
		id,
	)

	return err
}

func DeletedUser(id int) error {
	query := `
		DELETE FROM users WHERE id = ? 
	`
	_, err := config.DB.Exec(
		query,
		id,
	)

	return err
}

func CountUsers(search string) (int, error) {

	query := `
		SELECT COUNT(*)
		FROM users
		WHERE
			name LIKE ?
			OR email LIKE ?
	`

	searchKeyword := "%" + search + "%"

	var total int

	err := config.DB.QueryRow(
		query,
		searchKeyword,
		searchKeyword,
	).Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
