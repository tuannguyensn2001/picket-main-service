package answersheet_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"net/http"
	"picket-main-service/src/app"
	"picket-main-service/src/constant"
	"picket-main-service/src/entities"
	"time"
)

var ErrUserDoingTest = app.NewRawError("user doing test", http.StatusConflict)
var ErrTimeNotValid = app.NewRawError("time not valid", http.StatusConflict)

func (u *usecase) Start(ctx context.Context, testId int, userId int) error {
	checkDoing, err := u.CheckUserDoingTest(ctx, userId, testId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	if checkDoing {
		log.Error().Err(ErrUserDoingTest).Send()
		return ErrUserDoingTest
	}

	test, err := u.testUsecase.GetById(ctx, testId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	if test.TimeEnd != nil {
		if test.TimeEnd.Before(time.Now()) {
			log.Error().Err(ErrTimeNotValid).Send()
			return ErrTimeNotValid
		}
	}

	if test.TimeStart != nil {
		if test.TimeStart.After(time.Now()) {
			log.Error().Err(ErrTimeNotValid).Send()
			return ErrTimeNotValid
		}
	}

	event := map[string]interface{}{
		"user_id":    userId,
		"test_id":    testId,
		"event":      "START",
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
		Topic:   "start-test",
	}
	err = u.jobUsecase.Create(ctx, &job)
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
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "start-test",
		AllowAutoTopicCreation: true,
		MaxAttempts:            15,
	}
	err = w.WriteMessages(ctx, kafka.Message{
		Value: b.Bytes(),
	})
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}
