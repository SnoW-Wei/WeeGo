/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:37:30
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:37:36
 */
package migrations

import (
	"database/sql"
	"weego/app/models"
	"weego/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type RoleMenu struct {
		models.BaseModel

		RoleID   uint64 `gorm:"index;not null;"` // 角色ID
		MenuID   uint64 `gorm:"index;not null;"` // 菜单ID
		ActionID uint64 `gorm:"index;not null;"` // 动作ID

		models.CommonTimeStampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&RoleMenu{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&RoleMenu{})
	}

	migrate.Add("2022_03_21_063730_add_role_menu_table", up, down)
}
