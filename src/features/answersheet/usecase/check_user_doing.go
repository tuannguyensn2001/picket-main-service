package answersheet_usecase

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	answersheetpb "picket-main-service/src/pb/answer_sheet"
	"time"
)

func (u *usecase) CheckUserSubmittedTest(ctx context.Context, userId int, testId int) (bool, error) {
	conn := answersheetpb.NewAnswerSheetServiceClient(u.config.GetConnectToAnswersheetService())
	resp, err := conn.CheckUserSubmitted(ctx, &answersheetpb.CheckUserSubmittedRequest{
		TestId: int64(testId),
		UserId: int64(userId),
	})
	if err != nil {
		log.Error().Err(err).Send()
		return false, err
	}
	return resp.Data, nil
}

func (u *usecase) CheckUserDoingTest(ctx context.Context, userId int, testId int) (bool, error) {
	ctx, span := tracer.Start(ctx, "check user doing test")
	defer span.End()
	result, _ := u.redis.Get(ctx, fmt.Sprintf("check_user_doing_%d_%d", userId, testId)).Result()
	if len(result) != 0 {
		return true, nil
	}

	conn := answersheetpb.NewAnswerSheetServiceClient(u.config.GetConnectToAnswersheetService())
	ctx, span = tracer.Start(ctx, "call grpc to answersheet service to get latest time")
	//respGetLatestTime, err := conn.GetLatestStartTime(ctx, &answersheetpb.GetLatestStartTimeRequest{
	//	TestId: int64(testId),
	//	UserId: int64(userId),
	//})
	//span.End()
	//if err != nil {
	//	log.Error().Err(err).Send()
	//	return false, err
	//}
	//if respGetLatestTime.Data != nil {
	//	t := respGetLatestTime.Data.AsTime()
	//	test, err := u.testUsecase.GetById(ctx, testId)
	//	if err != nil {
	//		log.Error().Err(err).Send()
	//		return false, err
	//	}
	//
	//	if test.TimeEnd != nil && test.TimeEnd.Before(time.Now()) {
	//		return false, nil
	//	}
	//	if t.Add(time.Duration(test.TimeToDo) * time.Minute).Before(time.Now()) {
	//		return false, nil
	//	}
	//}
	ctx, span = tracer.Start(ctx, "call grpc to answersheet service to check user doing test")
	resp, err := conn.CheckUserDoingTest(ctx, &answersheetpb.CheckUserDoingTestRequest{
		UserId: int64(userId),
		TestId: int64(testId),
	})
	span.End()
	if err != nil {
		log.Error().Err(err).Send()
		return false, err
	}

	go func() {
		u.redis.Set(context.TODO(), fmt.Sprintf("check_user_doing_%d_%d", userId, testId), true, time.Duration(1)*time.Minute)
	}()

	return resp.Check, nil
}
