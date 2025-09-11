package product_usecase

import (
	"fmt"

	"github.com/Kiveri/rating/internal/domain/model"
)

type GetProductReq struct {
	ID int64
}

func (u *UseCase) GetByID(req GetProductReq) (*model.Product, error) {
	product, err := u.productRepo.GetByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("productRepo.GetByID: %w", err)
	}

	return product, nil
}
