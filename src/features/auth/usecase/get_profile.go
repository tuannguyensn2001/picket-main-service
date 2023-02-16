package auth_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"picket-main-service/src/entities"
)

func (u *usecase) GetProfile(ctx context.Context, id int) (*entities.User, error) {
	result, err := u.repository.FindById(ctx, id)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	return result, nil
}
