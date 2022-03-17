/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-17 19:58:32
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-17 20:01:36
 */
package cmd

import (
	"weego/pkg/console"
	"weego/pkg/helpers"

	"github.com/spf13/cobra"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate App Key, will print the generated key",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs, // 不允许传参
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("App Key:")
	console.Success(helpers.RandomNumber(32))
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}
