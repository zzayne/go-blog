package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/zzayne/go-blog/model"
)

type UserController struct{}

var userModel model.User

func (ctrl UserController) SignUp(c *gin.Context) {
	type UserForm struct {
		Name     string `json:"name" binding:"required,min=4,max=20"`
		Password string `json:"password" binding:"required,min=6,max=20"`
	}
	var userForm UserForm

	if err := c.ShouldBindWith(&userForm, binding.JSON); err != nil {
		fmt.Println(err)
		FailedMsg(c, err.Error())
		return
	}

	var user model.User
	user, err := userModel.GetUserByName(userForm.Name)
	if err != nil {
		FailedMsg(c, "账号不存在")
		return
	}

	SuccessData(c, user)
}
