package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int             `gorm:"column:id" json:"id"`
	Username  string          `gorm:"column:username" json:"username"`
	Email     string          `gorm:"column:email" json:"email"`
	Password  string          `gorm:"column:password" json:"-"`
	Type      int             `gorm:"column:type" json:"type"`
	Status    int             `gorm:"column:status" json:"status"`
	CreatedAt *time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
	Profile   *Profile        `json:"profile,omitempty"`
}
