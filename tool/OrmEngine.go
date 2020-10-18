package tool

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go_web_app/entity"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

// orm框架接入
func OrmEngine(cfg *Config) (*Orm, error) {
	databaseConfig := cfg.Database
	conn := databaseConfig.User + ":" + databaseConfig.Password + "@tcp(" + databaseConfig.Host +
		":" + databaseConfig.Port + ")/" + databaseConfig.DbName + "?charset=" + databaseConfig.Charset
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(databaseConfig.ShowSql)
	err = engine.Sync2(new(entity.SmsCode))
	if err != nil {
		return nil, err
	}
	err = engine.Sync2(new(entity.User))
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil
}
