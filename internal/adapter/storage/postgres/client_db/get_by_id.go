package client_db

import "github.com/Kiveri/rating/internal/domain/model"

func (r *Repo) GetByID(id int64) (*model.Client, error) {
	return &model.Client{}, nil
}
