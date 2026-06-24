package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {

	router := testRouter

	payload := map[string]string{
		"email":    "test@mail.com",
		"password": "password123",
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// ======================
	// ASSERT STATUS
	// ======================
	assert.Equal(t, http.StatusOK, w.Code)

	// ======================
	// PARSE RESPONSE
	// ======================
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)

	// ======================
	// CHECK "data"
	// ======================
	data := response["data"].(map[string]interface{})

	assert.NotEmpty(t, data["access_token"])
	assert.NotEmpty(t, data["refresh_token"])
}

func TestLoginWrongPassword(t *testing.T) {

	router := testRouter

	payload := map[string]string{
		"email":    "test@mail.com",
		"password": "wrongpassword",
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "invalid email or password", response["message"])
}

func TestLoginUserNotFound(t *testing.T) {

	router := testRouter

	payload := map[string]string{
		"email":    "notfound@mail.com",
		"password": "password123",
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
