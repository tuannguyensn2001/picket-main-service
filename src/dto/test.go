package dto

import (
	"picket-main-service/src/entities"
	"time"
)

type CreateTestInput struct {
	Name               string `validate:"required"`
	TimeToDo           int    `validate:"required,gt=0"`
	TimeStart          string
	TimeEnd            string
	DoOnce             bool
	Password           string
	PreventCheat       uint8 `validate:"required"`
	IsAuthenticateUser bool
	ShowMark           uint8 `validate:"required"`
	ShowAnswer         uint8 `validate:"required"`
}

type MultipleChoiceAnswer struct {
	TestMultipleChoiceId int
	Answer               string
	Score                float64
	Type                 int32
}

type TestMultipleChoice struct {
	FilePath string
	Score    float64
	Answers  []MultipleChoiceAnswer
}

type CreateTestContentInput struct {
	TestId         int
	Typeable       int
	MultipleChoice *TestMultipleChoice
}

type UpdateTestContentInput struct {
	TestId         int
	Typeable       int
	MultipleChoice *TestMultipleChoice
}

type GetContentOutput struct {
	Content        *entities.TestContent `json:"content"`
	TimeLeft       *time.Duration        `json:"time_left"`
	TimeLeftSecond float64               `json:"time_left_second"`
}
