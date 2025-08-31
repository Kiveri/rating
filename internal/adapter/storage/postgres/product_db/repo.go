package product_db

import (
	"time"

	"github.com/Kiveri/rating/internal/config"
)

type (
	Repo struct {
		cluster *config.Cluster
		timer   timer
	}

	timer interface {
		NowMoscow() time.Time
	}
)

func NewRepo(cluster *config.Cluster, timer timer) *Repo {
	return &Repo{
		cluster: cluster,
		timer:   timer,
	}
}
