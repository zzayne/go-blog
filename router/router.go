package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zzayne/go-blog/config"
	"github.com/zzayne/go-blog/controller"
	"github.com/zzayne/go-blog/middleware"
)

// Route 路由
func Route(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())

	apiPrefix := config.ServerConfig.APIPrefix

	article := new(controller.ArticleController)
	category := new(controller.CategoryController)

	api := router.Group(apiPrefix)
	{
		user := new(controller.UserController)
		api.POST("/signin", user.SignIn)

		api.GET("/articles", article.ClientList)

	}

	admin := router.Group(apiPrefix+"/admin", middleware.AdminRequired)
	{
		admin.GET("/articles", article.AdminList)

		admin.GET("/categories", category.List)
		admin.POST("/category", category.Create)
		admin.PUT("/category", category.Update)
		admin.DELETE("/category", category.Delete)

	}

}
