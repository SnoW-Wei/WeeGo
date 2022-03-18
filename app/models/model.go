/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-22 13:51:55
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-18 23:37:28
 */
package models

import (
	"time"

	"github.com/spf13/cast"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimeStampsField struct {
	CreatedAt time.Time `gorm:"comment:'创建时间';column:created_at;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"comment:'修改时间';column:updated_at;" json:"updated_at,omitempty"`
	//DeletedAt localtime.LocalTime `gorm:"column:deleted_at;comment:'删除时间" json:"deleted_at,omitempty" sql:"index"`
}

func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
