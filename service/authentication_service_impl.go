package service

import (
	"errors"
	"golang-jwttoken/config"
	"golang-jwttoken/data/request"
	"golang-jwttoken/helper"
	"golang-jwttoken/models"
	"golang-jwttoken/repository"
	"golang-jwttoken/utils"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewAuthenticationServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(user request.LoginRequest) (string, error) {
	//panic("unimplemented")
	userData, err := a.UserRepository.FindByUsername(user.Username)
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	verify_err := utils.VerifyPassword(userData.Password, user.Password)
	if verify_err != nil {
		return "", errors.New("Invalid username or password")
	}

	config, _ := config.LoadConfig(".")
	//Generate token
	token, token_err := utils.GenerateToken(config.TokenExpiredIn, userData.ID, config.TokenSecret)
	helper.ErrorPanic(token_err)
	return token, nil
}

// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(user request.CreateUserRequest) {
	//panic("unimplemented")
	hassPassword, err := utils.HashPassword(user.Password)
	helper.ErrorPanic(err)

	userInfo := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hassPassword,
	}

	a.UserRepository.Save(userInfo)
}
