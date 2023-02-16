package test_repository

import (
	"context"
	"gorm.io/gorm"
	"picket-main-service/src/base"
	"picket-main-service/src/constant"
	"picket-main-service/src/entities"
)

type repo struct {
	base.Repository
}

func New(db *gorm.DB) *repo {
	return &repo{
		Repository: base.Repository{
			Db: db,
		},
	}
}

func (r *repo) CreateListTestMultipleChoiceAnswers(ctx context.Context, list []entities.TestMultipleChoiceAnswer) error {
	db := r.GetDB(ctx)
	return db.WithContext(ctx).Create(&list).Error
}

func (r *repo) CreateTestContent(ctx context.Context, content *entities.TestContent) error {
	db := r.GetDB(ctx)
	return db.WithContext(ctx).Create(content).Error
}

func (r *repo) CreateTestMultipleChoice(ctx context.Context, test *entities.TestMultipleChoice) error {
	db := r.GetDB(ctx)
	return db.WithContext(ctx).Create(test).Error
}

func (r *repo) FindByCode(ctx context.Context, code string) (*entities.Test, error) {
	db := r.GetDB(ctx)
	var result entities.Test
	if err := db.WithContext(ctx).Where("code = ?", code).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *repo) FindContentByTestId(ctx context.Context, testId int) (*entities.TestContent, error) {
	db := r.GetDB(ctx)
	var result entities.TestContent
	if err := db.WithContext(ctx).Where("test_id = ?", testId).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repo) FindTestById(ctx context.Context, testId int) (*entities.Test, error) {
	db := r.GetDB(ctx)
	var result entities.Test
	if err := db.WithContext(ctx).Where("id = ?", testId).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repo) FindTestByUserId(ctx context.Context, userId int) ([]entities.Test, error) {
	db := r.GetDB(ctx)

	var result []entities.Test
	if err := db.WithContext(ctx).Where("created_by = ?", userId).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repo) FindTestMultipleChoiceAnswer(ctx context.Context, multipleChoiceId int) ([]entities.TestMultipleChoiceAnswer, error) {
	db := r.GetDB(ctx)
	var result []entities.TestMultipleChoiceAnswer
	if err := db.WithContext(ctx).Where("test_multiple_choice_id = ?", multipleChoiceId).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repo) FindTestMultipleChoiceByTestId(ctx context.Context, testId int) (*entities.TestMultipleChoice, error) {
	db := r.GetDB(ctx)
	var result entities.TestMultipleChoice
	if err := db.WithContext(ctx).Raw("select tmc.* from test_multiple_choice tmc join test_content tc on tmc.id = tc.typeable_id where tc.typeable = ? and tc.test_id = ? order by id desc limit 1", constant.MULTIPLE_CHOICE, testId).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repo) Create(ctx context.Context, test *entities.Test) error {
	db := r.GetDB(ctx)
	return db.WithContext(ctx).Create(test).Error
}

func (r *repo) SaveTestMultipleChoice(ctx context.Context, entity *entities.TestMultipleChoice) error {
	db := r.GetDB(ctx)
	return db.WithContext(ctx).Save(entity).Error
}
