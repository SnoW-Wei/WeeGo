/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-21 12:56:09
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-02-27 13:22:55
 */
package main

import (
	"flag"
	"fmt"
	"weego/bootstrap"
	btsConfig "weego/config"
	"weego/pkg/config"

	"github.com/gin-gonic/gin"
)

/**
* init 通常用来
* 对变量进行初始化
* 检查/修复程序的状态
* 注册
* 运行一次计算
**/
func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "加载 .env文件，如 --env=test 加载的是.env.test文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 设置 gin的运行模式，支持debug , release ,test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非release 模式gin 终端打印太多信息，干扰到我们程序中的Log
	// 故此设置为 Release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// new 一个 Gin Engine实例
	r := gin.New()
	// 初始化DB
	bootstrap.SetupDB()

	// 初始化Redis
	bootstrap.SetupRedis()

	// 注册一个路由
	bootstrap.SetupRouter(r)

	// 运行服务，默认 8080端口
	err := r.Run(":" + config.Get("app.port"))

	if err != nil {
		fmt.Println(err.Error())
	}
}
