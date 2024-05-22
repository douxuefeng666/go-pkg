/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 16:25:07
 * @LastEditTime: 2024-05-21 16:27:01
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package conn

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        int64           `json:"id"  gorm:"primaryKey"`
	CreatedAt *time.Time      `json:"created_at,omitempty" `
	UpdatedAt *time.Time      `json:"updated_at,omitempty" `
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" `
}

type CommModel struct {
	Id        int64      `json:"id"  gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty" `
	UpdatedAt *time.Time `json:"updated_at,omitempty" `
}
