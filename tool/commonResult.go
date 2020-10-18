package tool

import (
	"github.com/gin-gonic/gin"
	"go_web_app/entity"
	"net/http"
)

const (
	SUCCESS = 1
	FAILED  = -1
)

func Success(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, entity.JsonResult{
		Code: SUCCESS,
		Msg:  "成功",
		Data: v,
	})
}

func Failed(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, entity.JsonResult{
		Code: FAILED,
		Msg:  "请求失败",
		Data: v,
	})
}
