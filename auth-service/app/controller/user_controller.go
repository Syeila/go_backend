package controller

import (
	"auth-service/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Profile godoc
// @Summary User Profile
// @Description Get current user profile
// @Tags User
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /profile [get]
func Profile(c *gin.Context) {

	userIDFloat, _ := c.Get("user_id")

	userID := int(userIDFloat.(float64))

	user, err := service.GetProfile(userID)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
