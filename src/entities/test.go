package entities

import (
	"gorm.io/gorm"
	"time"
)

type Test struct {
	Id                 int             `json:"id" gorm:"column:id"`
	Code               string          `json:"code" gorm:"column:code"`
	Name               string          `json:"name" gorm:"column:name"`
	TimeToDo           int             `json:"time_to_do" gorm:"column:time_to_do"`
	TimeStart          *time.Time      `json:"time_start" gorm:"column:time_start"`
	TimeEnd            *time.Time      `json:"time_end" gorm:"column:time_end"`
	DoOnce             bool            `json:"do_once" gorm:"column:do_once"`
	Password           string          `json:"password" gorm:"column:password"`
	PreventCheat       uint8           `json:"prevent_cheat" gorm:"column:prevent_cheat"`
	IsAuthenticateUser bool            `json:"is_authenticate_user" gorm:"column:is_authenticate_user"`
	ShowMark           uint8           `json:"show_mark" gorm:"column:show_mark"`
	ShowAnswer         uint8           `json:"show_answer" gorm:"column:show_answer"`
	CreatedBy          int             `json:"created_by" gorm:"column:created_by"`
	CreatedAt          *time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          *time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt          *gorm.DeletedAt `json:"deleted_at,omitempty"`
	Version            int             `json:"version" gorm:"column:version"`
	Content            *TestContent    `json:"content" gorm:"-"`
}

type TestContent struct {
	Id             int                 `json:"id" gorm:"column:id"`
	TestId         int                 `json:"test_id" gorm:"column:test_id"`
	TypeableId     int                 `json:"typeable_id" gorm:"column:typeable_id"`
	Typeable       int                 `json:"typeable" gorm:"column:typeable"`
	CreatedAt      *time.Time          `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      *time.Time          `json:"updated_at" gorm:"column:updated_at"`
	MultipleChoice *TestMultipleChoice `json:"multiple_choice,omitempty" gorm:"-"`
}

type TestMultipleChoiceAnswer struct {
	Id                   int        `json:"id" gorm:"column:id"`
	TestMultipleChoiceId int        `json:"test_multiple_choice_id" gorm:"column:test_multiple_choice_id"`
	Answer               string     `json:"answer" gorm:"column:answer"`
	Score                float64    `json:"score" gorm:"column:score"`
	Type                 int        `json:"type" gorm:"column:type"`
	CreatedAt            *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt            *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type TestMultipleChoice struct {
	Id        int                        `json:"id" gorm:"column:id"`
	FilePath  string                     `json:"file_path" gorm:"column:file_path"`
	Score     float64                    `json:"score" gorm:"column:score"`
	CreatedAt *time.Time                 `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time                 `json:"updated_at" gorm:"column:updated_at"`
	Answers   []TestMultipleChoiceAnswer `json:"answers,omitempty"`
}

func (TestContent) TableName() string {
	return "test_content"
}

func (TestMultipleChoiceAnswer) TableName() string {
	return "test_multiple_choice_answers"
}

func (TestMultipleChoice) TableName() string {
	return "test_multiple_choice"
}
