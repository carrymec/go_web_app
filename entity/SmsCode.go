package entity

type SmsCode struct {
	Id         int    `xorm:"pk autoincr" json:"id"`
	Phone      string `xorm:"varchar(11)" json:"phone"`
	BizId      string `xorm:"varchar(64)" json:"biz_id"`
	Code       string `xorm:"varchar(4)" json:"code"`
	CreateTime int64  `xorm:"bigint" json:"create_time"`
}
