package client_usecase

import (
	"fmt"

	"github.com/Kiveri/rating/internal/domain/model"
)

type GetClientReq struct {
	ID int64
}

func (u *UseCase) GetByID(req GetClientReq) (*model.Client, error) {
	client, err := u.clientRepo.GetByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("clientRepo.GetByID: %w", err)
	}

	return client, nil
}
