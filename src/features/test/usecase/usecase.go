package test_usecase

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"picket-main-service/src/app"
	"picket-main-service/src/base"
	"picket-main-service/src/constant"
	"picket-main-service/src/dto"
	"picket-main-service/src/entities"
	randompkg "picket-main-service/src/pkg/random"
	"picket-main-service/src/utils"
	"strings"
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
}

func New(repository IRepository) *usecase {
	return &usecase{repository: repository}
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

func (u *usecase) GetById(ctx context.Context, id int) (*entities.Test, error) {
	return u.repository.FindByTestId(ctx, id)
}

func (u *usecase) GetContent(ctx context.Context, testId int) (*entities.TestContent, error) {

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
