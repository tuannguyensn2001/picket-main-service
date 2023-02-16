package auth_transport

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"picket-main-service/src/app"
	"picket-main-service/src/dto"
	"picket-main-service/src/entities"
	"picket-main-service/src/utils"
)

type IUsecase interface {
	Register(ctx context.Context, input dto.RegisterInput) error
	LoginGoogle(ctx context.Context, code string) (*dto.LoginOutput, error)
	Login(ctx context.Context, input dto.LoginInput) (*dto.LoginOutput, error)
	GetProfile(ctx context.Context, id int) (*entities.User, error)
}

type transport struct {
	usecase IUsecase
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) Register(ctx *gin.Context) {
	var input dto.RegisterInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.NewBadRequestError(err, "data not valid"))
	}
	err := t.usecase.Register(ctx.Request.Context(), input)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, app.Response{
		Message: "success",
	})
}

func (t *transport) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.NewBadRequestError(err, "data not valid"))
	}
	result, err := t.usecase.Login(ctx.Request.Context(), input)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, app.Response{
		Message: "success",
		Data:    result,
	})
}

func (t *transport) LoginGoogle(ctx *gin.Context) {
	var input dto.LoginGoogleInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.NewBadRequestError(err, "data not valid"))
	}

	result, err := t.usecase.LoginGoogle(ctx.Request.Context(), input.Code)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, app.Response{
		Data:    result,
		Message: "success",
	})
}

func (t *transport) GetProfile(ctx *gin.Context) {
	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		panic(app.NewForbiddenError(err, "forbidden"))
	}
	user, err := t.usecase.GetProfile(ctx.Request.Context(), userId)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, app.Response{
		Data:    user,
		Message: "success",
	})
}
