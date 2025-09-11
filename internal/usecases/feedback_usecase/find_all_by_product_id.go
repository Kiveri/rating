package feedback_usecase

import (
	"fmt"

	"github.com/Kiveri/rating/internal/domain/dto/filters"
	"github.com/Kiveri/rating/internal/domain/model"
)

type FindAllByProductIDReq struct {
	ProductID int64
}

func (u *UseCase) FindAllByProductId(req FindAllByProductIDReq) ([]*model.Feedback, error) {
	feedbacks, err := u.feedbackRepo.FindAllByFilter(&filters.FeedbackFindAllFilter{
		ProductID: req.ProductID,
	})
	if err != nil {
		return nil, fmt.Errorf("feedbackRepo.FindAllByFilter: %w by product id = %d", err, req.ProductID)
	}

	return feedbacks, nil
}
