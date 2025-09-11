package product_http_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Kiveri/rating/internal/adapter/storage"
	"github.com/Kiveri/rating/internal/controllers/http_controller"
	"github.com/Kiveri/rating/internal/usecases/product_usecase"
)

func (c *Controller) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http_controller.ValidationErrorRespond(w, http_controller.NewValidationError("product id not present", "id"))

		return
	}

	product, err := c.productUseCase.GetByID(product_usecase.GetProductReq{
		ID: productID,
	})
	if err != nil {
		if errors.Is(err, storage.ErrProductNotFound) {
			http_controller.NotFoundErrorRespond(w, http_controller.NewNotFoundError("product not found"))

			return
		}

		http_controller.InternalServerErrorRespond(w, err)

		return
	}

	if err = http_controller.EncodeResponse(w, mapProductToHttpResponse(product)); err != nil {
		return
	}
}
