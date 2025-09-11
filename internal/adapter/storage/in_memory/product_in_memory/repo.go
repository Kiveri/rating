package product_in_memory

import (
	"time"

	"github.com/Kiveri/rating/internal/domain/model"
)

type (
	Repo struct {
		products map[int64]*model.Product
		nextID   int64
		timer    timer
	}

	timer interface {
		NowMoscow() time.Time
	}
)

func NewRepo(timer timer) *Repo {
	return &Repo{
		products: make(map[int64]*model.Product),
		nextID:   1,
		timer:    timer,
	}
}

func (r *Repo) getNextID() int64 {
	nextID := r.nextID
	r.nextID++

	return nextID
}
