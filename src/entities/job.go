package entities

import "time"

type Job struct {
	Id           int        `json:"id" gorm:"column:id"`
	Payload      string     `json:"payload" gorm:"column:payload"`
	Status       string     `json:"status" gorm:"column:status"`
	ErrorMessage string     `json:"error_message" gorm:"column:error_message"`
	Topic        string     `json:"topic" gorm:"column:topic"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

const (
	INIT    = "INIT"
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)
