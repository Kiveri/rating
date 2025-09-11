package feedback_in_memory

import (
	"fmt"

	"github.com/Kiveri/rating/internal/adapter/storage"

	"github.com/Kiveri/rating/internal/domain/model"
)

func (r *Repo) GetByID(id int64) (*model.Feedback, error) {
	feedback, exists := r.feedbacks[id]
	if !exists {
		return nil, fmt.Errorf("%w with id = %d", storage.ErrFeedbackNotFound, id)
	}

	return feedback, nil
}
