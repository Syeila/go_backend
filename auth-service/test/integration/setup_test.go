package integration

import (
	"auth-service/app/router"
	"auth-service/config"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var testRouter *gin.Engine

func TestMain(m *testing.M) {

	// Test Environment
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3308")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_NAME", "auth_db")

	// Connect DB
	config.ConnectMySQL()

	// Cleanup
	cleanDatabase()

	// Seed User
	seedUser()

	log.Println(">> TEST MAIN RUNNING")

	// Setup Gin
	gin.SetMode(gin.TestMode)

	testRouter = gin.Default()

	router.SetupRouter(testRouter)

	code := m.Run()

	os.Exit(code)
}
