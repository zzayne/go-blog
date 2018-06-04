package controller

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/zzayne/go-blog/config"
	"github.com/zzayne/go-blog/model"
)

//UserController ...
type UserController struct{}

var userModel model.User

//SignIn 用户登陆
func (ctrl *UserController) SignIn(c *gin.Context) {
	type UserForm struct {
		Name     string `json:"name" binding:"required,min=4,max=20"`
		Password string `json:"password" binding:"required,min=6,max=40"`
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
	userModel.Pass = user.Pass

	if userModel.CheckPassword(userForm.Password) == false {
		FailedMsg(c, "用户名或者密码错误")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	tokenString, err := token.SignedString([]byte(config.AppConfig.TokenSecret))

	if err != nil {
		fmt.Println(err.Error())
		FailedMsg(c, "系统内部错误")
		return
	}

	SuccessData(c, gin.H{
		"token": tokenString,
	})

}
