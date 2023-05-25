package controller

import (
	"fmt"
	"gin_test/bulletin_board/data/request"
	"gin_test/bulletin_board/helper"
	service "gin_test/bulletin_board/service/user"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	userService service.UserService
}

func NewUsercontroller(service service.UserService) *UsersController {
	return &UsersController{
		userService: service,
	}
}

// Find All
func (controller *UsersController) GetUsers(ctx *gin.Context) {
	users := controller.userService.FindAll()
	// helper.ResponseHandler(ctx, http.StatusOK, "Get All Users Success.", users)
	ctx.HTML(http.StatusOK, "userList.html", gin.H{
		"users": users,
	})
}

// Delete
func (controller *UsersController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.userService.Delete(id)
	ctx.Redirect(http.StatusFound, "/users")
}

// Update
func (controller *UsersController) Update(ctx *gin.Context) {
	userId := ctx.Param("userId")
	username := ctx.PostForm("username")
	email := ctx.PostForm("email")
	utype := ctx.PostForm("type")
	phone := ctx.PostForm("phone")
	dob := ctx.PostForm("dob")
	address := ctx.PostForm("address")
	dobTime, err := time.Parse("2006-01-02", dob)
	if err != nil {
		fmt.Print("date wrong")
	}

	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	if method := ctx.Request.Header.Get("X-HTTP-Method-Override"); method == "PUT" {
		ctx.Request.Method = "PUT"
	}
	updateUserRequest := request.UpdateUserRequest{
		Id:              id,
		Username:        username,
		Email:           email,
		Type:            utype,
		Phone:           phone,
		Address:         address,
		Updated_User_ID: id,
		Date_Of_Birth:   &dobTime,
	}
	controller.userService.Update(updateUserRequest)
	ctx.Redirect(http.StatusFound, "/users")

}

// user create form
func (controller *UsersController) CreateUser(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "usercreateform.html", gin.H{})
}

// user update form
func (controller *UsersController) UpdateForm(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	user := controller.userService.FindById(id)
	ctx.HTML(http.StatusOK, "userupdate.html", gin.H{
		"User": user,
	})
}
