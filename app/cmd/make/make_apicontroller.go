/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-18 11:18:24
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-18 11:23:56
 */
package make

import (
	"fmt"
	"strings"
	"weego/pkg/console"

	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller , example : make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1), // 只允许且必传 1个参数
}

func runMakeAPIController(cmd *cobra.Command, args []string) {

	// 处理参数，要求附带 API 版本（V1或者V2）
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format : v1/user")
	}

	// apiVersion 用来拼接目标路径
	// name 用来生成 cmd.Model实例
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	// 组建目标目录
	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, model.TableName)

	// 基于模版创建文件（做好变量替换）
	createFileFromStub(filePath, "apicontroller", model)
}
