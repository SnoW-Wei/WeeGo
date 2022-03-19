/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-20 00:20:24
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 00:20:38
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

	seed.Add("SeedLinksTable", func(db *gorm.DB) {

		links := factories.MakeLinks(5)

		result := db.Table("links").Create(&links)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
