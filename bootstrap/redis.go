/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-26 10:56:37
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-02-26 10:58:44
 */
package bootstrap

import (
	"fmt"
	"weego/pkg/config"
	"weego/pkg/redis"
)

func SetupRedis() {

	redis.ConnectRedis(

		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
