package product_http_controller

import (
	"encoding/json"
	"net/http"

	"github.com/Kiveri/rating/internal/controllers/http_controller"
	"github.com/Kiveri/rating/internal/domain/model"
	"github.com/Kiveri/rating/internal/usecases/product_usecase"
)

type createProductReq struct {
	Name        string `json:"name"`
	Article     string `json:"article"`
	Price       int64  `json:"price"`
	ProductType string `json:"product_type"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req createProductReq
	err := decoder.Decode(&req)
	if err != nil {
		http_controller.InternalServerErrorRespond(w, err)

		return
	}

	if validationError := validateCreateProductRequest(req); validationError != nil {
		http_controller.ValidationErrorRespond(w, validationError)

		return
	}

	var productType model.ProductType
	switch req.ProductType {
	case "raw":
		productType = model.PRODUCT_TYPE_RAW
	case "prod":
		productType = model.PRODUCT_TYPE_PROD
	default:
		productType = 0
	}

	err = c.productUseCase.CreateProduct(product_usecase.CreateProductReq{
		Name:        req.Name,
		Article:     req.Article,
		Price:       req.Price,
		ProductType: productType,
	})
	if err != nil {
		http_controller.InternalServerErrorRespond(w, err)

		return
	}
}

func validateCreateProductRequest(req createProductReq) *http_controller.ValidationError {
	if req.Name == "" {
		return http_controller.NewValidationError("name is required", "name")
	}
	if req.Article == "" {
		return http_controller.NewValidationError("article is required", "article")
	}
	if req.Price <= 0 {
		return http_controller.NewValidationError("price must be greater than zero", "price")
	}
	if req.ProductType != "raw" && req.ProductType != "prod" {
		return http_controller.NewValidationError("product_type must be only 'raw' or 'prod'", "product_type")
	}

	return nil
}
