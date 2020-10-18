package tool

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go_web_app/entity"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	databaseConfig := cfg.Database
	//dsn := "root:root@tcp(127.0.0.1:3306)/optimize?charset=utf8mb4&parseTime=True"
	conn := databaseConfig.User + ":" + databaseConfig.Password + "@tcp(" + databaseConfig.Host +
		":" + databaseConfig.Port + ")/" + databaseConfig.DbName + "?charset=" + databaseConfig.Charset
	fmt.Println(conn)
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(databaseConfig.ShowSql)
	err = engine.Sync2(new(entity.SmsCode))
	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil
}
