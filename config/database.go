package config

import (
	"fmt"
	"golang-jwttoken/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(config *Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})

	helper.ErrorPanic(err)

	fmt.Println("Connectd successfully to the database")
	return db
}
