package product_usecase

import (
	"fmt"

	"github.com/Kiveri/rating/internal/domain/model"
)

type CreateProductReq struct {
	Name        string
	Article     string
	Price       int64
	ProductType model.ProductType
}

func (u *UseCase) CreateProduct(req CreateProductReq) error {
	product := model.NewProduct(req.Name, req.Article, req.Price, req.ProductType)
	err := u.productRepo.Create(product)
	if err != nil {
		return fmt.Errorf("productRepo.Create: %w", err)
	}

	return nil
}
