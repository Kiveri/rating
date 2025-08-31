package client_usecase

import (
	"fmt"

	"github.com/Kiveri/rating/internal/domain/model"
)

type CreateClientReq struct {
	Fio   string
	Phone string
	Email string
}

func (u *UseCase) Create(req CreateClientReq) error {
	client := model.NewClient(req.Email, req.Fio, req.Phone)
	err := u.clientRepo.Create(client)
	if err != nil {
		return fmt.Errorf("clientRepo.Create: %w", err)
	}

	return nil
}
