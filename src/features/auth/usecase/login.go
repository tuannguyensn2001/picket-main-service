package auth_usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"picket-main-service/src/dto"
	"strconv"
	"time"
)

func (u *usecase) Login(ctx context.Context, input dto.LoginInput) (*dto.LoginOutput, error) {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	user, err := u.repository.FindByEmail(ctx, input.Email)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
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
