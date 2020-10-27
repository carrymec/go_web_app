package tool

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSession(r *gin.Engine) {
	//初始化redis连接
	store, err := redis.NewStore(10, "tcp", GetConfig().RedisConfig.Addr+":"+GetConfig().RedisConfig.Port,
		"", []byte(GetConfig().RedisConfig.Password))
	if err != nil {
		fmt.Println(err.Error())
	}
	r.Use(sessions.Sessions("mysession", store))
}

func SetSession(c *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(c)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

func GetSession(c *gin.Context, key interface{}) interface{} {
	session := sessions.Default(c)
	if session == nil {
		return nil
	}
	return session.Get(key)
}
