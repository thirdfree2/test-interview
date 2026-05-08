package http

import (
	"be-interview-app/internal/delivery/middleware"

	"github.com/gin-gonic/gin"
)

// MapRoutes จัดกลุ่มและกำหนดเส้นทาง API ทั้งหมด
func MapRoutes(
	r *gin.Engine, 
	userHandler *UserHandler) {

	v1 := r.Group("/api/v1")
	{
		// public routes
		userGroup := v1.Group("/users")
		{
			userGroup.POST("/register", userHandler.Register)
			userGroup.POST("/login", userHandler.Login)
		}

		// protected routes
	{
		authGroup := v1.Group("/")
		authGroup.Use(middleware.AuthMiddleware())
		{
			authGroup.GET("/profile/:id", userHandler.GetUserProfile)
			authGroup.GET("/profile/me", userHandler.GetMe)
		}
	}
	}
}