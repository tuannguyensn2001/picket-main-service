package auth_usecase

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
	"picket-main-service/src/app"
	"picket-main-service/src/dto"
	"picket-main-service/src/entities"
	"strconv"
	"time"
)

var ErrCodeNotValid = app.NewBadRequestError(errors.New("code not valid"))
var tracer = otel.Tracer("auth_usecase")

func (u *usecase) LoginGoogle(ctx context.Context, code string) (*dto.LoginOutput, error) {
	if len(code) == 0 {
		return nil, ErrCodeNotValid
	}
	result, err := u.googleService.GetAccessTokenFromCode(ctx, code)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	googleAcount, err := u.googleService.GetUserProfileByAccessToken(ctx, result)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	user, err := u.repository.FindByEmail(ctx, googleAcount.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Send()
			return nil, err
		}
		err = u.repository.Transaction(ctx, func(ctx context.Context) error {
			user = &entities.User{
				Email:    googleAcount.Email,
				Username: googleAcount.Username,
			}
			err := u.repository.Create(ctx, user)
			if err != nil {
				log.Error().Err(err).Send()
				return err
			}
			profile := entities.Profile{
				UserId: user.Id,
				Avatar: googleAcount.Profile.Avatar,
			}
			err = u.repository.CreateProfile(ctx, &profile)
			if err != nil {
				log.Error().Err(err).Send()
				return err
			}

			return nil
		})

	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   user.Email,
		ID:        strconv.Itoa(user.Id),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(u.config.GetSecretKey()))
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return &dto.LoginOutput{
		AccessToken: token,
	}, nil

}
