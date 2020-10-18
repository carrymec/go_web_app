package dao

import (
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
