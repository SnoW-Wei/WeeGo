/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-22 16:52:07
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 12:31:22
 */
package user

import (
	"weego/app/models"
	"weego/pkg/database"
	"weego/pkg/hash"
)

type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`

	City string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Avatar string `json:"avatar,omitempty"`

	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimeStampsField
}

// Create 创建用户，通过User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
