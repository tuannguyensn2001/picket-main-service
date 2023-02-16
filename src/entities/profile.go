package entities

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	Id        int             `gorm:"column:id" json:"id"`
	UserId    int             `gorm:"column:user_id" json:"user_id"`
	Avatar    string          `gorm:"column:avatar" json:"avatar"`
	CreatedAt *time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}
