package answersheet_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"picket-main-service/src/app"
	"picket-main-service/src/constant"
	"picket-main-service/src/dto"
	"picket-main-service/src/entities"
	"time"
)

func (u *usecase) Submit(ctx context.Context, input dto.SubmitTestInput) error {
	ctx, span := tracer.Start(ctx, "submit test")
	defer span.End()
	_, err := u.GetTimeLeft(ctx, input.TestId, input.UserId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	check, err := u.CheckUserDoingTest(ctx, input.UserId, input.TestId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	if !check {
		return app.NewForbiddenError(errors.New("can't submit"))
	}

	check, err = u.CheckUserSubmittedTest(ctx, input.UserId, input.TestId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	event := map[string]interface{}{
		"user_id":    input.UserId,
		"test_id":    input.TestId,
		"event":      "END",
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(event); err != nil {
		log.Error().Err(err).Send()
		return err
	}
	job := entities.Job{
		Payload: b.String(),
		Status:  constant.INIT,
		Topic:   "submit-test",
	}
	ctx, span = tracer.Start(ctx, "insert job")
	err = u.jobUsecase.Create(ctx, &job)
	span.End()
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	b.Reset()
	err = json.NewEncoder(b).Encode(map[string]interface{}{
		"job_id":  job.Id,
		"payload": event,
	})
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP(u.config.GetKafkaUrl()),
		Topic:                  "submit-test",
		AllowAutoTopicCreation: true,
		MaxAttempts:            15,
		BatchSize:              1,
	}
	ctx, span = tracer.Start(ctx, "push to kafka")
	err = w.WriteMessages(ctx, kafka.Message{
		Value: b.Bytes(),
	})
	span.End()
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}
