package answersheet_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	answersheetpb "picket-main-service/src/pb/answer_sheet"
	"time"
)

func (u *usecase) CheckUserDoingTest(ctx context.Context, userId int, testId int) (bool, error) {
	ctx, span := tracer.Start(ctx, "check user doing test")
	defer span.End()
	conn := answersheetpb.NewAnswerSheetServiceClient(u.config.GetConnectToAnswersheetService())
	ctx, span = tracer.Start(ctx, "call grpc to answersheet service to get latest time")
	respGetLatestTime, err := conn.GetLatestStartTime(ctx, &answersheetpb.GetLatestStartTimeRequest{
		TestId: int64(testId),
		UserId: int64(userId),
	})
	span.End()
	if err != nil {
		log.Error().Err(err).Send()
		return false, err
	}
	if respGetLatestTime.Data != nil {
		t := respGetLatestTime.Data.AsTime()
		test, err := u.testUsecase.GetById(ctx, testId)
		if err != nil {
			log.Error().Err(err).Send()
			return false, err
		}

		if test.TimeEnd != nil && test.TimeEnd.Before(time.Now()) {
			return false, nil
		}
		if t.Add(time.Duration(test.TimeToDo) * time.Minute).Before(time.Now()) {
			return false, nil
		}
	}
	ctx, span = tracer.Start(ctx, "call grpc to answersheet service to check user doing test")
	resp, err := conn.CheckUserDoingTest(ctx, &answersheetpb.CheckUserDoingTestRequest{
		UserId: int64(userId),
		TestId: int64(testId),
	})
	span.End()
	if err != nil {
		return false, err
	}

	return resp.Check, nil
}
