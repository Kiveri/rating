package product_in_memory

import "github.com/Kiveri/rating/internal/domain/model"

func (r *Repo) Create(product *model.Product) error {
	if _, exist := r.products[product.ID]; exist {
		return nil
	}

	product.ID = r.getNextID()
	product.CreatedAt = r.timer.NowMoscow()
	product.UpdatedAt = r.timer.NowMoscow()

	r.products[product.ID] = product

	return nil
}
