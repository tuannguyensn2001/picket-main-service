package answersheet_usecase

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	answersheetpb "picket-main-service/src/pb/answer_sheet"
	"time"
)

func (u *usecase) GetLatestStartTime(ctx context.Context, testId int, userId int) (*time.Time, error) {

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

		return &result, nil
	}
	return nil, errors.New("user hasn't started test")

}
