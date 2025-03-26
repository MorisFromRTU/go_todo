package repo

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"todo-app/models"
)

type AuthRepository struct {
	DB *gorm.DB
}

var authValidate = validator.New()

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) RegisterUser(newUser *models.User) error {
	if err := taskValidate.Struct(newUser); err != nil {
		return err
	}
	return r.DB.Create(newUser).Error
}

func (r *AuthRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AuthRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
