package feedback_http_controller

import (
	"time"

	"github.com/Kiveri/rating/internal/domain/model"
	"github.com/samber/lo"
)

type feedbackHttpResponse struct {
	ID        int64      `json:"id"`
	Rate      uint8      `json:"rate"`
	Comment   *string    `json:"comment"`
	ProductID int64      `json:"product_id"`
	ClientID  int64      `json:"client_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func mapFeedbackToHttpResponse(feedback *model.Feedback) *feedbackHttpResponse {
	return &feedbackHttpResponse{
		ID:        feedback.ID,
		Rate:      feedback.Rate,
		Comment:   feedback.Comment,
		ProductID: feedback.ProductID,
		ClientID:  feedback.ClientID,
		CreatedAt: feedback.CreatedAt,
		UpdatedAt: feedback.UpdatedAt,
		DeletedAt: feedback.DeletedAt,
	}
}

func mapFeedbacksListHttpResponse(feedbacks []*model.Feedback) []*feedbackHttpResponse {
	return lo.Map(feedbacks, func(feedback *model.Feedback, _ int) *feedbackHttpResponse {
		return &feedbackHttpResponse{
			ID:        feedback.ID,
			Rate:      feedback.Rate,
			Comment:   feedback.Comment,
			ProductID: feedback.ProductID,
			ClientID:  feedback.ClientID,
			CreatedAt: feedback.CreatedAt,
			UpdatedAt: feedback.UpdatedAt,
			DeletedAt: feedback.DeletedAt,
		}
	})
}
