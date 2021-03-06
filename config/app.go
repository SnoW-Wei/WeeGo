/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-21 17:11:46
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 15:04:20
 */
package config

import "weego/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{

			// 应用名称
			"name": config.Env("APP_NAME", "weego"),

			// 当前环境，用于区分多环境，一般为 local，pre，production，test
			"env": config.Env("APP_ENV", "production"),

			// 是否进入调试模式
			"debug": config.Env("APP_DEBUG", false),

			// 应用服务端口
			"port": config.Env("APP_PORT", "9527"),

			// 加密会话，JWT加密
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

			// 用以生成链接
			"url": config.Env("URL", "http://127.0.0.1:9527"),

			// 设置时区，JWT 里会使用，日志记录里也会使用
			"timezone": config.Env("TIMEZONE", "Asia/shanghai"),

			// API 域名，未设置的话所有 API URL 加 api 前缀，如 http://domain.com/api/v1/users
			"api_domain": config.Env("API_DOMAIN"),
		}
	})
}
