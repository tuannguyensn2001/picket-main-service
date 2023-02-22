package job_usecase

import (
	"context"
	"picket-main-service/src/entities"
)

type IRepository interface {
	Create(ctx context.Context, job *entities.Job) error
}

type usecase struct {
	repository IRepository
}

func New(repository IRepository) *usecase {
	return &usecase{repository: repository}
}

func (u *usecase) Create(ctx context.Context, job *entities.Job) error {
	return u.repository.Create(ctx, job)
}
