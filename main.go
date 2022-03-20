/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-21 12:56:09
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 12:17:01
 */
package main

import (
	"fmt"
	"os"
	"weego/app/cmd"
	"weego/app/cmd/make"
	"weego/bootstrap"
	btsConfig "weego/config"
	"weego/pkg/config"
	"weego/pkg/console"

	"github.com/spf13/cobra"
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

	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "Start web server",
		Long:  `Defalut will run "serve" command, you can use "-h" flag to see all commands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，以来命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
			bootstrap.SetupCache()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		cmd.CmdCache,
	)

	// 配置默认运用 web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)
	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Fail to run app with %v:%s", os.Args, err.Error()))
	}
}
