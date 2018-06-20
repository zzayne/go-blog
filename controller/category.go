package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zzayne/go-blog/config"
	"github.com/zzayne/go-blog/model"
)

type CategoryController struct{}

var cateModel model.Category

func (ctl *CategoryController) List(c *gin.Context) {
	var pageSize, pageNo int
	var err error
	pageSize = config.AppConfig.PageSize

	if pageNo, err = strconv.Atoi(c.Query("pageNo")); err != nil {
		pageNo = 1
		err = nil
	}

	if pageNo < 1 {
		pageNo = 1
	}

	categories, err := cateModel.List(model.Pager{
		PageSize: pageSize,
		PageNo:   pageNo,
	})

	if err != nil {
		FailedMsg(c, err.Error())
		return
	}

	//temp
	SuccessPageData(c, categories, cateModel.TotalCount())
}

func (ctl CategoryController) Create(c *gin.Context) {
	ctl.Save(c, true)
}
func (ctl CategoryController) Update(c *gin.Context) {
	ctl.Save(c, false)
}

func (ctl CategoryController) Save(c *gin.Context, isNew bool) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		FailedMsg(c, "参数错误")
		return
	}

	if category.ParentID != 0 {
		if _, err := cateModel.Find(category.ParentID); err != nil {
			FailedMsg(c, "无效的父分类")
			return
		}
	}

	if err := cateModel.Save(category, isNew); err != nil {
		FailedMsg(c, "保存失败，请重试")
		return
	}
	SuccessMsg(c, "保存成功")

}

func (ctl CategoryController) Delete(c *gin.Context) {

	var delID int
	var err error

	if delID, err = strconv.Atoi(c.Query("id")); err != nil {
		FailedMsg(c, "参数错误")
		return
	}
	if err = cateModel.Delete(delID); err != nil {
		FailedMsg(c, err.Error())
		return
	}
	SuccessMsg(c, "删除成功")

}
