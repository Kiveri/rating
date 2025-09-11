package service_provider

import (
	"github.com/Kiveri/rating/internal/adapter/storage/in_memory/client_in_memory"
	"github.com/Kiveri/rating/internal/adapter/storage/in_memory/feedback_in_memory"
	"github.com/Kiveri/rating/internal/adapter/storage/in_memory/product_in_memory"
	"github.com/Kiveri/rating/internal/adapter/storage/postgres/client_db"
	"github.com/Kiveri/rating/internal/adapter/storage/postgres/feedback_db"
	"github.com/Kiveri/rating/internal/adapter/storage/postgres/product_db"
	"github.com/Kiveri/rating/internal/config"
	"github.com/Kiveri/rating/internal/controllers/http_controller/client_http_controller"
	"github.com/Kiveri/rating/internal/controllers/http_controller/feedback_http_controller"
	"github.com/Kiveri/rating/internal/controllers/http_controller/product_http_controller"
	"github.com/Kiveri/rating/internal/pkg/timer"
	"github.com/Kiveri/rating/internal/usecases/client_usecase"
	"github.com/Kiveri/rating/internal/usecases/feedback_usecase"
	"github.com/Kiveri/rating/internal/usecases/product_usecase"
)

type ServiceProvider struct {
	dbCluster *config.Cluster
	config    *config.Config
	timer     *timer.Timer

	clientRepoInMemory   *client_in_memory.Repo
	feedbackRepoInMemory *feedback_in_memory.Repo
	productRepoInMemory  *product_in_memory.Repo

	clientRepoDb   *client_db.Repo
	feedbackRepoDb *feedback_db.Repo
	productRepoDb  *product_db.Repo

	clientUseCase   *client_usecase.UseCase
	feedbackUseCase *feedback_usecase.UseCase
	productUseCase  *product_usecase.UseCase

	clientHttpController   *client_http_controller.Controller
	feedbackHttpController *feedback_http_controller.Controller
	productHttpController  *product_http_controller.Controller
}

func NewServiceProvider(config *config.Config) *ServiceProvider {
	return &ServiceProvider{
		config: config,
	}
}
