package answersheet_usecase

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"picket-main-service/src/dto"
)

func (u *usecase) GetContent(ctx context.Context, testId int, userId int) (*dto.GetContentOutput, error) {
	ctx, span := tracer.Start(ctx, "get content")
	defer span.End()

	check, err := u.CheckUserDoingTest(ctx, userId, testId)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	if !check {
		return nil, errors.New("forbidden")
	}
	ctx, span = tracer.Start(ctx, "get content")
	content, err := u.testUsecase.GetContent(ctx, testId)
	span.End()
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	ctx, span = tracer.Start(ctx, "get time left")
	timeLeft, err := u.GetTimeLeft(ctx, testId, userId)
	span.End()
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	output := dto.GetContentOutput{
		Content:        content,
		TimeLeft:       timeLeft,
		TimeLeftSecond: timeLeft.Seconds(),
	}
	return &output, nil
}
