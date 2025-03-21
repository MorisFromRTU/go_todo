package main

import (
	"fmt"
	"log"
	"os"
	"todo-app/db"
	"todo-app/handlers"
	"todo-app/repo"
	"todo-app/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	if err := db.InitDB(dsn); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	taskRepo := repo.NewTaskRepository(db.DB)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/tasks", taskHandler.GetTasks)
	r.POST("/tasks", taskHandler.AddTask)
	r.PUT("/tasks/:id", taskHandler.CompleteTask)
	r.PATCH("/tasks/:id", taskHandler.ChangeTaskTitle)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
	r.Run(":8000")
}
