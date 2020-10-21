package entity

type User struct {
	Id           int64   `xorm:"pk autoincr" json:"id"`
	UserName     string  `xorm:"varchar(32)" json:"user_name"`
	Mobile       string  `xorm:"varchar(11)" json:"mobile"`
	RegisterTime int64   `xorm:"bigint" json:"register_time"`
	Avatar       string  `xorm:"varchar(255)" json:"avatar"`
	Balance      float64 `xorm:"double" json:"balance"`
	IsActive     int8    `xorm:"tinyint" json:"is_active"`
	City         string  `xorm:"varchar(32)" json:"city"`
	Password     string  `xorm:"varchar(64)" json:"password"`
}
