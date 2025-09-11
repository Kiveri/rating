package client_usecase

import "github.com/Kiveri/rating/internal/domain/model"

type (
	UseCase struct {
		clientRepo clientRepo
	}

	clientRepo interface {
		Create(client *model.Client) error
		GetByID(id int64) (*model.Client, error)
	}
)

func NewUseCase(clientRepo clientRepo) *UseCase {
	return &UseCase{
		clientRepo: clientRepo,
	}
}
