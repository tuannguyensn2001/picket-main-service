package answersheet_usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

func (u *usecase) GetTimeLeft(ctx context.Context, testId int, userId int) (*time.Duration, error) {
	test, err := u.testUsecase.GetById(ctx, testId)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	if test.TimeEnd != nil {
		left := test.TimeEnd.Sub(time.Now())
		return &left, nil
	}

	latest, err := u.GetLatestStartTime(ctx, testId, userId)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	zone, _ := latest.Zone()
	zoneNow, _ := time.Now().Zone()
	log.Info().Interface("latest", latest.Format(" 15:04:05 02/01/2006")).
		Interface("timezone", zone).
		Interface("time_zone_now", zoneNow).
		Interface("now", time.Now().Format("15:04:05 02/01/2006")).
		Interface("add", latest.Add(time.Duration(test.TimeToDo)*time.Minute).Format("15:04:05 02/01/2006")).
		Send()

	left := latest.Add(time.Duration(test.TimeToDo) * time.Minute).Sub(time.Now())

	return &left, nil
}
