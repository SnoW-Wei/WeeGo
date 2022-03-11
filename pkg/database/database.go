/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-21 20:46:28
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-07 00:11:01
 */
package database

import (
	"database/sql"
	"fmt"

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
