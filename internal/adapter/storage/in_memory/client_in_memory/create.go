package client_in_memory

import (
	"github.com/Kiveri/rating/internal/domain/model"
)

func (r *Repo) Create(client *model.Client) error {
	if _, exist := r.clients[client.ID]; exist {
		return nil
	}

	client.ID = r.getNextID()
	client.CreatedAt = r.timer.NowMoscow()
	client.UpdatedAt = r.timer.NowMoscow()

	r.clients[client.ID] = client

	return nil
}
