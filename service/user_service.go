package service

import (
	"golang-jwttoken/data/request"
	"golang-jwttoken/data/response"
)

type UserService interface {
	Update(user request.UpdateUserRequest)
	Remove(id int)
	Delete(id int)
	GetAll() []response.UserResponse
}
