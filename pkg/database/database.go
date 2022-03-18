/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-21 20:46:28
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-18 21:17:32
 */
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"weego/pkg/config"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// DB对象
var DB *gorm.DB
var SQLDB *sql.DB

// Connect 连接数据库
// TODO 支持多库连接
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	// 使用gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMySQLTables()
	case "sqlite":
		err = deleteAllSqliteTable()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteAllSqliteTable() error {
	tables := []string{}

	// 读取所有数据表
	err := DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type = 'table'").Error
	if err != nil {
		return err
	}

	// 删除所有表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	tables := []string{}

	// 读取所有数据表
	err := DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).Error
	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	DB.Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {

		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	//  开启 Mysql 外键检测
	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}
