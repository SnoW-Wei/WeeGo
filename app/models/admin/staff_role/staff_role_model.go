/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:40:57
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:41:13
 */
package staff_role

import (
	"weego/app/models"
	"weego/pkg/database"
)

type StaffRole struct {
	models.BaseModel

	UserID uint64 `gorm:"index;default:0;"` // 用户内码
	RoleID uint64 `gorm:"index;default:0;"` // 角色内码

	models.CommonTimeStampsField
}

func (staffRole *StaffRole) Create() {
	database.DB.Create(&staffRole)
}

func (staffRole *StaffRole) Save() (rowsAffected int64) {
	result := database.DB.Save(&staffRole)
	return result.RowsAffected
}

func (staffRole *StaffRole) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&staffRole)
	return result.RowsAffected
}
