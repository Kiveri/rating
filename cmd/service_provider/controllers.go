package service_provider

import (
	"github.com/Kiveri/rating/internal/controllers/http_controller/client_http_controller"
	"github.com/Kiveri/rating/internal/controllers/http_controller/feedback_http_controller"
	"github.com/Kiveri/rating/internal/controllers/http_controller/product_http_controller"
)

func (sp *ServiceProvider) GetClientHttpController() *client_http_controller.Controller {
	if sp.clientHttpController == nil {
		sp.clientHttpController = client_http_controller.NewController(sp.getClientUseCase())
	}

	return sp.clientHttpController
}

func (sp *ServiceProvider) GetFeedbackHttpController() *feedback_http_controller.Controller {
	if sp.feedbackHttpController == nil {
		sp.feedbackHttpController = feedback_http_controller.NewController(sp.getFeedbackUseCase())
	}

	return sp.feedbackHttpController
}

func (sp *ServiceProvider) GetProductHttpController() *product_http_controller.Controller {
	if sp.productHttpController == nil {
		sp.productHttpController = product_http_controller.NewController(sp.getProductUseCase())
	}

	return sp.productHttpController
}
