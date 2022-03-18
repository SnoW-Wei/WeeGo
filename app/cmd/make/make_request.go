/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-18 12:05:45
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-18 12:12:37
 */
package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file , example make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeRequest(cmd *cobra.Command, args []string) {
	// 格式化模型名称，返回一个model对象
	model := makeModelFromString(args[0])
	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/requests/%s_request.go", model.PackageName)
	// 基于模版创建文件（做好变量替换）
	createFileFromStub(filePath, "request", model)
}
