package service

import (
	"golang-jwttoken/data/request"
	"golang-jwttoken/data/response"
	"golang-jwttoken/helper"
	"golang-jwttoken/repository"
	"golang-jwttoken/utils"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(repository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: repository,
		Validate:       validate,
	}
}

// GetAll implements UserService
func (u *UserServiceImpl) GetAll() []response.UserResponse {
	//panic("unimplemented")
	result := u.UserRepository.FindAll()
	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Id:       value.ID,
			Username: value.Username,
			Email:    value.Email,
		}
		users = append(users, user)
	}
	return users
}

// Delete implements UserService
func (u *UserServiceImpl) Delete(id int) {
	//panic("unimplemented")
	u.UserRepository.Delete(id)
}

// Remove implements UserService
func (u *UserServiceImpl) Remove(id int) {
	//panic("unimplemented")
	u.UserRepository.Remove(id)
}

// Update implements UserService
func (u *UserServiceImpl) Update(user request.UpdateUserRequest) {
	//panic("unimplemented")
	userData, err := u.UserRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	userData.Email = user.Email
	if user.Password != "" {
		hashData, hash_err := utils.HashPassword(user.Password)
		helper.ErrorPanic(hash_err)
		user.Password = hashData
	}
	u.UserRepository.Save(userData)
}
