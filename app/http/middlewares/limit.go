/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-14 15:55:40
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-17 16:34:44
 */
package middlewares

import (
	"net/http"
	"weego/pkg/app"
	"weego/pkg/limiter"
	"weego/pkg/logger"
	"weego/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// LimitIP 全局限流中间件，针对 IP 进行限流
// limit 为格式化字符串，如 “5-S”，示例：
//
// * 5 reqs/second: "5-S"
// * 10 reqs/minute: "10-M"
// * 1000 reqs/hour: "1000-H"
// * 2000 reqs/day: "2000-D"
//
func LimitIP(limit string) gin.HandlerFunc {
	if app.IsTest() {
		limit = "1000000-H"
	}

	return func(c *gin.Context) {
		// 针对IP限流
		key := limiter.GetKeyIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}

		c.Next()
	}
}

// LimitPerRoute 限流中间件，用在单独的路由中
func LimitPerRoute(limit string) gin.HandlerFunc {
	if app.IsTest() {
		limit = "1000000-H"
	}

	return func(c *gin.Context) {
		// 针对单个路由，增加访问次数
		c.Set("limiter-once", false)

		// 针对IP + 路由限流
		key := limiter.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}

		c.Next()
	}
}

func limitHandler(c *gin.Context, key string, limit string) bool {
	// 获取超额的情况
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}

	// ------ 设置表头信息 ------
	// X-RateLimit-Limit :10000最大访问次数
	// X-RateLimit-Remaining :9993剩余访问次数
	// X-RateLimit-Reset :1513784506 到该时间点，访问次数会重置为 X-ReateLimit-Limit
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	// 超额
	if rate.Reached {
		// 提示用户超额
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "接口请求太平凡",
		})
		return false
	}
	return true
}
