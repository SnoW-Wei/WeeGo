/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-21 15:48:02
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-17 17:22:05
 */
package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// SUccess 打印一条成功消息，绿色输出
func Success(msg string) {
	colorOut(msg, "green")
}

// Error 打印一条报错消息，红色输出
func Error(msg string) {
	colorOut(msg, "red")
}

// Waring 打印一条提示消息，黄色输出
func Waring(msg string) {
	colorOut(msg, "yellow")
}

// Exit 打印一条消息，红色输出，并退出程序
func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

// ExitIf 语法糖，自带err != nil判断
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

// colorOut 内部使用，设置高亮颜色
func colorOut(msg string, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(msg, color))
}
