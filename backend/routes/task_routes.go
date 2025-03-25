package routes

import (
	"todo-app/handlers"
	"todo-app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.Engine, h *handlers.TaskHandler) {
	taskGroup := r.Group("/tasks")
	taskGroup.Use(middleware.AuthMiddleware())
	{
		taskGroup.GET("", h.GetTasks)
		taskGroup.POST("", h.AddTask)
		taskGroup.PUT("/:id", h.CompleteTask)
		taskGroup.PATCH("/:id", h.ChangeTaskTitle)
		taskGroup.DELETE("/:id", h.DeleteTask)
	}
}
