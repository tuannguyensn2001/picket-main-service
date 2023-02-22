package answersheet_transport

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"picket-main-service/src/app"
	"picket-main-service/src/dto"
	"picket-main-service/src/utils"
	"strconv"
)

type IUsecase interface {
	Start(ctx context.Context, testId int, userId int) error
	CheckUserDoingTest(ctx context.Context, userId int, testId int) (bool, error)
	GetContent(ctx context.Context, testId int, userId int) (*dto.GetContentOutput, error)
	GetCurrentTest(ctx context.Context, userId int, testId int) (dto.GetCurrentTestOutput, error)
	UserAnswer(ctx context.Context, userId int, input dto.UserAnswerInput) error
}

type transport struct {
	usecase IUsecase
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}

func (t *transport) Start(ctx *gin.Context) {
	var input dto.StartTestInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.NewBadRequestError(err, "data not valid"))
	}
	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		panic(app.NewForbiddenError(err))
	}
	err = t.usecase.Start(ctx.Request.Context(), input.TestId, userId)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, app.Response{
		Message: "success",
	})
}

func (t *transport) CheckUserDoingTest(ctx *gin.Context) {
	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		panic(app.NewForbiddenError(err))
	}
	testId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(err)
	}
	check, err := t.usecase.CheckUserDoingTest(ctx.Request.Context(), userId, testId)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "message",
		"data":    check,
		"check":   check,
	})
}

func (t *transport) GetContent(ctx *gin.Context) {
	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		panic(app.NewForbiddenError(err))
	}
	testId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(app.NewBadRequestError(err))
	}
	result, err := t.usecase.GetContent(ctx.Request.Context(), testId, userId)

	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, app.Response{
		Message: "success",
		Data:    result,
	})
}

func (t *transport) GetCurrentTest(ctx *gin.Context) {
	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		panic(app.NewForbiddenError(err))
	}
	testId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		panic(app.NewBadRequestError(err))
	}
	result, err := t.usecase.GetCurrentTest(ctx.Request.Context(), userId, testId)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, app.Response{
		Data:    result,
		Message: "success",
	})
}

func (t *transport) UserAnswer(ctx *gin.Context) {
	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		panic(app.NewForbiddenError(err))
	}
	var input dto.UserAnswerInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.NewBadRequestError(err, "data not valid"))
	}
	err = t.usecase.UserAnswer(ctx.Request.Context(), userId, input)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, app.Response{
		Message: "success",
	})
}
