package answersheet_usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	answersheetpb "picket-main-service/src/pb/answer_sheet"
	"strconv"
	"time"
)

func (u *usecase) GetLatestStartTime(ctx context.Context, testId int, userId int) (*time.Time, error) {

	ctx, span := tracer.Start(ctx, "get latest start time")
	defer span.End()
	result, err := u.redis.Get(ctx, fmt.Sprintf("get_latest_start_time_%d_%d", testId, userId)).Result()
	if len(result) != 0 {
		i, err := strconv.ParseInt(result, 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(i, 0)
		return &tm, nil
	}
	conn := answersheetpb.NewAnswerSheetServiceClient(u.config.GetConnectToAnswersheetService())
	resp, err := conn.GetLatestStartTime(ctx, &answersheetpb.GetLatestStartTimeRequest{
		TestId: int64(testId),
		UserId: int64(userId),
	})
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	if resp.Data != nil {
		//t, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
		result := resp.Data.AsTime()
		log.Info().Str("latest", result.Format("15:04:05 02/01/2006"))

		go func() {
			u.redis.Set(context.TODO(), fmt.Sprintf("get_latest_start_time_%d_%d", testId, userId), result.Unix(), 1*time.Minute).Err()

		}()
		return &result, nil
	}

	return nil, errors.New("user hasn't started test")

}
