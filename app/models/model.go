/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-02-22 13:51:55
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-04 13:07:41
 */
package models

import (
	"time"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

//TODO 自动补充时间
type CommonTimeStampsField struct {
	CreateAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdateAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}
