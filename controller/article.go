package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zzayne/go-blog/model"
)

//AccountController ...
type ArticleController struct{}

var articleModel model.Article

func (ctrl ArticleController) List(c *gin.Context) {
	SuccessData(c, articleModel.TotalCount())
}
