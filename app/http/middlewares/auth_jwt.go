/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-06 21:53:05
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 19:34:21
 */
package middlewares

import (
	"fmt"
	"weego/app/models/user"
	"weego/pkg/config"
	"weego/pkg/jwt"
	"weego/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT(guard string) gin.HandlerFunc {
	return func(c *gin.Context) {

		claims, err := jwt.NewJWT(guard).ParserToken(c)

		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}
		if ok, _ := jwt.NewJWT(guard).BlackListCache(c, "c"); !ok {
			response.Unauthorized(c, "令牌无效")
			return
		}

		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}

		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
