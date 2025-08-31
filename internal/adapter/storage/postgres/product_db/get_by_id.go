package product_db

import "github.com/Kiveri/rating/internal/domain/model"

func (r *Repo) GetByID(id int64) (*model.Product, error) {
	return &model.Product{}, nil
}
