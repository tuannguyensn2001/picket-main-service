package answersheet_usecase

type IRepository interface {
}

type usecase struct {
	repository IRepository
}

func New(repository IRepository) *usecase {
	return &usecase{repository: repository}
}
