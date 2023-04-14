package repository

import "golang-jwttoken/models"

type UserRepository interface {
	Save(user models.User)
	Update(user models.User)
	Delete(userId int)
	Remove(userId int)
	FindById(userId int) (models.User, error)
	FindAll() []models.User
	FindByUsername(username string) (models.User, error)
}
