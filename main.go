/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-21 12:56:09
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-02-24 20:14:25
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

	bootstrap.SetupLogger()
	// new 一个 Gin Engine实例
	r := gin.New()
	// 初始化DB
	bootstrap.SetupDB()

	// 注册一个路由
	bootstrap.SetupRouter(r)

	// 运行服务，默认 8080端口
	err := r.Run(":" + config.Get("app.port"))

	if err != nil {
		fmt.Println("err.Error()")
	}
}
