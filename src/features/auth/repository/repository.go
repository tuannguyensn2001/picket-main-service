package auth_repository

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

func (r *repo) Create(ctx context.Context, user *entities.User) error {
	db := r.GetDB(ctx)
	return db.WithContext(ctx).Create(user).Error
}

func (r *repo) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	db := r.GetDB(ctx)
	var result entities.User
	if err := db.WithContext(ctx).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repo) CreateProfile(ctx context.Context, profile *entities.Profile) error {
	db := r.GetDB(ctx)
	return db.WithContext(ctx).Create(profile).Error
}

func (r *repo) FindById(ctx context.Context, id int) (*entities.User, error) {
	db := r.GetDB(ctx)
	var result entities.User
	if err := db.WithContext(ctx).Preload("Profile").Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
