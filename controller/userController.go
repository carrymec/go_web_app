package controller

import (
	"github.com/gin-gonic/gin"
	"go_web_app/entity"
	"go_web_app/service"
	"net/http"
)

type UserController struct {
}

func (user *UserController) Router(r *gin.Engine) {
	r.GET("/api/sendMsg", user.sendMsg)
}
func (user *UserController) sendMsg(c *gin.Context) {

	//发送验证码
	phone, ok := c.GetQuery("phone")
	if !ok {
		c.JSON(http.StatusOK, entity.JsonResult{
			Code: -1,
			Msg:  "参数异常",
		})
		return
	}
	userService := service.UserService{}
	sendOk, code := userService.SendCode(phone)
	if sendOk {
		c.JSON(http.StatusOK, entity.JsonResult{
			Code: 0,
			Msg:  "发送成功",
			Data: code,
		})
		return
	} else {
		c.JSON(http.StatusOK, entity.JsonResult{
			Code: -1,
			Msg:  "发送失败",
			Data: nil,
		})
		return
	}
}
