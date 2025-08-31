package product_usecase

import "github.com/Kiveri/rating/internal/domain/model"

type (
	UseCase struct {
		productRepo productRepo
	}

	productRepo interface {
		Create(product *model.Product) error
	}
)

func NewUseCase(productRepo productRepo) *UseCase {
	return &UseCase{
		productRepo: productRepo,
	}
}
