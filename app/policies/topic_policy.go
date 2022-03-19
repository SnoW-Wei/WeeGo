/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 23:37:12
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 23:38:14
 */
package policies

import (
	"weego/app/models/topic"
	"weego/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
