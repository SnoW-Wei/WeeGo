package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
* init 通常用来
* 对变量进行初始化
* 检查/修复程序的状态
* 注册
* 运行一次计算

func init() {

}
*/

func main() {

	// new 一个 Gin Engine实例
	r := gin.New()

	// 注册中间键
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	r.NoRoute(func(c *gin.Context) {

		accpetString := c.Request.Header.Get("Accept")

		if strings.Contains(accpetString, "text/html") {

			c.String(http.StatusNotFound, "页面返回 404")
		} else {

			c.JSON(http.StatusNotFound, gin.H{
				"error_code":   404,
				"error_mesage": "路由未定义，请求确认URL和请求方法是否正确",
			})
		}
	})

	// 运行服务，默认 8080端口
	r.Run()
}
