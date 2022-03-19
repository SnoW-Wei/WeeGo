/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 11:22:15
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 11:29:24
 */
package seeders

import (
	"fmt"
	"weego/database/factories"
	"weego/pkg/console"
	"weego/pkg/logger"
	"weego/pkg/seed"

	"gorm.io/gorm"
)

func init() {
	//添加 Seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		// 创建 10 个用户对象
		users := factories.MakeUsers(10)
		// 批量创建用户（注意批量创建不会调用模型钩子）
		result := db.Table("users").Create(&users)
		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}
		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seede", result.Statement.Table, result.RowsAffected))
	})
}
