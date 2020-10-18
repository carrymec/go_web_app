package main

import (
	"github.com/gin-gonic/gin"
	"go_web_app/entity"
	"go_web_app/middleware"
	"net/http"
)

func main() {

	r := gin.Default()

	r.POST("/login",middleware.PrintInfo(), func(c *gin.Context) {
		var login entity.Login
		var result entity.JsonResult
		if err := c.BindJSON(&login);err != nil{
			result.Code = -1
			result.Msg = "参数异常"
			c.JSON(-1,result)
			return
		}
		result.Code = 1
		result.Data = &login
		result.Msg = "登录成功"
		c.JSON(http.StatusOK,result)
	})

	r.Run(":9000")
}
