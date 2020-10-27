package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web_app/entity"
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
	r.GET("/api/userInfo", user.getUserInfo)
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
		sess, _ := json.Marshal(hasUser)
		err := tool.SetSession(c, "user_"+string(hasUser.Id), sess)
		if err != nil {
			tool.Failed(c, "登陆失败")
			return
		}
		tool.Success(c, hasUser)
		return
	}
	tool.Failed(c, "登陆失败")
	return
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

// 获取用户信息
func (user *UserController) getUserInfo(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		tool.Failed(c, "参数异常")
		return
	}
	fmt.Println("id is :" + id)
	session := tool.GetSession(c, "user_"+id)

	if session == nil {
		tool.Failed(c, "用户未登录")
		return
	}
	err := json.Unmarshal(session.([]byte), entity.User{})
	if err != nil {
		tool.Failed(c, "用户未登录")
		return
	}
	tool.Success(c, session)
	return
}
