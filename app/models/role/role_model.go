/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:25:09
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:32:02
 */
package role

import (
	"weego/app/models"
	"weego/pkg/database"
)

type RoleResource struct {
	models.BaseModel

	Name     string  `gorm:"size:100;index;default:'';not null;"` // 角色名称
	Sequence int     `gorm:"index;default:0;"`                    // 排序值
	Memo     *string `gorm:"size:1024;"`                          // 备注
	Status   int     `gorm:"index;default:0;"`                    // 状态(1:启用 2:禁用)
	Creator  uint64  `gorm:""`                                    // 创建者

	models.CommonTimeStampsField
}

func (roleResource *RoleResource) Create() {
	database.DB.Create(&roleResource)
}

func (roleResource *RoleResource) Save() (rowsAffected int64) {
	result := database.DB.Save(&roleResource)
	return result.RowsAffected
}

func (roleResource *RoleResource) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&roleResource)
	return result.RowsAffected
}
