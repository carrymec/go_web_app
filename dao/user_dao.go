package dao

import (
	"fmt"
	"go_web_app/entity"
	"go_web_app/tool"
)

type UserDao struct {
	*tool.Orm
}

func (dao *UserDao) InsertCode(sms entity.SmsCode) int64 {
	rows, err := dao.InsertOne(&sms)
	if err != nil {
		dao.Logger().Error(err)
	}
	return rows
}

func (dao *UserDao) FindByPhone(phone string) *entity.User {
	var user entity.User
	if _, err := dao.Where("mobile = ?", phone).Get(&user); err != nil {
		fmt.Println(err.Error())
	}
	return &user
}

func (dao *UserDao) Valid(phone string, code string) *entity.SmsCode {
	var smsCode entity.SmsCode
	if _, err := dao.Where("phone=? and code=?", phone, code).Get(&smsCode); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &smsCode
}

func (dao *UserDao) InsertUser(user entity.User) int64 {
	rows, err := dao.InsertOne(&user)
	if err != nil {
		fmt.Printf(err.Error())
		return 0
	}
	return rows
}

func (dao *UserDao) QueryByUsername(username string) *entity.User {
	var user entity.User
	if _, err := dao.Where("user_name=?",
		username).Get(&user); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &user
}
