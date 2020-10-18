package service

import (
	"fmt"
	"go_web_app/dao"
	"go_web_app/entity"
	"go_web_app/tool"
	"math/rand"
	"time"
)

type UserService struct {
}

func (service *UserService) SendCode(phone string) (bool, string) {

	//产生验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(
		time.Now().UnixNano())).Int31n(10000))
	//调用三方短信服务
	//返回结果
	smsCode := entity.SmsCode{
		Phone:      phone,
		BizId:      "11111",
		Code:       code,
		CreateTime: time.Now().Unix(),
	}
	userDao := dao.UserDao{Orm: tool.DbEngine}
	rows := userDao.InsertCode(smsCode)
	if rows > 0 {
		return true, code
	}
	return false, ""
}
