/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-23 21:13:07
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-05 11:58:15
 */
package app

import (
	"time"
	"weego/pkg/config"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {

	return config.Get("app.env") == "production"
}

func IsTest() bool {
	return config.Get("app.env") == "test"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezoen"))
	return time.Now().In(chinaTimezone)
}
