package test_usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"picket-main-service/src/app"
	"picket-main-service/src/base"
	"picket-main-service/src/constant"
	"picket-main-service/src/dto"
	"picket-main-service/src/entities"
	randompkg "picket-main-service/src/pkg/random"
	"picket-main-service/src/utils"
	"strings"
	"sync"
	"time"
)

type IRepository interface {
	Create(ctx context.Context, test *entities.Test) error
	CreateTestContent(ctx context.Context, test *entities.TestContent) error
	CreateTestMultipleChoice(ctx context.Context, test *entities.TestMultipleChoice) error
	CreateListTestMultipleChoiceAnswers(ctx context.Context, list []entities.TestMultipleChoiceAnswer) error
	base.IBaseRepository
	FindByTestId(ctx context.Context, id int) (*entities.Test, error)
	FindContentByTestId(ctx context.Context, testId int) (*entities.TestContent, error)
	FindTestMultipleChoiceByTestId(ctx context.Context, testId int) (*entities.TestMultipleChoice, error)
	FindTestByUserId(ctx context.Context, userId int) ([]entities.Test, error)
	FindTestMultipleChoiceAnswer(ctx context.Context, multipleChoiceId int) ([]entities.TestMultipleChoiceAnswer, error)
	FindByCode(ctx context.Context, code string) (*entities.Test, error)
}

type usecase struct {
	repository IRepository
	redis      IRedis
	lock       locking
}

type locking struct {
	getContent sync.Mutex
	getById    sync.Mutex
}

var tracer = otel.Tracer("test_usecase")

type IRedis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

func New(repository IRepository, redis IRedis) *usecase {
	var mu sync.Mutex
	lock := locking{
		getContent: mu,
		getById:    mu,
	}
	return &usecase{repository: repository, redis: redis, lock: lock}
}

func (u *usecase) ParseTimeTest(ctx context.Context, val string) (*time.Time, error) {
	if len(val) == 0 {
		return nil, nil
	}
	return utils.ParseTime("HH:MM:SS DD/MM/YYYY", val)
}

func (u *usecase) Create(ctx context.Context, input dto.CreateTestInput, userId int) error {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		log.Error().Err(err).Send()
		return err
	}
	timeStart, err := u.ParseTimeTest(ctx, input.TimeStart)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	timeEnd, err := u.ParseTimeTest(ctx, input.TimeEnd)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	if timeStart != nil && timeEnd != nil {
		if timeStart.After(*timeEnd) {
			log.Error().Err(err).Send()
			return app.NewBadRequestError(errors.New("time start before time end"))
		}
	}

	test := entities.Test{
		Code:               strings.ToUpper(randompkg.StringWithLength(5)),
		Name:               input.Name,
		TimeToDo:           input.TimeToDo,
		TimeStart:          timeStart,
		TimeEnd:            timeEnd,
		DoOnce:             input.DoOnce,
		Password:           input.Password,
		PreventCheat:       input.PreventCheat,
		IsAuthenticateUser: input.IsAuthenticateUser,
		ShowAnswer:         input.ShowAnswer,
		ShowMark:           input.ShowMark,
		CreatedBy:          userId,
		Version:            1,
	}
	err = u.repository.Create(ctx, &test)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) UpdateContent(ctx context.Context, input dto.UpdateTestContentInput) error {

	content, err := u.repository.FindContentByTestId(ctx, input.TestId)
	if err != nil {
		return err
	}
	if content.Typeable == input.Typeable {
		var handler func(ctx context.Context, input dto.UpdateTestContentInput) error
		switch content.Typeable {
		case 1:
			handler = u.UpdateMultipleChoiceContent
		}
		err = handler(ctx, input)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *usecase) UpdateMultipleChoiceContent(ctx context.Context, input dto.UpdateTestContentInput) error {

	return nil
}

func (u *usecase) GetPreview(ctx context.Context, id int) (*entities.Test, error) {
	return u.repository.FindByTestId(ctx, id)
}

func (u *usecase) GetTestsByUserId(ctx context.Context, userId int) ([]entities.Test, error) {
	result, err := u.repository.FindTestByUserId(ctx, userId)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	return result, nil
}

func (u *usecase) GetByIdFromRedis(ctx context.Context, id int) *entities.Test {
	ctx, span := tracer.Start(ctx, "get test from redis")
	defer span.End()
	result := u.redis.Get(ctx, fmt.Sprintf("test-%d", id))
	if result.Err() != nil {
		log.Error().Err(result.Err()).Send()
		return nil
	}
	var output entities.Test
	err := json.NewDecoder(strings.NewReader(result.Val())).Decode(&output)
	if err != nil {
		log.Error().Err(err).Send()
		return nil
	}
	return &output
}
func (u *usecase) SaveTestToRedis(ctx context.Context, test *entities.Test) error {
	ctx, span := tracer.Start(ctx, "save test to redis")
	defer span.End()
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(&test)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	var timeDuration time.Duration
	if test.TimeEnd != nil {
		timeDuration = test.TimeEnd.Sub(time.Now())
	} else {
		timeDuration = time.Duration(test.TimeToDo) * time.Minute
	}
	result := u.redis.Set(ctx, fmt.Sprintf("test-%d", test.Id), b.String(), timeDuration)
	if result.Err() != nil {
		log.Error().Err(result.Err()).Send()
		return result.Err()
	}

	return nil
}

func (u *usecase) GetById(ctx context.Context, id int) (*entities.Test, error) {
	test := u.GetByIdFromRedis(ctx, id)
	if test != nil {
		return test, nil
	}

	u.lock.getById.Lock()
	defer u.lock.getById.Unlock()
	test = u.GetByIdFromRedis(ctx, id)
	if test != nil {
		return test, nil
	}

	test, err := u.repository.FindByTestId(ctx, id)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	retry.Do(func() error {
		return u.SaveTestToRedis(ctx, test)
	})

	return test, nil
}

func (u *usecase) GetContentFromRedis(ctx context.Context, testId int) *entities.TestContent {
	ctx, span := tracer.Start(ctx, "get content from redis")
	if span.IsRecording() {
		span.SetAttributes(attribute.Int("test_id", testId))
	}
	defer span.End()
	result := u.redis.Get(ctx, fmt.Sprintf("test-content-%d", testId))
	if result.Err() != nil {
		return nil
	}
	var output entities.TestContent
	err := json.NewDecoder(strings.NewReader(result.Val())).Decode(&output)
	if err != nil {
		return nil
	}
	log.Info().Interface("content", output).Msg("get content from redis")
	return &output
}

func (u *usecase) SaveContentToRedis(ctx context.Context, content *entities.TestContent) error {
	ctx, span := tracer.Start(ctx, "save content to redis")
	defer span.End()
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(content)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	timeDuration, ok := ctx.Value("time_expire_test").(time.Duration)
	if !ok {
		timeDuration = 2 * time.Minute
	}
	result := u.redis.Set(ctx, fmt.Sprintf("test-content-%d", content.TestId), b.String(), timeDuration)
	if result.Err() != nil {
		log.Error().Err(err).Send()
		return result.Err()
	}
	return nil
}

func (u *usecase) GetContent(ctx context.Context, testId int) (*entities.TestContent, error) {
	content := u.GetContentFromRedis(ctx, testId)
	if content != nil {
		return content, nil
	}

	log.Info().Msg("lock")
	u.lock.getContent.Lock()
	defer u.lock.getContent.Unlock()

	content = u.GetContentFromRedis(ctx, testId)
	if content != nil {
		return content, nil
	}

	content, err := u.repository.FindContentByTestId(ctx, testId)
	if err != nil {
		return nil, err
	}
	multipleChoice, err := u.repository.FindTestMultipleChoiceByTestId(ctx, testId)
	if err != nil {
		return nil, err
	}
	answers, err := u.repository.FindTestMultipleChoiceAnswer(ctx, multipleChoice.Id)
	if err != nil {
		return nil, err
	}
	multipleChoice.Answers = answers
	content.MultipleChoice = multipleChoice

	log.Info().Msg("unlock")
	//go func() {

	//}()
	retry.Do(func() error {

		test, err := u.repository.FindByTestId(ctx, testId)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}
		var timeExpire time.Duration
		if test.TimeEnd != nil {
			timeExpire = test.TimeEnd.Sub(time.Now())
		} else {
			timeExpire = time.Duration(test.TimeToDo) * time.Minute
		}
		ctx = context.WithValue(ctx, "time_expire_test", timeExpire)
		return u.SaveContentToRedis(ctx, content)
	}, retry.Attempts(5))

	return content, nil
}

func (u *usecase) CreateContent(ctx context.Context, input dto.CreateTestContentInput) error {
	var handler func(ctx context.Context, input dto.CreateTestContentInput) error
	switch input.Typeable {
	case 1:
		handler = u.CreateMultipleChoiceContent
	}

	err := handler(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) CreateMultipleChoiceContent(ctx context.Context, input dto.CreateTestContentInput) error {

	test, err := u.repository.FindByTestId(ctx, input.TestId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	//testContent, err := u.repository.FindContentByTestId(ctx, test.Id)
	//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	zap.S().Error(err)
	//	return err
	//}
	//if testContent != nil {
	//	zap.S().Error(err)
	//	return errpkg.Test.TestHasContent
	//}

	ctx = u.repository.BeginTransaction(ctx)

	multipleChoice := entities.TestMultipleChoice{
		FilePath: input.MultipleChoice.FilePath,
		Score:    input.MultipleChoice.Score,
	}
	err = u.repository.CreateTestMultipleChoice(ctx, &multipleChoice)
	if err != nil {
		u.repository.Rollback(ctx)
		log.Error().Err(err).Send()
		return err
	}

	answers := make([]entities.TestMultipleChoiceAnswer, len(input.MultipleChoice.Answers))
	for index, item := range input.MultipleChoice.Answers {
		answers[index] = entities.TestMultipleChoiceAnswer{
			Answer:               item.Answer,
			Score:                item.Score,
			Type:                 int(item.Type),
			TestMultipleChoiceId: multipleChoice.Id,
		}
	}
	err = u.repository.CreateListTestMultipleChoiceAnswers(ctx, answers)
	if err != nil {
		u.repository.Rollback(ctx)
		log.Error().Err(err).Send()
		return err
	}

	content := entities.TestContent{
		TypeableId: multipleChoice.Id,
		Typeable:   1,
		TestId:     test.Id,
	}
	err = u.repository.CreateTestContent(ctx, &content)
	if err != nil {
		u.repository.Rollback(ctx)
		log.Error().Err(err).Send()
		return err
	}

	u.repository.Commit(ctx)

	return nil
}

func (u *usecase) CheckTestCanDo(ctx context.Context, testId int) error {

	test, err := u.repository.FindByTestId(ctx, testId)
	if err != nil {
		return err
	}
	if test.TimeEnd != nil {
		if test.TimeEnd.Before(time.Now()) {
			return errors.New("time not valid")
		}
	}
	if test.TimeStart != nil {
		if test.TimeStart.After(time.Now()) {
			return errors.New("time not valid")
		}
	}

	return nil
}

func (u *usecase) CheckTestAndQuestionValid(ctx context.Context, testId int, questionId int) error {

	content, err := u.repository.FindContentByTestId(ctx, testId)
	if err != nil {
		return err
	}
	if content.Typeable == constant.MULTIPLE_CHOICE {

		questions, err := u.repository.FindTestMultipleChoiceAnswer(ctx, content.TypeableId)

		if err != nil {
			return err
		}
		valid := false
		for _, item := range questions {
			if item.Id == questionId {
				valid = true
				break
			}
		}
		if !valid {
			return errors.New("question not valid")
		}
	}

	return nil
}
