package client_http_controller

import (
	"github.com/Kiveri/rating/internal/domain/model"
	"github.com/Kiveri/rating/internal/usecases/client_usecase"
)

type (
	Controller struct {
		clientUseCase clientUseCase
	}

	clientUseCase interface {
		Create(req client_usecase.CreateClientReq) error
		GetByID(req client_usecase.GetClientReq) (*model.Client, error)
	}
)

func NewController(clientUseCase clientUseCase) *Controller {
	return &Controller{
		clientUseCase: clientUseCase,
	}
}
