package main

import (
	"fmt"
	"weego/bootstrap"

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

	// 注册一个路由
	bootstrap.SetupRouter(r)

	// 运行服务，默认 8080端口
	err := r.Run(":3000")

	if err != nil {
		fmt.Println("err.Error()")
	}
}
