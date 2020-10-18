package main

import (
	"github.com/gin-gonic/gin"
	"go_web_app/controller"
	"go_web_app/middleware"
	"go_web_app/tool"
)

func main() {
	config, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}
	_, err = tool.OrmEngine(config)
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	RegisterRouter(r)
	r.Use(middleware.Cors())
	_ = r.Run(config.AppHost + ":" + config.AppPort)
}

// 路由注册
func RegisterRouter(r *gin.Engine) {
	new(controller.Login).Router(r)
	new(controller.UserController).Router(r)
}
