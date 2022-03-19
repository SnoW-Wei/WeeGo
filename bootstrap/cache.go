/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-20 00:58:10
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 01:03:41
 */
package bootstrap

import (
	"fmt"
	"weego/pkg/cache"
	"weego/pkg/config"
)

func SetupCache() {

	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)
	cache.InitWithCacheStore(rds)
}
