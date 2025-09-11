package client_http_controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Kiveri/rating/internal/adapter/storage"
	"github.com/Kiveri/rating/internal/controllers/http_controller"
	"github.com/Kiveri/rating/internal/usecases/client_usecase"
)

func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	clientID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http_controller.ValidationErrorRespond(w, http_controller.NewValidationError("client id not present", "id"))

		return
	}

	client, err := c.clientUseCase.GetByID(client_usecase.GetClientReq{
		ID: clientID,
	})
	if err != nil {
		if errors.Is(err, storage.ErrClientNotFound) {
			http_controller.NotFoundErrorRespond(w, http_controller.NewNotFoundError("client not found"))

			return
		}

		http_controller.InternalServerErrorRespond(w, err)

		return
	}

	if err = http_controller.EncodeResponse(w, mapClientToHttpResponse(client)); err != nil {
		return
	}
}
