package product_in_memory

import (
	"fmt"

	"github.com/Kiveri/rating/internal/domain/model"
)

func (r *Repo) Update(product *model.Product) error {
	if _, exist := r.products[product.ID]; !exist {
		return fmt.Errorf("%w with id = %d", errProductNotFound, product.ID)
	}

	product.UpdatedAt = r.timer.NowMoscow()

	r.products[product.ID] = product

	return nil
}
