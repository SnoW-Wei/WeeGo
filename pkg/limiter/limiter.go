/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-14 15:29:37
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-17 12:10:20
 */
package limiter

import (
	"strings"
	"weego/pkg/config"
	"weego/pkg/logger"
	"weego/pkg/redis"

	"github.com/gin-gonic/gin"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

// GetKeyIP获取 Limitor 的 Key,IP
func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

// GetKeyRouteWithIP Limitor 的 Key,路由+IP，针对单个路由做限流
func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

func CheckRate(c *gin.Context, key string, formatted string) (limiterlib.Context, error) {

	// 实现化以来的 limiter 包的limiter.Rate 对象
	var context limiterlib.Context
	rate, err := limiterlib.NewRateFromFormatted(formatted)

	if err != nil {
		logger.LogIf(err)
		return context, err
	}
	// 初始化存储，使用我们程序里共用的redis.Redis对象
	store, err := sredis.NewStoreWithOptions(redis.Redis.Client, limiterlib.StoreOptions{
		// 为 limiter 设置前缀，保持redis 里数据的整洁
		Prefix: config.GetString("app.name") + ":limiter",
	})

	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	// 使用上面的初始化的 limiter.Rate 对象和存储对象
	limiterObj := limiterlib.New(store, rate)

	// 获取限流的结果
	if c.GetBool("limiter-once") {
		//peck() 取结果，不增加访问次数
		return limiterObj.Peek(c, key)
	} else {
		// 确保多个路由组里调用 limitIP 进行限流时，只增加一次访问次数
		c.Set("limiter-once", true)

		//Get() 取结果，增加访问次数
		return limiterObj.Get(c, key)
	}
}

// routeToKeyString辅助方法，将url中的 / 格式为 -
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
