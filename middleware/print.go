package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func PrintInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		method := c.Request.Method
		c.Next()
		status := c.Writer.Status()
		fmt.Printf("请求地址为 %v,请求method为 %v,请求状态为 %v\n",path,method,status)
	}
}
