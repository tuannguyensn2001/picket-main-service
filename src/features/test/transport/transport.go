package test_transport

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"picket-main-service/src/app"
	"picket-main-service/src/entities"
	"strconv"
)

type IUsecase interface {
	GetPreview(ctx context.Context, id int) (*entities.Test, error)
}

type transport struct {
	usecase IUsecase
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) GetPreview(ctx *gin.Context) {
	testId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(app.NewBadRequestError(err, "data not valid"))
	}
	result, err := t.usecase.GetPreview(ctx.Request.Context(), testId)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, app.Response{
		Message: "success",
		Data:    result,
	})
}
