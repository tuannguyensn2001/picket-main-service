package answersheet_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	answersheetpb "picket-main-service/src/pb/answer_sheet"
	"time"
)

func (u *usecase) CheckUserDoingTest(ctx context.Context, userId int, testId int) (bool, error) {

	client, err := grpc.Dial("localhost:30000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Send()
		return false, err
	}
	defer client.Close()
	conn := answersheetpb.NewAnswerSheetServiceClient(client)
	respGetLatestTime, err := conn.GetLatestStartTime(ctx, &answersheetpb.GetLatestStartTimeRequest{
		TestId: int64(testId),
		UserId: int64(userId),
	})
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
	resp, err := conn.CheckUserDoingTest(ctx, &answersheetpb.CheckUserDoingTestRequest{
		UserId: int64(userId),
		TestId: int64(testId),
	})
	if err != nil {
		return false, err
	}

	return resp.Check, nil
}
