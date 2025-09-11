package product_http_controller

import (
	"github.com/Kiveri/rating/internal/domain/model"
	"github.com/Kiveri/rating/internal/usecases/product_usecase"
)

type (
	Controller struct {
		productUseCase productUseCase
	}

	productUseCase interface {
		CreateProduct(req product_usecase.CreateProductReq) error
		GetByID(req product_usecase.GetProductReq) (*model.Product, error)
	}
)

func NewController(productUseCase productUseCase) *Controller {
	return &Controller{
		productUseCase: productUseCase,
	}
}
