package dto

type StartTestInput struct {
	TestId int `json:"test_id" form:"test_id" binding:"required" `
}

type GetCurrentTestOutput = []GetCurrentTestAnswer

type GetCurrentTestAnswer struct {
	QuestionId int    `json:"question_id"`
	Answer     string `json:"answer"`
}

type UserAnswerInput struct {
	TestId         int    `json:"test_id" form:"test_id" binding:"required"`
	QuestionId     int    `json:"question_id" form:"question_id" binding:"required"`
	Answer         string `json:"answer" form:"answer" binding:"required"`
	PreviousAnswer string `json:"previous_answer" form:"previous_answer" `
}

type SubmitTestInput struct {
	UserId int
	TestId int `json:"test_id" form:"test_id" binding:"required"`
}
