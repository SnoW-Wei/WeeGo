/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 22:59:31
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 23:03:37
 */
package migrations

import (
	"database/sql"
	"weego/app/models"
	"weego/app/models/category"
	"weego/app/models/user"
	"weego/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Topic struct {
		models.BaseModel

		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:longtext;not null"`
		UserID     string `gorm:"type:bigint;not null;index"`
		CategoryID string `gorm:"type:bigint;not null;index"`

		// 会创建 user_id 和 category_id 外键的约束
		User     user.User         `json:"user"`
		Category category.Category `json:"category"`

		models.CommonTimeStampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2022_03_19_145931_add_topics_table", up, down)
}
