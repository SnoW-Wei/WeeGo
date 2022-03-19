/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 11:16:34
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 11:20:44
 */
package seed

import "gorm.io/gorm"

// 存放所有 Seeder
var seeders []Seeder

// 按顺序执行的Seeder 数组
// 支持一些必须按顺序执行的seeder, 例如：topic 创建的
// 时必须依赖user，所有 topicSeeder 应该在UserSeeder后执行
var orderedSeederNames []string

type SeederFunc func(*gorm.DB)

// Seeder 对应每一个 database/seeders 目录下的Seeder文件
type Seeder struct {
	Func SeederFunc
	Name string
}

// Add 注册到 seeders 数组中
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

// SetRunorder 设置“按顺序执行的 Seeder数组”
func SetRunOrder(names []string) {
	orderedSeederNames = names
}
