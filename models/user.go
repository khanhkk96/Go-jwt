package models

import "gorm.io/gorm"

type User struct {
	ID       int    `gorm:"type:int;primary_key"`
	Username string `gorm:"type:varchar(50);not null"`
	Email    string `gorm:"type:varchar(256);uniqueIndex;not null"`
	Password string `gorm:"not null"`
	gorm.Model
}
