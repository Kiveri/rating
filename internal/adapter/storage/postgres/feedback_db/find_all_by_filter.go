package feedback_db

import (
	"github.com/Kiveri/rating/internal/domain/dto/filters"
	"github.com/Kiveri/rating/internal/domain/model"
)

func (r *Repo) FindAllByFilter(filter *filters.FeedbackFindAllFilter) ([]*model.Feedback, error) {
	return nil, nil
}
