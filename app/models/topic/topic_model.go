/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-19 22:55:54
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-19 22:58:19
 */
package topic

import (
	"weego/app/models"
	"weego/app/models/category"
	"weego/app/models/user"
	"weego/pkg/database"
)

type Topic struct {
	models.BaseModel

	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	// 通过 user_id关联用户
	User user.User `json:"user"`

	// 通过 category_id关联分类
	Category category.Category `json:"category"`
	models.CommonTimeStampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
