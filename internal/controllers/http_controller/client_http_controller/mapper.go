package client_http_controller

import (
	"time"

	"github.com/Kiveri/rating/internal/domain/model"
)

type clientHttpResponse struct {
	ID        int64     `json:"id"`
	Fio       string    `json:"fio"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func mapClientToHttpResponse(client *model.Client) *clientHttpResponse {
	return &clientHttpResponse{
		ID:        client.ID,
		Fio:       client.Fio,
		Phone:     client.Phone,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}
}
