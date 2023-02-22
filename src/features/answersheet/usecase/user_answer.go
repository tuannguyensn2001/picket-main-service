package answersheet_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"picket-main-service/src/constant"
	"picket-main-service/src/dto"
	"picket-main-service/src/entities"
	"time"
)

func (u *usecase) UserAnswer(ctx context.Context, userId int, input dto.UserAnswerInput) error {

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	check, err := u.CheckUserDoingTest(ctx, userId, input.TestId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	if !check {
		log.Error().Err(ErrUserDoingTest)
		return ErrUserDoingTest
	}

	err = u.testUsecase.CheckTestCanDo(ctx, input.TestId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	err = u.testUsecase.CheckTestAndQuestionValid(ctx, input.TestId, input.QuestionId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	event := map[string]interface{}{
		"user_id":         userId,
		"test_id":         input.TestId,
		"event":           "ANSWER",
		"question_id":     input.QuestionId,
		"previous_answer": input.PreviousAnswer,
		"answer":          input.Answer,
		"created_at":      time.Now(),
		"updated_at":      time.Now(),
	}
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(event); err != nil {
		log.Error().Err(err).Send()
		return err
	}
	job := entities.Job{
		Payload: b.String(),
		Status:  constant.INIT,
		Topic:   "answer-test",
	}
	if err := u.jobUsecase.Create(ctx, &job); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "answer-test",
		AllowAutoTopicCreation: true,
		MaxAttempts:            15,
	}
	key := fmt.Sprintf("%d-%d", userId, input.TestId)
	b.Reset()
	if err := json.NewEncoder(b).Encode(map[string]interface{}{
		"job_id":  job.Id,
		"payload": event,
	}); err != nil {
		log.Error().Err(err).Send()
		return err
	}
	if err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: b.Bytes(),
	}); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}
