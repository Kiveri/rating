package client_in_memory

import (
	"errors"
	"time"

	"github.com/Kiveri/rating/internal/domain/model"
)

var errClientNotFound = errors.New("client not found")

type (
	Repo struct {
		clients map[int64]*model.Client
		nextID  int64
		timer   timer
	}

	timer interface {
		NowMoscow() time.Time
	}
)

func NewRepo(timer timer) *Repo {
	return &Repo{
		clients: make(map[int64]*model.Client),
		nextID:  1,
		timer:   timer,
	}
}

func (r *Repo) getNextID() int64 {
	nextID := r.nextID
	r.nextID++

	return nextID
}
