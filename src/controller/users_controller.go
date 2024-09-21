package controller

import (
	"net/http"
	"strconv"

	"github.com/Anrewp/go_balance/src/data/request"
	"github.com/Anrewp/go_balance/src/data/response"
	"github.com/Anrewp/go_balance/src/helper"
	"github.com/Anrewp/go_balance/src/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usersService service.UsersService
}

func NewUserController(service service.UsersService) *UserController {
	return &UserController{
		usersService: service,
	}
}

func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	controller.usersService.Create(createUserRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUsersRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	updateUserRequest.Id = id

	controller.usersService.Update(updateUserRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.usersService.Delete(id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.usersService.FindById(id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	userResponse := controller.usersService.FindAll()

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
