package router

import (
	"auth-service/app/controller"
	"auth-service/app/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "auth-service/docs"
)

func SetupRouter(r *gin.Engine) {

	r.POST("/register", controller.Register)
	r.POST("/login", middleware.LoginRateLimit(), controller.Login)
	r.POST("/refresh", controller.RefreshToken)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/users", controller.GetUsers)

	protected := r.Group("/")

	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", controller.Profile)
	protected.GET("/sessions", controller.Sessions)
	protected.POST("/logout", controller.Logout)
	protected.PUT("/users/:id", controller.UpdateUser)
	protected.DELETE("/users/:id", controller.DeleteUser)
}
