package answersheet_usecase

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"picket-main-service/src/config"
	"picket-main-service/src/entities"
)

type IRepository interface {
}

var tracer = otel.Tracer("answersheet_usecase")

type usecase struct {
	repository  IRepository
	testUsecase ITestUsecase
	jobUsecase  IJobUsecase
	config      config.IConfig
	redis       IRedis
}

type ITestUsecase interface {
	GetById(ctx context.Context, id int) (*entities.Test, error)
	GetContent(ctx context.Context, testId int) (*entities.TestContent, error)
	CheckTestCanDo(ctx context.Context, testId int) error
	CheckTestAndQuestionValid(ctx context.Context, testId int, questionId int) error
}

type IRedis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
}

type IJobUsecase interface {
	Create(ctx context.Context, job *entities.Job) error
}

func New(repository IRepository, testUsecase ITestUsecase, jobUsecase IJobUsecase, config config.IConfig, redis IRedis) *usecase {
	return &usecase{repository: repository, testUsecase: testUsecase, jobUsecase: jobUsecase, config: config, redis: redis}
}
