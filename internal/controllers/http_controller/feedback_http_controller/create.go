package feedback_http_controller

import (
	"encoding/json"
	"net/http"

	"github.com/Kiveri/rating/internal/controllers/http_controller"
	"github.com/Kiveri/rating/internal/usecases/feedback_usecase"
)

type createFeedbackReq struct {
	Rate      uint8   `json:"rate"`
	Comment   *string `json:"comment"`
	ProductID int64   `json:"product_id"`
	ClientID  int64   `json:"client_id"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req createFeedbackReq
	err := decoder.Decode(&req)
	if err != nil {
		http_controller.InternalServerErrorRespond(w, err)

		return
	}

	if validationError := validateCreateFeedback(req); validationError != nil {
		http_controller.ValidationErrorRespond(w, validationError)

		return
	}

	err = c.feedbackUseCase.CreateFeedback(feedback_usecase.CreateFeedbackReq{
		Rate:      req.Rate,
		Comment:   req.Comment,
		ProductID: req.ProductID,
		ClientID:  req.ClientID,
	})
	if err != nil {
		http_controller.InternalServerErrorRespond(w, err)

		return
	}
}

func validateCreateFeedback(req createFeedbackReq) *http_controller.ValidationError {
	if req.Rate > 5 {
		return http_controller.NewValidationError("rate must be from 0 to 5", "rate")
	}
	if req.ProductID <= 0 {
		return http_controller.NewValidationError("product id must be greater than zero", "product_id")
	}
	if req.ClientID <= 0 {
		return http_controller.NewValidationError("client id must be greater than zero", "client_id")
	}

	return nil
}
