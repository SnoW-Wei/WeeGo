/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-21 15:32:25
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-02-25 13:56:20
 */
package bootstrap

import (
	"net/http"
	"strings"
	"weego/routes"

	"weego/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	// 注册全局中间键
	registerGlobalMiddleWare(router)

	// 注册API路由
	routes.RegisterAPIRoutes(router)

	// 配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {

	router.NoRoute(func(c *gin.Context) {
		accpetString := c.Request.Header.Get("Accept")

		if strings.Contains(accpetString, "text/html") {

			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":   404,
				"error_mesage": "路由未定义，请确认URL 和请求方法是否正确。",
			})
		}
	})
}
