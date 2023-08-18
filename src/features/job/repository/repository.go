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

func (r *repo) UpdateSuccess(ctx context.Context, jobId int) error {
	db := r.GetDB(ctx).WithContext(ctx)

	if err := db.Model(&entities.Job{}).Where("id = ?", jobId).Update("status", entities.SUCCESS).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateFail(ctx context.Context, jobId int, errorMessage string) error {
	db := r.GetDB(ctx).WithContext(ctx)

	if err := db.Model(&entities.Job{}).Where("id = ?", jobId).Updates(map[string]interface{}{
		"status":        entities.FAIL,
		"error_message": errorMessage,
	}).Error; err != nil {
		return err
	}

	return nil
}
