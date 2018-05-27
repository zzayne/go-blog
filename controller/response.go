package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zzayne/go-blog/model"
)

func SuccessMsg(c *gin.Context, msg string) {
	r := new(model.Result)
	r.Msg = msg
	r.Code = model.ErrorCode.SUCCESS
	r.Success = true

	c.JSON(200, r)
}
func SuccessData(c *gin.Context, data interface{}) {
	r := new(model.Result)
	r.Data = data
	r.Code = model.ErrorCode.SUCCESS
	r.Success = true
	r.Msg = "success"

	c.JSON(200, r)
}

//FailedMsg ...
func FailedMsg(c *gin.Context, msg string) {
	r := model.Result{
		Code:    model.ErrorCode.ERROR,
		Success: false,
		Msg:     msg,
	}
	c.JSON(200, r)
}

//UnauthorizedResult ...
func Unauthorized(c *gin.Context, msg string) {
	r := model.Result{
		Code:    model.ErrorCode.Unauthorized,
		Success: false,
		Msg:     "无权限访问",
	}
	c.JSON(200, r)
}
