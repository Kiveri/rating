package client_in_memory

import (
	"fmt"

	"github.com/Kiveri/rating/internal/domain/model"
)

func (r *Repo) GetByID(id int64) (*model.Client, error) {
	client, exist := r.clients[id]
	if !exist {
		return nil, fmt.Errorf("%w with id = %d", errClientNotFound, id)
	}

	return client, nil
}
