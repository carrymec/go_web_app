package service

import (
	"fmt"
	"go_web_app/dao"
	"go_web_app/entity"
	"go_web_app/param"
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
	bizId := fmt.Sprintf("%06v", rand.New(rand.NewSource(
		time.Now().UnixNano())).Int31n(1000000))
	//调用三方短信服务
	//返回结果
	smsCode := entity.SmsCode{
		Phone:      phone,
		BizId:      bizId,
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

func (service *UserService) FindByPhone(login param.SmsLogin) *entity.User {
	//查询数据库 phone是否存在
	userDao := dao.UserDao{Orm: tool.DbEngine}
	smsCode := userDao.Valid(login.Phone, login.Code)
	if smsCode.Id == 0 {
		return nil
	}
	user := userDao.FindByPhone(login.Phone)
	if user.Id != 0 {
		return user
	}
	//新增用户
	newUser := entity.User{}
	newUser.UserName = login.Phone
	newUser.Mobile = login.Phone
	newUser.RegisterTime = time.Now().Unix()

	newUser.Id = userDao.InsertUser(newUser)

	return &newUser
}
