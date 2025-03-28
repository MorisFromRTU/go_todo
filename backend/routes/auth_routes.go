package routes

import (
	"todo-app/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, h *handlers.AuthHanlder) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", h.RegisterUser)
		authGroup.GET("", h.GetUsers)
		authGroup.POST("/login", h.LoginUser)
	}
}
