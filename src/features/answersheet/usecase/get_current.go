package answersheet_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"picket-main-service/src/dto"
	answersheetpb "picket-main-service/src/pb/answer_sheet"
)

func (u *usecase) GetCurrentTest(ctx context.Context, userId int, testId int) (dto.GetCurrentTestOutput, error) {

	conn := answersheetpb.NewAnswerSheetServiceClient(u.config.GetConnectToAnswersheetService())
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
