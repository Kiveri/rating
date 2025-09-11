package client_http_controller

import (
	"encoding/json"
	"net/http"

	"github.com/Kiveri/rating/internal/controllers/http_controller"
	"github.com/Kiveri/rating/internal/usecases/client_usecase"
)

type createClientReq struct {
	Fio   string `json:"fio"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req createClientReq
	err := decoder.Decode(&req)
	if err != nil {
		http_controller.InternalServerErrorRespond(w, err)

		return
	}

	if validationError := validateCreateClientRequest(req); validationError != nil {
		http_controller.ValidationErrorRespond(w, validationError)

		return
	}

	err = c.clientUseCase.Create(client_usecase.CreateClientReq{
		Fio:   req.Fio,
		Phone: req.Phone,
		Email: req.Email,
	})
	if err != nil {
		http_controller.InternalServerErrorRespond(w, err)

		return
	}
}

func validateCreateClientRequest(req createClientReq) *http_controller.ValidationError {
	if req.Fio == "" {
		return http_controller.NewValidationError("fio is required", "fio")
	}
	if req.Phone == "" {
		return http_controller.NewValidationError("phone is required", "phone")
	}
	if req.Email == "" {
		return http_controller.NewValidationError("email is required", "email")
	}

	return nil
}
