package repository

import (
	"errors"
	"golang-jwttoken/data/request"
	"golang-jwttoken/helper"
	"golang-jwttoken/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

// Delete implements UserRepository
func (i *UserRepositoryImpl) Delete(userId int) {
	//panic("unimplemented")
	var user models.User
	result := i.Db.Where("id = ?", userId).Unscoped().Delete(&user)
	helper.ErrorPanic(result.Error)
}

// Delete implements UserRepository
func (i *UserRepositoryImpl) Remove(userId int) {
	//panic("unimplemented")
	var user models.User
	result := i.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository
func (i *UserRepositoryImpl) FindAll() []models.User {
	//panic("unimplemented")
	var users []models.User
	result := i.Db.Omit("Password").Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UserRepository
func (i *UserRepositoryImpl) FindById(userId int) (models.User, error) {
	//panic("unimplemented")
	var user models.User
	result := i.Db.Omit("Password").Find(&user, userId)
	if result != nil {
		return user, nil
	}
	return user, errors.New("User is not found")
}

// FindByUsername implements UserRepository
func (i *UserRepositoryImpl) FindByUsername(username string) (models.User, error) {
	//panic("unimplemented")
	var user models.User
	result := i.Db.First(&user, "username = ?", username)

	if result.Error != nil {
		return user, errors.New("Invalid username or password")
	}
	return user, nil
}

// Save implements UserRepository
func (i *UserRepositoryImpl) Save(user models.User) {
	//panic("unimplemented")
	result := i.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository
func (i *UserRepositoryImpl) Update(user models.User) {
	//panic("unimplemented")
	var updateUser = request.UpdateUserRequest{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	result := i.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}
