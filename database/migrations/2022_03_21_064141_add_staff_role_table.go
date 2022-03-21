/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:41:41
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:42:04
 */
package migrations

import (
	"database/sql"
	"weego/app/models"
	"weego/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type StaffRole struct {
		models.BaseModel

		UserID uint64 `gorm:"index;default:0;"` // 用户内码
		RoleID uint64 `gorm:"index;default:0;"` // 角色内码

		models.CommonTimeStampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&StaffRole{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&StaffRole{})
	}

	migrate.Add("2022_03_21_064141_add_staff_role_table", up, down)
}
