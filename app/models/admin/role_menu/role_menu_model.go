/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:36:56
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:37:07
 */
package role_menu

import (
	"weego/app/models"
	"weego/pkg/database"
)

type RoleMenu struct {
	models.BaseModel

	RoleID   uint64 `gorm:"index;not null;"` // 角色ID
	MenuID   uint64 `gorm:"index;not null;"` // 菜单ID
	ActionID uint64 `gorm:"index;not null;"` // 动作ID

	models.CommonTimeStampsField
}

func (roleMenu *RoleMenu) Create() {
	database.DB.Create(&roleMenu)
}

func (roleMenu *RoleMenu) Save() (rowsAffected int64) {
	result := database.DB.Save(&roleMenu)
	return result.RowsAffected
}

func (roleMenu *RoleMenu) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&roleMenu)
	return result.RowsAffected
}
