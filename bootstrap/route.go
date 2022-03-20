/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-21 15:32:25
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 15:52:06
 */
package bootstrap

import (
	"net/http"
	"strings"
	"weego/routes"

	"weego/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	// 注册全局中间键
	registerGlobalMiddleWare(router)

	// 注册API路由
	routes.RegisterAPIRoutes(router)

	// 静态资源访问
	setStaticRouter(router)

	// 配置 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
		middlewares.ForceUA(),
	)
}

func setStaticRouter(router *gin.Engine) {
	router.Static("/uploads", "./public/")
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
