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
	r.GET("/api/captcha", user.captcha)
	r.POST("/api/verity", user.verity)
	r.POST("/api/loginByNameAndPwd", user.loginByNameAndPwd)
}

// 发送验证码
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

// 短信登陆
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

// 验证码生成
func (user *UserController) captcha(c *gin.Context) {
	captcha := tool.Generate()
	if captcha == nil {
		tool.Failed(c, "生成验证码失败")
		return
	}
	tool.Success(c, captcha)
}

// 验证验证码
func (user *UserController) verity(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		tool.Failed(c, "id为空")
		return
	}
	code, ok := c.GetQuery("code")
	if !ok {
		tool.Failed(c, "验证码为空")
		return
	}
	ok = tool.Verity(id, code)
	if ok {
		tool.Success(c, "验证成功")
		return
	}
	tool.Failed(c, "验证失败")
}

// 用户名和密码登录
func (user *UserController) loginByNameAndPwd(c *gin.Context) {
	var login param.Login
	err := tool.Decode(c.Request.Body, &login)
	if err != nil {
		tool.Failed(c, "参数异常")
		return
	}
	userService := service.UserService{}
	hasUser := userService.LoginByNameAndPwd(login)
	if hasUser != nil {
		tool.Success(c, hasUser)
		return
	}
	tool.Failed(c, "登陆失败")
}
