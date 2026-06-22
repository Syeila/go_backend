package controller

import (
	"auth-service/app/domain/dto"
	"auth-service/app/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register User
// @Description Register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register Request"
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	err := service.Register(req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register Success",
	})
}

// Login godoc
// @Summary Login User
// @Description Login user and get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login Request"
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()

	tokens, err := service.Login(
		req,
		ipAddress,
		userAgent,
	)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tokens,
	})
}

// RefreshToken godoc
// @Summary Refresh Access Token
// @Description Generate new access token using refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh Token"
// @Success 200 {object} map[string]interface{}
// @Router /refresh [post]
func RefreshToken(c *gin.Context) {

	var req dto.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	tokens, err := service.RefreshToken(req)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tokens,
	})
}

// Logout godoc
// @Summary Logout User
// @Description Logout and revoke refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LogoutRequest true "Logout Request"
// @Success 200 {object} map[string]interface{}
// @Router /logout [post]
func Logout(c *gin.Context) {

	var req dto.LogoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	err := service.Logout(req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logout success",
	})
}

// Sessions godoc
// @Summary Get User Sessions
// @Description Get all active sessions for current user
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /sessions [get]
// Get Session
func Sessions(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")

	if !exists {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user not found",
		})

		return
	}

	userID := int(
		userIDValue.(float64),
	)

	sessions, err := service.GetSessions(
		userID,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": sessions,
	})
}

// GetUser godoc
// @Summary Get All User
// @Description Get current all user
// @Tags Auth
// @Produce json
// @Param page query int false "Page Number"
// @Param limit query int false "Limit Data"
// @Param search query string false "Search Name or Email"
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
func GetUsers(c *gin.Context) {

	var req dto.PaginationRequest

	if err := c.ShouldBindQuery(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}

	users, err := service.GetUsers(
		req.Page,
		req.Limit,
		req.Search,
	)

	log.Println("PAGE =", req.Page)
	log.Println("LIMIT =", req.Limit)
	log.Println("SEARCH =", req.Search)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update User
// @Description Update existing user by ID
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body dto.RegisterRequest true "Update User Request"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	err = service.UpdateUser(id, req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update User Success",
	})
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete user by ID
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	fmt.Println(test)

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	fmt.Println("ID:", id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	err = service.DeleteUser(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete User Success",
	})
}
