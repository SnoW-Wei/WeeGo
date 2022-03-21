/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-06 23:38:42
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 16:07:09
 */
package middlewares

import (
	"weego/pkg/jwt"
	"weego/pkg/response"

	"github.com/gin-gonic/gin"
)

func GuestJWT(guard string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {

			// 解析 token 成功，说明登录成功
			_, err := jwt.NewJWT(guard).ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
