/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-17 20:21:33
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-17 20:26:12
 */
package cmd

import (
	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Like the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清楚测试代码
func runPlay(cmd *cobra.Command, args []string) {

}
