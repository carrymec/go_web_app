package controller

import (
	"github.com/gin-gonic/gin"
	"go_web_app/entity"
	"go_web_app/param"
	"net/http"
)

type Login struct {
}

func (l *Login) Router(engine *gin.Engine) {
	engine.GET("/l", l.hello)
	engine.POST("login", l.login)
}
func (l *Login) hello(c *gin.Context) {
	c.JSON(http.StatusOK, entity.JsonResult{
		Code: 1,
		Msg:  "ok",
		Data: "hello",
	})
}

func (l *Login) login(c *gin.Context) {
	var login param.Login
	var result entity.JsonResult
	if err := c.BindJSON(&login); err != nil {
		result.Code = -1
		result.Msg = "参数异常"
		c.JSON(-1, result)
		return
	}
	result.Code = 1
	result.Data = &login
	result.Msg = "登录成功"
	c.JSON(http.StatusOK, result)
}
