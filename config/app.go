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
		}
	})
}
