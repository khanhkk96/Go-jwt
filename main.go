package main

import (
	"golang-jwttoken/config"
	"golang-jwttoken/controllers"
	_ "golang-jwttoken/docs"
	"golang-jwttoken/helper"
	"golang-jwttoken/models"
	"golang-jwttoken/repository"
	"golang-jwttoken/router"
	"golang-jwttoken/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// @title Tag Service API
// @version 1.0
// @description A Tag service API in Go using Gin framework

// @host	localhost:8000
// @BasePath /api
// @securityDefinitions.apikey  BearerAuth
// @in header
// @name Authorization
func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load enviroment variables", err)
	}

	//Database
	db := config.ConnectDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&models.User{})

	//Init Repository
	userRepository := repository.NewUserRepositoryImpl(db)

	//Init service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)
	userService := service.NewUserServiceImpl(userRepository, validate)

	//Init controller
	authenticationController := controllers.NewAuthenticationController(authenticationService)
	userController := controllers.NewUserController(userService)

	//Init router
	router := router.NewRouter(userRepository, authenticationController, userController)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
