package feedback_usecase

import (
	"github.com/Kiveri/rating/internal/domain/dto/filters"
	"github.com/Kiveri/rating/internal/domain/model"
)

type (
	UseCase struct {
		feedbackRepo feedbackRepo
		productRepo  productRepo
	}

	feedbackRepo interface {
		Create(feedback *model.Feedback) error
		Delete(id int64) error
		GetByID(id int64) (*model.Feedback, error)
		Update(feedback *model.Feedback) error
		FindAllByFilter(filter *filters.FeedbackFindAllFilter) ([]*model.Feedback, error)
	}
	productRepo interface {
		GetByID(id int64) (*model.Product, error)
		Update(product *model.Product) error
	}
)

func NewUseCase(feedbackRepo feedbackRepo, productRepo productRepo) *UseCase {
	return &UseCase{
		feedbackRepo: feedbackRepo,
		productRepo:  productRepo,
	}
}
