package answersheet_usecase

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"picket-main-service/src/app"
	"time"
)

var ErrTimeTestNotValid = app.NewBadRequestError(errors.New("time test not valid"))

func (u *usecase) GetTimeLeft(ctx context.Context, testId int, userId int) (*time.Duration, error) {
	test, err := u.testUsecase.GetById(ctx, testId)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	//if test.TimeEnd != nil {
	//	log.Info().Str("time-end", test.TimeEnd.UTC().Format("15:04:05 02/01/2006")).
	//		Str("now", time.Now().UTC().Format("15:04:05 02/01/2006")).
	//		Send()
	//	left := test.TimeEnd.Sub(time.Now())
	//	return &left, nil
	//}

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
	if left < 0 {
		return nil, ErrTimeTestNotValid
	}

	return &left, nil
}
