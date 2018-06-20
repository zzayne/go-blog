package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zzayne/go-blog/model"
)

//SuccessMsg 成功消息返回
func SuccessMsg(c *gin.Context, msg string) {
	r := new(model.Result)
	r.Msg = msg
	r.Code = model.ErrorCode.SUCCESS
	r.Success = true

	c.JSON(200, r)
}

//SuccessPageData 带总页码的返回结果
func SuccessPageData(c *gin.Context, data interface{}, totalCount int) {
	r := new(model.PageResult)
	r.Data = data
	r.Code = model.ErrorCode.SUCCESS
	r.Success = true
	r.Msg = "success"
	r.TotalCount = totalCount
	c.JSON(200, r)

}

//SuccessData 成功数据返回
func SuccessData(c *gin.Context, data interface{}) {
	r := new(model.Result)
	r.Data = data
	r.Code = model.ErrorCode.SUCCESS
	r.Success = true
	r.Msg = "success"

	c.JSON(200, r)
}

//FailedMsg 失败消息返回
func FailedMsg(c *gin.Context, msg string) {
	r := model.Result{
		Code:    model.ErrorCode.ERROR,
		Success: false,
		Msg:     msg,
	}
	c.JSON(200, r)
}

//FailedResult 失败返回对象
func FailedResult(c *gin.Context, msg string, code int) {
	r := model.Result{
		Code:    code,
		Success: false,
		Msg:     msg,
	}
	c.JSON(200, r)
}

//Unauthorized 未授权返回
func Unauthorized(c *gin.Context, msg string) {
	r := model.Result{
		Code:    model.ErrorCode.Unauthorized,
		Success: false,
		Msg:     "无权限访问",
	}
	c.JSON(200, r)
}
