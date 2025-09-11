package product_in_memory

import (
	"fmt"

	"github.com/Kiveri/rating/internal/adapter/storage"

	"github.com/Kiveri/rating/internal/domain/model"
)

func (r *Repo) GetByID(id int64) (*model.Product, error) {
	product, exist := r.products[id]
	if !exist {
		return nil, fmt.Errorf("%w with id = %d", storage.ErrProductNotFound, id)
	}

	return product, nil
}
