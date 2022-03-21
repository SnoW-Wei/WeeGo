/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:40:06
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:40:20
 */
package migrations

import (
	"database/sql"
	"weego/app/models"
	"weego/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Staff struct {
		models.BaseModel

		UserName string  `gorm:"size:64;uniqueIndex;default:'';not null;"` // 用户名
		RealName string  `gorm:"size:64;index;default:'';"`                // 真实姓名
		Password string  `gorm:"size:40;default:'';"`                      // 密码
		Email    *string `gorm:"size:255;"`                                // 邮箱
		Phone    *string `gorm:"size:20;"`                                 // 手机号
		Status   int     `gorm:"index;default:0;"`                         // 状态(1:启用 2:停用)
		Creator  uint64  `gorm:""`                                         // 创建者

		models.CommonTimeStampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Staff{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Staff{})
	}

	migrate.Add("2022_03_21_064006_add_staff_table", up, down)
}
