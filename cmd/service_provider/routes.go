package service_provider

import "net/http"

func (sp *ServiceProvider) GetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /clients/create", sp.GetClientHttpController().Create)
	mux.HandleFunc("GET /clients/{id}", sp.GetClientHttpController().GetByID)

	mux.HandleFunc("POST /feedbacks/create", sp.GetFeedbackHttpController().Create)
	mux.HandleFunc("GET /feedbacks/{id}", sp.GetFeedbackHttpController().GetById)
	mux.HandleFunc("GET /feedbacks/product/{id}", sp.GetFeedbackHttpController().FindAllByProductID)

	mux.HandleFunc("POST /products/create", sp.GetProductHttpController().Create)
	mux.HandleFunc("GET /products/{id}", sp.GetProductHttpController().GetById)

	return mux
}
