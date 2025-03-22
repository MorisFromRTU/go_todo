package handlers

import (
	"net/http"
	"todo-app/services"

	"github.com/gin-gonic/gin"
)

type AuthHanlder struct {
	service *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHanlder {
	return &AuthHanlder{service: s}
}

func (h *AuthHanlder) RegisterUser(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := h.service.RegisterUser(requestBody.Username, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusCreated, nil)
}

func (h *AuthHanlder) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *AuthHanlder) LoginUser(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	token, err := h.service.LoginUser(
		requestBody.Username,
		requestBody.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
