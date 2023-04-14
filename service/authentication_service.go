package service

import "golang-jwttoken/data/request"

type AuthenticationService interface {
	Login(user request.LoginRequest) (string, error)
	Register(user request.CreateUserRequest)
}
