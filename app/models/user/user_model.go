/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-22 16:52:07
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-04 11:44:21
 */
package user

import (
	"weego/app/models"
	"weego/pkg/database"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimeStampsField
}

// Create 创建用户，通过User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
