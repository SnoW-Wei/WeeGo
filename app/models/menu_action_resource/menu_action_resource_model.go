/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:19:09
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:19:27
 */
package menu_action_resource

import (
	"weego/app/models"
	"weego/pkg/database"
)

type MenuActionResource struct {
	models.BaseModel

	ActionID uint64 `gorm:"index;not null;"` // 菜单动作ID
	Method   string `gorm:"size:50;"`        // 资源请求方式(支持正则)
	Path     string `gorm:"size:255;"`       // 资源请求路径（支持/:id匹配）

	models.CommonTimeStampsField
}

func (menuActionResource *MenuActionResource) Create() {
	database.DB.Create(&menuActionResource)
}

func (menuActionResource *MenuActionResource) Save() (rowsAffected int64) {
	result := database.DB.Save(&menuActionResource)
	return result.RowsAffected
}

func (menuActionResource *MenuActionResource) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&menuActionResource)
	return result.RowsAffected
}
