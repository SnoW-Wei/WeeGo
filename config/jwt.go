/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-05 11:59:14
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-05 12:02:17
 */
package config

import "weego/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			// 使用 config.GetString("app.key")
			// "sign_key": "",

			// 过期时间，单位是分钟，一般不超过两小时
			"expire_time": config.Env("JWT_EXPIRE_TIME", 120),

			// 允许刷新数据，单位分钟，86400 为两个月，从Token 的签名数据算起
			"max_refresh_time": config.Env("JWT_MAX_REFRESH_TIME", 86400),

			// debug 模式下的过期时间，方便本地开发调试
			"debug_expire_time": 86400,
		}
	})
}
