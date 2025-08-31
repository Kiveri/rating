package feedback_db

import "github.com/Kiveri/rating/internal/domain/model"

func (r *Repo) GetByID(id int64) (*model.Feedback, error) {
	return &model.Feedback{}, nil
}
