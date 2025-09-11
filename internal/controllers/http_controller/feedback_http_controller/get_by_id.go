package feedback_http_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Kiveri/rating/internal/adapter/storage"
	"github.com/Kiveri/rating/internal/controllers/http_controller"
	"github.com/Kiveri/rating/internal/usecases/feedback_usecase"
)

func (c *Controller) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	feedbackID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http_controller.ValidationErrorRespond(w, http_controller.NewValidationError("feedback id not present", "id"))

		return
	}

	feedback, err := c.feedbackUseCase.GetByID(feedback_usecase.GetFeedbackReq{
		ID: feedbackID,
	})
	if err != nil {
		if errors.Is(err, storage.ErrFeedbackNotFound) {
			http_controller.NotFoundErrorRespond(w, http_controller.NewNotFoundError("feedback not found"))

			return
		}

		http_controller.InternalServerErrorRespond(w, err)

		return
	}

	if err = http_controller.EncodeResponse(w, mapFeedbackToHttpResponse(feedback)); err != nil {
		return
	}
}
