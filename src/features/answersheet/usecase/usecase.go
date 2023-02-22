package answersheet_usecase

import (
	"context"
	"picket-main-service/src/entities"
)

type IRepository interface {
}

type usecase struct {
	repository  IRepository
	testUsecase ITestUsecase
	jobUsecase  IJobUsecase
}

type ITestUsecase interface {
	GetById(ctx context.Context, id int) (*entities.Test, error)
	GetContent(ctx context.Context, testId int) (*entities.TestContent, error)
	CheckTestCanDo(ctx context.Context, testId int) error
	CheckTestAndQuestionValid(ctx context.Context, testId int, questionId int) error
}

type IJobUsecase interface {
	Create(ctx context.Context, job *entities.Job) error
}

func New(repository IRepository, testUsecase ITestUsecase, jobUsecase IJobUsecase) *usecase {
	return &usecase{repository: repository, testUsecase: testUsecase, jobUsecase: jobUsecase}
}
