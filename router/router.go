package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zzayne/go-blog/config"
	"github.com/zzayne/go-blog/controller"
)

// Route 路由
func Route(router *gin.Engine) {

	apiPrefix := config.ServerConfig.APIPrefix
	api := router.Group(apiPrefix)
	{
		atricle := new(controller.ArticleController)
		api.GET("/test", atricle.List)

		user := new(controller.UserController)
		api.POST("/signup", user.SignUp)
	}
}
