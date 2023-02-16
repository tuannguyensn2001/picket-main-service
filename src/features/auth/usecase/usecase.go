package auth_usecase

import (
	"context"
	"picket-main-service/src/base"
	"picket-main-service/src/entities"
)

type IRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	CreateProfile(ctx context.Context, profile *entities.Profile) error
	base.IBaseRepository
	FindById(ctx context.Context, id int) (*entities.User, error)
}

type IConfig interface {
	GetSecretKey() string
}

type IGoogleService interface {
	GetUserProfileByAccessToken(ctx context.Context, accessToken string) (*entities.User, error)
	GetAccessTokenFromCode(ctx context.Context, code string) (string, error)
}

type usecase struct {
	repository    IRepository
	config        IConfig
	googleService IGoogleService
}

func New(repository IRepository, config IConfig, googleService IGoogleService) *usecase {
	return &usecase{repository: repository, config: config, googleService: googleService}
}
