/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:27:04
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:30:59
 */
package migrations

import (
	"database/sql"
	"weego/app/models"
	"weego/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Role struct {
		models.BaseModel

		Name     string  `gorm:"size:100;index;default:'';not null;"` // 角色名称
		Sequence int     `gorm:"index;default:0;"`                    // 排序值
		Memo     *string `gorm:"size:1024;"`                          // 备注
		Status   int     `gorm:"index;default:0;"`                    // 状态(1:启用 2:禁用)
		Creator  uint64  `gorm:""`

		models.CommonTimeStampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Role{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Role{})
	}

	migrate.Add("2022_03_21_062704_add_role_table", up, down)
}
