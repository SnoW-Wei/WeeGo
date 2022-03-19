/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-20 00:11:04
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-20 00:13:59
 */
package link

import (
	"weego/app/models"
	"weego/pkg/database"
)

type Link struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
    URL string `json:"url,omitempty"`

	models.CommonTimeStampsField
}

func (link *Link) Create() {
	database.DB.Create(&link)
}

func (link *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&link)
	return result.RowsAffected
}

func (link *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&link)
	return result.RowsAffected
}
