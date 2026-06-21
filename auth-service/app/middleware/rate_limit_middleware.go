package middleware

import (
	"context"
	"net/http"
	"time"

	"auth-service/config"

	"github.com/gin-gonic/gin"
)

func LoginRateLimit() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := context.Background()

		ip := c.ClientIP()

		key := "login:" + ip

		count, err := config.RedisClient.Get(
			ctx,
			key,
		).Int()

		if err != nil {

			count = 0
		}

		if count >= 5 {

			c.JSON(
				http.StatusTooManyRequests,
				gin.H{
					"message": "too many login attempts, try again later",
				},
			)

			c.Abort()

			return
		}

		newCount, err := config.RedisClient.Incr(
			ctx,
			key,
		).Result()

		if err != nil {

			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "redis error",
				},
			)

			c.Abort()

			return
		}

		// pertama kali request
		if newCount == 1 {

			config.RedisClient.Expire(
				ctx,
				key,
				time.Minute,
			)
		}

		c.Next()
	}
}
