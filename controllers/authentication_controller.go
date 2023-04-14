package controllers

import (
	"golang-jwttoken/data/request"
	"golang-jwttoken/data/response"
	"golang-jwttoken/helper"
	"golang-jwttoken/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: service}
}

// Login		godoc
// @Sumary		Login account
// @Param		account body request.LoginRequest true "Account info"
// @Tags 		Auth
// @Produce		application/json
// @Success		200 {object} response.WebResponse{}
// @Router		/auth/login [post]
func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)

	token, err_token := controller.authenticationService.Login(loginRequest)
	if err_token != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Successfullly log in!",
		Data:    resp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// Login		godoc
// @Sumary		Register
// @Tags 		Auth
// @Param		user body request.CreateUserRequest true "User information"
// @Produce		application/json
// @Success		200 {object} response.WebResponse{}
// @Router		/auth/register [post]
func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	controller.authenticationService.Register(createUserRequest)
	webResponse := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
