package feedback_usecase

import (
	"fmt"

	"github.com/Kiveri/rating/internal/domain/model"
)

type GetFeedbackReq struct {
	ID int64
}

func (u *UseCase) GetByID(req GetFeedbackReq) (*model.Feedback, error) {
	feedback, err := u.feedbackRepo.GetByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("feedbackRepo.GetByID: %w", err)
	}

	return feedback, nil
}
