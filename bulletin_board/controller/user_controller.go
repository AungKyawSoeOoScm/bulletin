package controller

import (
	interfaces "gin_test/bulletin_board/dao/user"
	"gin_test/bulletin_board/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	UsersInterface interfaces.UsersInterface
}

func NewUsercontroller(interfaces interfaces.UsersInterface) *UsersController {
	return &UsersController{UsersInterface: interfaces}
}

func (controller *UsersController) GetUsers(ctx *gin.Context) {
	users := controller.UsersInterface.FindAll()
	helper.ResponseHandler(ctx, http.StatusOK, "Get All Users Success.", users)
}
