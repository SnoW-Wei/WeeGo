/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-20 13:00:18
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 13:02:42
 */
package migrations

import (
	"database/sql"
	"weego/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		City string `gorm:"type:varchar(10);"`
        Introduction string `gorm:"type:varchar(255);"`
        Avatar string `gorm:"type:varchar(255);defalut:null"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&User{}, "City")
        migrator.DropColumn(&User{}, "Introduction")
        migrator.DropColumn(&User{}, "Avatar")
	}

	migrate.Add("2022_03_20_050018_add_fields_to_user", up, down)
}
