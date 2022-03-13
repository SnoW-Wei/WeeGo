/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-22 13:51:55
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-12 03:04:19
 */
package models

import (
	"weego/pkg/localtime"

	"github.com/spf13/cast"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimeStampsField struct {
	CreatedAt localtime.LocalTime `gorm:"comment:'创建时间';type:timestamp;" json:"created_at,omitempty"`
	UpdatedAt localtime.LocalTime `gorm:"comment:'修改时间';type:timestamp;" json:"updated_at,omitempty"`
	//DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;" json:"deleted_at,omitempty"`
}

func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
