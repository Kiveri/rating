package feedback_http_controller

import (
	"github.com/Kiveri/rating/internal/domain/model"
	"github.com/Kiveri/rating/internal/usecases/feedback_usecase"
)

type (
	Controller struct {
		feedbackUseCase feedbackUseCase
	}

	feedbackUseCase interface {
		CreateFeedback(req feedback_usecase.CreateFeedbackReq) error
		GetByID(req feedback_usecase.GetFeedbackReq) (*model.Feedback, error)
		FindAllByProductId(req feedback_usecase.FindAllByProductIDReq) ([]*model.Feedback, error)
	}
)

func NewController(feedbackUseCase feedbackUseCase) *Controller {
	return &Controller{
		feedbackUseCase: feedbackUseCase,
	}
}
