/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:19:43
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:19:51
 */
package migrations

import (
	"database/sql"
	"weego/app/models"
	"weego/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type MenuActionResource struct {
		models.BaseModel

		ActionID uint64 `gorm:"index;not null;"` // 菜单动作ID
		Method   string `gorm:"size:50;"`        // 资源请求方式(支持正则)
		Path     string `gorm:"size:255;"`       // 资源请求路径（支持/:id匹配）

		models.CommonTimeStampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&MenuActionResource{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&MenuActionResource{})
	}

	migrate.Add("2022_03_21_061943_add_menu_action_resource_table", up, down)
}
