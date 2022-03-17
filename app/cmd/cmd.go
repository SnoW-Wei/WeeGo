/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-17 18:29:56
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-17 19:30:23
 */
package cmd

import (
	"os"
	"weego/pkg/helpers"

	"github.com/spf13/cobra"
)

// Env 存储全局选项 --env 的值
var Env string

// RegisterGlobalFlags 注册全局选项(flag)
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .env file , example: --env=test will use .env.test file")
}

func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := helpers.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}
