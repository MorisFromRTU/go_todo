package services

import (
	"errors"
	"fmt"
	"todo-app/auth"
	"todo-app/models"
	"todo-app/repo"

	"gorm.io/gorm"
)

type AuthService struct {
	repo *repo.AuthRepository
}

func NewAuthService(repo *repo.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(username, password string) error {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
	}
	newUser := &models.User{
		UserName: username,
		Password: hashedPassword,
	}

	err = s.repo.RegisterUser(newUser)
	return err
}

func (s *AuthService) GetUsers() (*[]models.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *AuthService) LoginUser(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid username or password")
		}
		return "", fmt.Errorf("failed to find user: %w", err)
	}

	if !auth.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid username or password")
	}

	token, err := auth.GenerateToken(user.Id)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return token, nil
}
