package answersheet_usecase

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	answersheetpb "picket-main-service/src/pb/answer_sheet"
	"time"
)

func (u *usecase) GetLatestStartTime(ctx context.Context, testId int, userId int) (*time.Time, error) {
	client, err := grpc.Dial("localhost:30000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	defer client.Close()
	conn := answersheetpb.NewAnswerSheetServiceClient(client)
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
