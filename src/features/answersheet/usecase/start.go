package answersheet_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/avast/retry-go"
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
var ErrDidTest = app.NewRawError("user did test", http.StatusConflict)

func (u *usecase) Start(ctx context.Context, testId int, userId int) error {
	ctx, span := tracer.Start(ctx, "start test")
	defer span.End()
	//lastest, err := u.GetLatestStartTime(ctx, testId, userId)
	//if err != nil {
	//	log.Error().Err(err).Send()
	//}
	//if lastest != nil {
	//	log.Error().Err(ErrDidTest).Interface("time", lastest).Send()
	//	return ErrDidTest
	//}

	checkDoing, err := u.CheckUserDoingTest(ctx, userId, testId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	if checkDoing {
		log.Error().Err(ErrUserDoingTest).Send()
		return ErrUserDoingTest
	}

	ctx, span = tracer.Start(ctx, "get test by id")
	test, err := u.testUsecase.GetById(ctx, testId)
	span.End()
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
		log.Info().
			Str("time start", test.TimeStart.UTC().Format("15:04:05 02/01/2006")).
			Str("now", time.Now().UTC().Format("15:04:05 02/01/2006")).
			Bool("after", test.TimeStart.After(time.Now())).
			Send()
		if test.TimeStart.After(time.Now()) {
			log.Error().Err(ErrTimeNotValid).Send()
			return ErrTimeNotValid
		}
	}

	go func() {
		err := retry.Do(func() error {
			return u.SyncTestContent(context.Background(), testId)
		})
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Send()
		}
	}()

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
		Topic:                  "start-test",
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
