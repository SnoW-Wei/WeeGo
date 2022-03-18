package migrations

import (
    "database/sql"
    "weego/app/models"
    "weego/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    type User struct {
        models.BaseModel

        Name     string `gorm:"type:varchar(10);not null;index"`
        Email    string `gorm:"type:varchar(50);index;default:null"`
        Phone    string `gorm:"type:char(11);index;default:null"`
        Password string `gorm:"type:varchar(64)"`

        models.CommonTimeStampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&User{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&User{})
    }

    migrate.Add("2022_03_18_100933_add_users_table", up, down)
}