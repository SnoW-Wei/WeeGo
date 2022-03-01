/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-01 19:12:30
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-01 19:30:03
 */
package config

import "weego/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认是 Mailhog 的配置
			"smtp": map[string]interface{}{
				"host":     config.Env("MAIL_HOST", "127.0.0.1"),
				"port":     config.Env("MAIL_PORT", 1025),
				"username": config.Env("MAIL_USERNAME", ""),
				"password": config.Env("MAIL_PASSWORD", ""),
			},
			"from": map[string]interface{}{
				"address": config.Env("MAIL_FROM_ADDRESS", "weego@example.com"),
				"name":    config.Env("MAIL_FROM_NAME", "weego"),
			},
		}
	})
}
