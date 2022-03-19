/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 11:27:50
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 11:29:38
 */
package seeders

import "weego/pkg/seed"

func Initialize() {
	// 触发加载本目录下其他文件中的 init 方法

	// 制定优先于同目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
