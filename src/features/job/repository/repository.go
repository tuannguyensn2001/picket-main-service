package job_repository

import (
	"context"
	"gorm.io/gorm"
	"picket-main-service/src/base"
	"picket-main-service/src/entities"
)

type repo struct {
	base.Repository
}

func New(db *gorm.DB) *repo {
	return &repo{
		Repository: base.Repository{
			Db: db,
		},
	}
}

func (r *repo) Create(ctx context.Context, job *entities.Job) error {
	db := r.GetDB(ctx)
	return db.WithContext(ctx).Create(job).Error
}
