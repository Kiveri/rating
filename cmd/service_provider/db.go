package service_provider

import (
	"context"
	"log"

	"github.com/Kiveri/rating/internal/config"
)

func (sp *ServiceProvider) getDbCluster(ctx context.Context) *config.Cluster {
	if sp.dbCluster == nil {
		dbCluster, err := config.NewCluster(ctx)
		if err != nil {
			log.Fatal(err)
		}

		sp.dbCluster = dbCluster
	}

	return sp.dbCluster
}
