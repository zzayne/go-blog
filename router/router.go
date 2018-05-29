package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zzayne/go-blog/config"
	"github.com/zzayne/go-blog/controller"
	"github.com/zzayne/go-blog/middleware"
)

// Route 路由
func Route(router *gin.Engine) {

	apiPrefix := config.ServerConfig.APIPrefix
	api := router.Group(apiPrefix)
	{
		user := new(controller.UserController)
		api.POST("/signin", user.SignIn)

	}

	admin := router.Group(apiPrefix+"/admin", middleware.AdminRequired)
	{
		admin.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "you are admin",
			})
		})
		article := new(controller.ArticleController)
		admin.GET("/articles", article.List)

	}

}
