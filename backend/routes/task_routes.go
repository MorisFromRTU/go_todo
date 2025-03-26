package routes

import (
	"todo-app/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.Engine, h *handlers.TaskHandler) {
	taskGroup := r.Group("/tasks")
	{
		taskGroup.GET("", h.GetTasks)
		taskGroup.POST("", h.AddTask)
		taskGroup.PUT("/:id", h.CompleteTask)
		taskGroup.PATCH("/:id", h.ChangeTaskTitle)
		taskGroup.DELETE("/:id", h.DeleteTask)
	}
}
