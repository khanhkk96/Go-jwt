package controllers

import (
	"golang-jwttoken/data/request"
	"golang-jwttoken/data/response"
	"golang-jwttoken/helper"
	"golang-jwttoken/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	//userRepository repository.UserRepository
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

// Get all users	godoc
// @Security 		BearerAuth
// @Sumary			GetAll
// @Tags			User
// @Produce			application/json
// @Success			200 {object} response.WebResponse{}
// @Router			/user [get]
func (controller *UserController) GetUsers(ctx *gin.Context) {
	users := controller.UserService.GetAll()

	response := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, response)
}

// Update account	godoc
// @Sumary			Update
// @Tags			User
// @Security 		BearerAuth
// @Param			user body request.UpdateUserRequest true "User information"
// @Produce			application/json
// @Success			200 {object} response.WebResponse{}
// @Router			/user [put]
func (controller *UserController) UpdateUser(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Request.URL.Query().Get("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	updateUserRequest.Id = id
	controller.UserService.Update(updateUserRequest)

	response := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully update user data!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, response)
}

// Login		godoc
// @Sumary		Remove
// @Tags		User
// @Security 	BearerAuth
// @Description	Remove an account
// @Param		userId path string true "Remove user by id"
// @Produce		application/json
// @Success		200 {object} response.WebResponse{}
// @Router		/user/{userId} [delete]
func (controler *UserController) Remove(ctx *gin.Context) {
	userId := ctx.Request.URL.Query().Get("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	controler.UserService.Remove(id)

	response := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully remove user",
	}

	ctx.JSON(http.StatusOK, response)
}

// Login		godoc
// @Sumary		Delete
// @Tags		User
// @Security 	BearerAuth
// @Description	Delete an account
// @Param		userId path string true "Delete user by id"
// @Produce		application/json
// @Success		200 {object} response.WebResponse{}
// @Router		/user/delete/{userId} [delete]
func (controler *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Request.URL.Query().Get("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	controler.UserService.Delete(id)

	response := response.WebResponse{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully delete user",
	}

	ctx.JSON(http.StatusOK, response)
}
