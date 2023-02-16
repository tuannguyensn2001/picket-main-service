package auth_usecase

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"picket-main-service/src/app"
	"picket-main-service/src/constant"
	"picket-main-service/src/dto"
	"picket-main-service/src/entities"
)

var ErrUserExisted = app.NewBadRequestError(errors.New("user existed"))

func (u *usecase) Register(ctx context.Context, input dto.RegisterInput) error {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		log.Error().Err(err).Send()
		return err
	}
	user, err := u.repository.FindByEmail(ctx, input.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Send()
			return err
		}
	}
	if user != nil {
		return ErrUserExisted
	}
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	err = u.repository.Transaction(ctx, func(ctx context.Context) error {
		user := entities.User{
			Email:    input.Email,
			Username: input.Username,
			Password: string(password),
			Status:   constant.ACTIVE,
			Type:     constant.TYPE_ACCOUNT_NORMAL,
		}
		err := u.repository.Create(ctx, &user)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}
		profile := entities.Profile{
			UserId: user.Id,
			Avatar: constant.DEFAULT_AVATAR,
		}
		err = u.repository.CreateProfile(ctx, &profile)
		if err != nil {
			log.Error().Err(err).Send()
		}

		return nil
	})

	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}
