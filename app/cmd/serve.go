/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-17 17:57:56
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-17 18:10:58
 */
package cmd

import (
	"weego/bootstrap"
	"weego/pkg/config"
	"weego/pkg/console"
	"weego/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	// 设置 gin 的运行模式，支持 debug 和 release test三种模式
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式下，gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release , 有特殊情况手动改为debug 即可

	gin.SetMode(gin.ReleaseMode)

	// gin 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务器
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit("Unable to start the server, error: " + err.Error())
	}
}
