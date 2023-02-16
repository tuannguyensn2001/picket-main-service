package answersheet_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	answersheetpb "picket-main-service/src/pb/answer_sheet"
)

func (u *usecase) CheckUserDoingTest(ctx context.Context, userId int, testId int) error {

	client, err := grpc.Dial("localhost:30000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	defer client.Close()
	conn := answersheetpb.NewAnswerSheetServiceClient(client)
	respGetLatestTime, err := conn.GetLatestStartTime(ctx, &answersheetpb.GetLatestStartTimeRequest{
		TestId: int64(testId),
		UserId: int64(userId),
	})
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	if respGetLatestTime.Data != nil {
		t := respGetLatestTime.Data.AsTime()

	}

	return nil
}
