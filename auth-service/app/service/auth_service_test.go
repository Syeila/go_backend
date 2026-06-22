package service

import (
	"math"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHash(t *testing.T) {
	password := "123456"

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		t.Fatalf("failed generate hash: %v", err)
	}

	if string(hash) == password {
		t.Fatal("hash should not equal original password.")
	}
}

func TestPasswordCompare(t *testing.T) {
	password := "123456"

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		t.Fatalf("failed generate hash: %v", err)
	}

	err = bcrypt.CompareHashAndPassword(
		hash,
		[]byte(password),
	)

	if err != nil {
		t.Fatal("password should match")
	}

}

func TestPaginationCalculation(t *testing.T) {

	page := 3
	limit := 10

	offset := (page - 1) * limit

	if offset != 20 {
		t.Errorf(
			"expected offset 20, got %d",
			offset,
		)
	}

	total := 95

	totalPages := int(
		math.Ceil(
			float64(total) / float64(limit),
		),
	)

	if totalPages != 10 {
		t.Errorf(
			"expected total pages 10, got %d",
			totalPages,
		)
	}

}
