package feedback_usecase

import (
	"fmt"

	"github.com/Kiveri/rating/internal/domain/model"
)

type CreateFeedbackReq struct {
	Rate      uint8
	Comment   *string
	ProductID int64
	ClientID  int64
}

func (u *UseCase) CreateFeedback(req CreateFeedbackReq) error {
	product, err := u.productRepo.GetByID(req.ProductID)
	if err != nil {
		return fmt.Errorf("productRepo.GetByID: %w", err)
	}
	product.CalculateAverageRateByAddNew(req.Rate, product.FeedbacksCount)

	feedback := model.NewFeedback(req.Rate, req.Comment, req.ProductID, req.ClientID)
	err = u.feedbackRepo.Create(feedback)
	if err != nil {
		return fmt.Errorf("feedbackRepo.Create: %w", err)
	}

	err = u.productRepo.Update(product)
	if err != nil {
		return fmt.Errorf("productRepo.Update: %w", err)
	}

	return nil
}
