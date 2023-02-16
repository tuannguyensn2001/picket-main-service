package utils

import (
	"context"
	"errors"
	"strings"
)

var ErrBearerTokenNotValid = errors.New("bearer token is not valid")
var ErrValueContext = errors.New("context don't have value")

func GetBearerToken(s string) (string, error) {
	if len(s) == 0 {
		return "", ErrBearerTokenNotValid
	}
	split := strings.Split(s, " ")
	if len(split) != 2 {
		return "", ErrBearerTokenNotValid
	}
	if split[0] != "Bearer" {
		return "", ErrBearerTokenNotValid
	}
	return split[1], nil
}

func GetUserIdFromCtx(ctx context.Context) (int, error) {
	userId, ok := ctx.Value("user_id").(int)
	if !ok {
		return -1, ErrValueContext
	}
	return userId, nil

}
