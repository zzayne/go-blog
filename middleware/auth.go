package middleware

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zzayne/go-blog/config"
	"github.com/zzayne/go-blog/controller"
	"github.com/zzayne/go-blog/model"
)

func getUser(c *gin.Context) (model.User, error) {
	var user model.User
	tokenString := c.Request.Header.Get("Token")

	if tokenString == "" {
		return user, errors.New("未登录")
	}

	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.TokenSecret), nil
	})

	if tokenErr != nil {
		return user, errors.New("未登录")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["id"].(float64))
		var userModel model.User
		var err error
		user, err = userModel.GetUserByID(userID)
		if err != nil {
			return user, errors.New("未登录")
		}
		return user, nil
	}
	return user, errors.New("未登录")
}

// AdminRequired 必须是管理员
func AdminRequired(c *gin.Context) {

	var user model.User
	var err error

	if user, err = getUser(c); err != nil {
		controller.FailedResult(c, "未登录", model.ErrorCode.LoginTimeout)
		c.Abort()

		return
	}
	if user.Role == model.UserRoleAdmin {
		c.Set("user", user)
		c.Next()
	} else {
		controller.FailedMsg(c, "没有权限")
		c.Abort()
		return
	}
}
