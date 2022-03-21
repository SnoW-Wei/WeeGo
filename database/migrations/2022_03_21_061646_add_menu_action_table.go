/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:16:46
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:16:56
 */
package migrations

import (
	"database/sql"
	"weego/app/models"
	"weego/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type MenuAction struct {
		models.BaseModel

		MenuID uint64 `gorm:"index;not null;"` // 菜单ID
		Code   string `gorm:"size:100;"`       // 动作编号
		Name   string `gorm:"size:100;"`       // 动作名称

		models.CommonTimeStampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&MenuAction{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&MenuAction{})
	}

	migrate.Add("2022_03_21_061646_add_menu_action_table", up, down)
}
