/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:05:01
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:05:13
 */
package migrations

import (
	"database/sql"
	"weego/app/models"
	"weego/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Menu struct {
		models.BaseModel

		Name       string  `gorm:"size:50;index;default:'';not null;"` // 菜单名称
		Icon       *string `gorm:"size:255;"`                          // 菜单图标
		Router     *string `gorm:"size:255;"`                          // 访问路由
		ParentID   *uint64 `gorm:"index;default:0;"`                   // 父级内码
		ParentPath *string `gorm:"size:512;index;default:'';"`         // 父级路径
		IsShow     int     `gorm:"index;default:0;"`                   // 是否显示(1:显示 2:隐藏)
		Status     int     `gorm:"index;default:0;"`                   // 状态(1:启用 2:禁用)
		Sequence   int     `gorm:"index;default:0;"`                   // 排序值
		Memo       *string `gorm:"size:1024;"`                         // 备注
		Creator    uint64  `gorm:""`                                   // 创建人

		models.CommonTimeStampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Menu{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Menu{})
	}

	migrate.Add("2022_03_21_060501_add_menu_table", up, down)
}
