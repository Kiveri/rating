package feedback_http_controller

import (
	"net/http"
	"strconv"

	"github.com/Kiveri/rating/internal/controllers/http_controller"
	"github.com/Kiveri/rating/internal/usecases/feedback_usecase"
)

func (c *Controller) FindAllByProductID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http_controller.ValidationErrorRespond(w, http_controller.NewValidationError("product id not present", "id"))

		return
	}

	if validationError := validateFindAllByProductID(productID); validationError != nil {
		http_controller.ValidationErrorRespond(w, validationError)

		return
	}

	feedbacks, err := c.feedbackUseCase.FindAllByProductId(feedback_usecase.FindAllByProductIDReq{
		ProductID: productID,
	})
	if err != nil {
		http_controller.InternalServerErrorRespond(w, err)

		return
	}

	if err = http_controller.EncodeResponse(w, mapFeedbacksListHttpResponse(feedbacks)); err != nil {
		return
	}
}

func validateFindAllByProductID(productID int64) *http_controller.ValidationError {
	if productID <= 0 {
		return http_controller.NewValidationError("product id must be positive", "product_id")
	}

	return nil
}
