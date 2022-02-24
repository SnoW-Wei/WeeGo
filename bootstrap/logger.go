/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-24 19:26:55
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-02-24 20:09:13
 */
package bootstrap

import (
	"weego/pkg/config"
	"weego/pkg/logger"
)

func SetupLogger() {

	logger.InitLogger(

		config.GetString("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.GetString("log.type"),
		config.GetString("log.level"),
	)
}
