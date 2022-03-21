/*
 * @Description: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-21 14:39:35
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-21 14:39:50
 */
package staff

import (
	"weego/app/models"
	"weego/pkg/database"
)

type Staff struct {
	models.BaseModel

	UserName string  `gorm:"size:64;uniqueIndex;default:'';not null;"` // 用户名
	RealName string  `gorm:"size:64;index;default:'';"`                // 真实姓名
	Password string  `gorm:"size:40;default:'';"`                      // 密码
	Email    *string `gorm:"size:255;"`                                // 邮箱
	Phone    *string `gorm:"size:20;"`                                 // 手机号
	Status   int     `gorm:"index;default:0;"`                         // 状态(1:启用 2:停用)
	Creator  uint64  `gorm:""`                                         // 创建者

	models.CommonTimeStampsField
}

func (staff *Staff) Create() {
	database.DB.Create(&staff)
}

func (staff *Staff) Save() (rowsAffected int64) {
	result := database.DB.Save(&staff)
	return result.RowsAffected
}

func (staff *Staff) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&staff)
	return result.RowsAffected
}
