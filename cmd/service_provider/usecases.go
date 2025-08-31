package service_provider

import (
	"github.com/Kiveri/rating/internal/usecases/client_usecase"
	"github.com/Kiveri/rating/internal/usecases/feedback_usecase"
	"github.com/Kiveri/rating/internal/usecases/product_usecase"
)

const (
	inMemory = "in-memory"
	postgres = "postgres"
)

func (sp *ServiceProvider) GetClientUseCase() *client_usecase.UseCase {
	if sp.clientUseCase == nil {
		if sp.config.StorageType == inMemory {
			sp.clientUseCase = client_usecase.NewUseCase(sp.getClientRepoInMemory())
		}
		if sp.config.StorageType == postgres {
			sp.clientUseCase = client_usecase.NewUseCase(sp.getClientRepoDb())
		}
	}

	return sp.clientUseCase
}

func (sp *ServiceProvider) GetFeedBackUseCase() *feedback_usecase.UseCase {
	if sp.feedbackUseCase == nil {
		if sp.config.StorageType == inMemory {
			sp.feedbackUseCase = feedback_usecase.NewUseCase(
				sp.getFeedbackRepoInMemory(),
				sp.getProductRepoInMemory(),
			)
		}
		if sp.config.StorageType == postgres {
			sp.feedbackUseCase = feedback_usecase.NewUseCase(
				sp.getFeedbackRepoDb(),
				sp.getProductRepoDb(),
			)
		}
	}

	return sp.feedbackUseCase
}

func (sp *ServiceProvider) GetProductUseCase() *product_usecase.UseCase {
	if sp.productUseCase == nil {
		if sp.config.StorageType == inMemory {
			sp.productUseCase = product_usecase.NewUseCase(sp.getProductRepoInMemory())
		}
		if sp.config.StorageType == postgres {
			sp.productUseCase = product_usecase.NewUseCase(sp.getProductRepoDb())
		}
	}

	return sp.productUseCase
}
