package integration

import (
	"auth-service/config"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func seedUser() {

	hash, _ := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)

	res, err := config.DB.Exec(`
		INSERT INTO users
		(name,email,password)
		VALUES (?,?,?)
	`,
		"Test User",
		"test@mail.com",
		string(hash),
	)

	if err != nil {
		log.Fatal("SEED ERROR:", err)
	}

	rows, _ := res.RowsAffected()
	log.Println("ROWS INSERTED:", rows)

	log.Println(">>> SEED USER RUNNING")
}
