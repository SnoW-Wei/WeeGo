/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 23:07:08
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 23:09:42
 */
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	Title      string `valid:"title" json:"title,omitempty"`
	Body       string `valid:"body" json:"body,omitempty"`
	CategoryID string `valid:"category_id" json:"category_id,omitempty"`
}

func TopicSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:3", "max_cn:40"},
		"body":        []string{"required", "min_cn:10", "max_cn:50000"},
		"category_id": []string{"required", "exists:categories,id"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:帖子标题为必填项",
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度需小于 40",
		},
		"body": []string{
			"required:帖子内容为必填项",
			"min_cn:长度需大于 10",
		},
		"category_id": []string{
			"required:帖子分类为必填项",
			"exists:帖子分类未找到",
		},
	}
	return validate(data, rules, messages)
}
