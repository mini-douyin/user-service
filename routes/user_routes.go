package routes

import (
	"github.com/gin-gonic/gin"
	"user-service/handlers"
	"user-service/repositories"
	"user-service/services"
)

func UserRoutes(r *gin.Engine) {
	userRepo := &repositories.PGUserRepository{}
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// routers
	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", userHandler.Register)
		v1.POST("/users/login", userHandler.Login)
	}
}
