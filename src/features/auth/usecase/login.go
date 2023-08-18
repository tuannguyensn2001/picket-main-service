package auth_usecase

import (
	"context"
	"github.com/form3tech-oss/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"picket-main-service/src/dto"
	"strconv"
	"time"
)

func (u *usecase) Login(ctx context.Context, input dto.LoginInput) (*dto.LoginOutput, error) {
	ctx, span := tracer.Start(ctx, "validate")
	validate := validator.New()
	span.End()
	if err := validate.Struct(input); err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	ctx, span = tracer.Start(ctx, "query email in db")
	user, err := u.repository.FindByEmail(ctx, input.Email)
	span.End()
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	ctx, span = tracer.Start(ctx, "compare hash")
	time.Sleep(2 * time.Second)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	span.End()
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
	ctx, span = tracer.Start(ctx, "generate token")
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(u.config.GetSecretKey()))
	span.End()
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return &dto.LoginOutput{
		AccessToken: token,
	}, nil

}
