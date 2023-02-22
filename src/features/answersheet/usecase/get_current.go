package answersheet_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"picket-main-service/src/dto"
	answersheetpb "picket-main-service/src/pb/answer_sheet"
)

func (u *usecase) GetCurrentTest(ctx context.Context, userId int, testId int) (dto.GetCurrentTestOutput, error) {
	client, err := grpc.Dial("localhost:30000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	conn := answersheetpb.NewAnswerSheetServiceClient(client)
	resp, err := conn.GetCurrentTest(ctx, &answersheetpb.GetCurrentTestRequest{
		UserId: int64(userId),
		TestId: int64(testId),
	})
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	result := make(dto.GetCurrentTestOutput, len(resp.Data))
	for index, item := range resp.Data {
		result[index] = dto.GetCurrentTestAnswer{
			QuestionId: int(item.QuestionId),
			Answer:     item.Answer,
		}
	}
	return result, nil
}
