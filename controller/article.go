package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zzayne/go-blog/config"
	"github.com/zzayne/go-blog/model"
)

//AccountController ...
type ArticleController struct{}

var articleModel model.Article

func (ctrl *ArticleController) ClientList(c *gin.Context) {
	queryList(c, false)
}
func (ctrl *ArticleController) AdminList(c *gin.Context) {
	queryList(c, true)
}

func queryList(c *gin.Context, isBackend bool) {
	var articles []model.Article
	var err error
	var pageSize, pageNo int
	pageSize = config.AppConfig.PageSize

	if pageNo, err = strconv.Atoi(c.Query("pageNo")); err != nil {
		pageNo = 1
		err = nil
	}

	if pageNo < 1 {
		pageNo = 1
	}

	articles, err = articleModel.List(1, model.Pager{
		PageSize:   pageSize,
		PageNo:     pageNo,
		OrderField: "created_at",
		OrderASC:   "desc",
	}, isBackend)

	if err != nil {
		FailedMsg(c, err.Error())
		return
	}
	SuccessData(c, articles)
}
