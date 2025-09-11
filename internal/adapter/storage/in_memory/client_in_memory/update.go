package client_in_memory

import (
	"fmt"

	"github.com/Kiveri/rating/internal/adapter/storage"

	"github.com/Kiveri/rating/internal/domain/model"
)

func (r *Repo) Update(client *model.Client) error {
	if _, exist := r.clients[client.ID]; !exist {
		return fmt.Errorf("%w with id = %d", storage.ErrClientNotFound, client.ID)
	}

	client.UpdatedAt = r.timer.NowMoscow()

	r.clients[client.ID] = client

	return nil
}
