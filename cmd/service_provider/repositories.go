package service_provider

import (
	"context"

	"github.com/Kiveri/rating/internal/adapter/storage/in_memory/client_in_memory"
	"github.com/Kiveri/rating/internal/adapter/storage/in_memory/feedback_in_memory"
	"github.com/Kiveri/rating/internal/adapter/storage/in_memory/product_in_memory"
	"github.com/Kiveri/rating/internal/adapter/storage/postgres/client_db"
	"github.com/Kiveri/rating/internal/adapter/storage/postgres/feedback_db"
	"github.com/Kiveri/rating/internal/adapter/storage/postgres/product_db"
)

func (sp *ServiceProvider) getClientRepoInMemory() *client_in_memory.Repo {
	if sp.clientRepoInMemory == nil {
		sp.clientRepoInMemory = client_in_memory.NewRepo(sp.getTimer())
	}

	return sp.clientRepoInMemory
}

func (sp *ServiceProvider) getFeedbackRepoInMemory() *feedback_in_memory.Repo {
	if sp.feedbackRepoInMemory == nil {
		sp.feedbackRepoInMemory = feedback_in_memory.NewRepo(sp.getTimer())
	}

	return sp.feedbackRepoInMemory
}

func (sp *ServiceProvider) getProductRepoInMemory() *product_in_memory.Repo {
	if sp.productRepoInMemory == nil {
		sp.productRepoInMemory = product_in_memory.NewRepo(sp.getTimer())
	}

	return sp.productRepoInMemory
}

func (sp *ServiceProvider) getClientRepoDb() *client_db.Repo {
	if sp.clientRepoDb == nil {
		sp.clientRepoDb = client_db.NewRepo(
			sp.getDbCluster(context.Background()),
			sp.getTimer(),
		)
	}

	return sp.clientRepoDb
}

func (sp *ServiceProvider) getFeedbackRepoDb() *feedback_db.Repo {
	if sp.feedbackRepoDb == nil {
		sp.feedbackRepoDb = feedback_db.NewRepo(
			sp.getDbCluster(context.Background()),
			sp.getTimer(),
		)
	}

	return sp.feedbackRepoDb
}

func (sp *ServiceProvider) getProductRepoDb() *product_db.Repo {
	if sp.productRepoDb == nil {
		sp.productRepoDb = product_db.NewRepo(
			sp.getDbCluster(context.Background()),
			sp.getTimer(),
		)
	}

	return sp.productRepoDb
}
