/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 12:04:25
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 12:07:28
 */
package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create model's factory file, example: make factory user",
	Run:   runMakeFactory,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1个参数
}

func runMakeFactory(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 model对象
	model := makeModelFromString(args[0])
	// 拼接目标文件路径
	filePath := fmt.Sprintf("database/factories/%s_factory.go", model.PackageName)
	// 基于模版创建文件（做好变量替换）
	createFileFromStub(filePath, "factory", model)
}
