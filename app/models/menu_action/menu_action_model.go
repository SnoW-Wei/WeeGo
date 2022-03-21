/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:10:22
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:14:49
 */
package menu_action

import (
	"weego/app/models"
	"weego/pkg/database"
)

type MenuAction struct {
	models.BaseModel

	MenuID uint64 `gorm:"index;not null;"` // 菜单ID
	Code   string `gorm:"size:100;"`       // 动作编号
	Name   string `gorm:"size:100;"`       // 动作名称

	models.CommonTimeStampsField
}

func (menuAction *MenuAction) Create() {
	database.DB.Create(&menuAction)
}

func (menuAction *MenuAction) Save() (rowsAffected int64) {
	result := database.DB.Save(&menuAction)
	return result.RowsAffected
}

func (menuAction *MenuAction) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&menuAction)
	return result.RowsAffected
}
