package product_http_controller

import (
	"math"
	"time"

	"github.com/Kiveri/rating/internal/domain/model"
)

type productHttpResponse struct {
	ID             int64             `json:"id"`
	Name           string            `json:"name"`
	Article        string            `json:"article"`
	Price          int64             `json:"price"`
	ProductType    model.ProductType `json:"product_type"`
	AverageRate    float64           `json:"average_rate"`
	FeedbacksCount uint64            `json:"feedbacks_count"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
}

func mapProductToHttpResponse(product *model.Product) *productHttpResponse {
	return &productHttpResponse{
		ID:             product.ID,
		Name:           product.Name,
		Article:        product.Article,
		Price:          product.Price,
		ProductType:    product.ProductType,
		AverageRate:    math.Round(product.AverageRate*10) / 10,
		FeedbacksCount: product.FeedbacksCount,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
	}
}
