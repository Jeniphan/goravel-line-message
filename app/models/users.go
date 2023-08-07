package models

import (
	"github.com/golang-module/carbon/v2"
	"github.com/goravel/framework/database/orm"
	"gorm.io/gorm"
)

type Users struct {
	orm.Model
	Id         uint            `gorm:"primaryKey" json:"id"`
	UserLineID string          `gorm:"size:255;column:user_line_id" form:"user_line_id" json:"user_line_id"`
	UserType   string          `gorm:"size:255;column:user_type" form:"user_type" json:"user_type"`
	CreateAt   carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt  carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

func (r *Users) Users() string {
	return "users"
}
