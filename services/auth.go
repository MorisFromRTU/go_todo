package services

import (
	"todo-app/auth"
	"todo-app/models"
	"todo-app/repo"
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
