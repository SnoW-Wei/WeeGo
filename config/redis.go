/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-25 17:15:58
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 00:56:54
 */
package config

import "weego/pkg/config"

func init() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{

			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"password": config.Env("REDIS_PASSWORD", ""),
			// 业务类存储 1 （图片验证码，短信验证码，会话）
			"database": config.Env("REDIS_MAIN_DB", 1),

			// 缓存cache 包使用0 ，缓存清空理应当不影响业务
			"database_cache": config.Env("REDIS_CACHE_DB", 0),
		}
	})
}
