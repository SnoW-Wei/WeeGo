/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 21:25:13
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 21:26:41
 */
package migrations

import (
	"database/sql"
	"weego/app/models"
	"weego/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Category struct {
		models.BaseModel

		Name         string `json:"name,omitempty"`
		Descripttion string `json:"descripttion,omitempty"`

		models.CommonTimeStampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Category{})
	}

	migrate.Add("2022_03_19_132513_add_categories_table", up, down)
}
