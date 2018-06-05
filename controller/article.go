package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zzayne/go-blog/config"
	"github.com/zzayne/go-blog/model"
)

// ArticleController ...
type ArticleController struct{}

var articleModel model.Article

// ClientList 前端文章列表
func (ctrl *ArticleController) ClientList(c *gin.Context) {
	queryList(c, false)
}

// AdminList 管理端文章列表
func (ctrl *ArticleController) AdminList(c *gin.Context) {
	queryList(c, true)
}

func queryList(c *gin.Context, isBackend bool) {
	var articles []model.Article
	var err error
	var pageSize, pageNo, cateID int
	var noContent bool

	pageSize = config.AppConfig.PageSize

	if pageNo, err = strconv.Atoi(c.Query("pageNo")); err != nil {
		pageNo = 1
		err = nil
	}

	if pageNo < 1 {
		pageNo = 1
	}

	if cateID, err = strconv.Atoi(c.Query("categoryID")); err != nil {
		cateID = 1
		err = nil
	}

	var temp = c.Query("noContent")

	noContent = temp == "true"

	articles, err = articleModel.List(cateID, model.Pager{
		PageSize:   pageSize,
		PageNo:     pageNo,
		OrderField: "created_at",
		OrderASC:   "desc",
	}, isBackend, noContent)

	if err != nil {
		FailedMsg(c, err.Error())
		return
	}
	//temp
	SuccessPageData(c, articles, 5)
}

// Create 创建文章
func (ctrl *ArticleController) Create(c *gin.Context) {
	saveArticle(c, false)
}

// Update 更新文章
func (ctrl *ArticleController) Update(c *gin.Context) {
	saveArticle(c, true)
}

func saveArticle(c *gin.Context, isEdit bool) {
	var article model.Article

	if err := c.ShouldBindJSON(&article); err != nil {
		fmt.Println(err.Error())
		FailedMsg(c, "参数错误")
		return
	}

	userInter, _ := c.Get("user")
	user := userInter.(model.User)

	if err := articleModel.Save(user.ID, article, isEdit); err != nil {
		FailedMsg(c, err.Error())
		return
	}
	SuccessMsg(c, "保存成功")
}

//Preview ...
func (ctrl *ArticleController) Preview(c *gin.Context) {
	info(c, true)
}

//View ...
func (ctrl *ArticleController) View(c *gin.Context) {
	info(c, false)
}

//info 查看内容
func info(c *gin.Context, isBackend bool) {

	var article model.Article
	var err error
	articleID, paramsErr := strconv.Atoi(c.Param("id"))

	if paramsErr != nil {
		FailedMsg(c, "文章id错误")
		return
	}

	format := c.Query("f")

	if article, err = articleModel.Info(articleID, isBackend, format); err != nil {
		FailedMsg(c, err.Error())
		return
	}
	SuccessData(c, article)
}

//Delete 删除
func (ctrl *ArticleController) Delete(c *gin.Context) {

	var delID int
	var err error

	if delID, err = strconv.Atoi(c.Query("id")); err != nil {
		FailedMsg(c, "参数错误")
		return
	}
	if err = articleModel.Delete(delID); err != nil {
		FailedMsg(c, err.Error())
		return
	}
	SuccessMsg(c, "删除成功")

}
