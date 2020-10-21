package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PrintInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		method := c.Request.Method
		c.Next()
		status := c.Writer.Status()
		idToken := c.Request.Header.Get("Id-token")
		fmt.Println("idToken is :" + idToken)
		fmt.Printf("请求地址为 %v,请求method为 %v,请求状态为 %v\n", path, method, status)
	}
}

// 跨域访问配置
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headers []string
		for key, _ := range context.Request.Header {
			headers = append(headers, key)
		}
		headerStr := strings.Join(headers, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin,access-contol-allow-headers,%s", headerStr)
		} else {
			headerStr = "access-control-allow-origin,access-control-allow-headers"
		}
		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
			context.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
			context.Writer.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
			context.Writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
			context.Writer.Header().Add("Access-Control-Max-Age", "172800")
			context.Writer.Header().Set("content-type", "application/json;charset=UTF-8") //返回数据格式是json
		}
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request")
		}
		context.Next()
	}
}
