package integration

import "auth-service/config"

func cleanDatabase() {

	config.DB.Exec("DELETE FROM refresh_tokens")
	config.DB.Exec("DELETE FROM sessions")
	config.DB.Exec("DELETE FROM users")
}
