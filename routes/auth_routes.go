package routes

import (
	"todo-app/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, h *handlers.AuthHanlder) {
	authGroup := r.Group("/users")
	{
		authGroup.POST("/register", h.RegisterUser)
		authGroup.GET("", h.GetUsers)
	}
}
