package db

import (
	"fmt"
	"log"

	"todo-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	fmt.Println("Connected to the database")

	if err := DB.AutoMigrate(
		&models.User{},
		&models.Task{},
	); err != nil {
		return err
	}

	return nil
}
