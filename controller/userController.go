package controller

import (
	"github.com/gin-gonic/gin"
	"go_web_app/param"
	"go_web_app/service"
	"go_web_app/tool"
)

type UserController struct {
}

func (user *UserController) Router(r *gin.Engine) {
	r.GET("/api/sendMsg", user.sendMsg)
	r.POST("/api/login_sms", user.loginSms)
}
func (user *UserController) sendMsg(c *gin.Context) {
	//发送验证码
	phone, ok := c.GetQuery("phone")
	if !ok {
		tool.Failed(c, "参数异常")
		return
	}
	userService := service.UserService{}
	sendOk, code := userService.SendCode(phone)
	if sendOk {
		tool.Success(c, code)
		return
	} else {
		tool.Failed(c, "发送失败")
		return
	}
}

func (user *UserController) loginSms(c *gin.Context) {
	var smsLogin param.SmsLogin
	err := tool.Decode(c.Request.Body, &smsLogin)
	if err != nil {
		tool.Failed(c, "参数异常")
		return
	}
	userService := service.UserService{}
	hasUser := userService.FindByPhone(smsLogin)
	if hasUser != nil {
		tool.Success(c, hasUser)
		return
	}
	tool.Failed(c, "登陆失败")
}
