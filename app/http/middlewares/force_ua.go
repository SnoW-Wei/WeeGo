/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-20 14:51:26
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 14:53:37
 */
package middlewares

import (
	"errors"
	"weego/pkg/response"

	"github.com/gin-gonic/gin"
)

func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"), "请求必须附带User-Agent标头")
			return
		}
		c.Next()
	}
}
