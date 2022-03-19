/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-23 21:13:07
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 16:55:08
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

// URL 传参 PATH 拼接站点的 URL
func URL(path string) string {
	return config.Get("app.url") + path
}

// V1URL 拼接带v1 标示 URL
func V1URL(path string) string {
	return URL("/v1/" + path)
}
