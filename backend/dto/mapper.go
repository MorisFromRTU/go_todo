package dto

import (
	"todo-app/models"
)

func MapToUserResponse(user models.User) UserResponse {
	return UserResponse{
		Id:       user.Id,
		UserName: user.UserName,
	}
}

func MapToUsersResponse(users []models.User) []UserResponse {
	var response []UserResponse
	for _, user := range users {
		response = append(response, MapToUserResponse(user))
	}
	return response
}
