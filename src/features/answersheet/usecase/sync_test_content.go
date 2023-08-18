package answersheet_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

func (u *usecase) SyncTestContent(ctx context.Context, testId int) error {
	content, err := u.testUsecase.GetContent(ctx, testId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	test, err := u.testUsecase.GetById(ctx, testId)
	if err != nil {
		log.Error().Err(err).Send()
		return nil
	}

	test.Content = content

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(test); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	w := kafka.Writer{
		Addr:                   kafka.TCP(u.config.GetKafkaUrl()),
		Topic:                  "sync-test",
		AllowAutoTopicCreation: true,
		MaxAttempts:            15,
		BatchSize:              1,
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
