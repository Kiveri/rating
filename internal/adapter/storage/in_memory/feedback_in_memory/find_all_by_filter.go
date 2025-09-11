package feedback_in_memory

import (
	"github.com/Kiveri/rating/internal/domain/dto/filters"
	"github.com/Kiveri/rating/internal/domain/model"
)

func (r *Repo) FindAllByFilter(filter *filters.FeedbackFindAllFilter) ([]*model.Feedback, error) {
	var result []*model.Feedback

	if filter == nil {
		for _, feedback := range r.feedbacks {
			if feedback.DeletedAt == nil {
				result = append(result, feedback)
			}
		}

		return result, nil
	}

	for _, feedback := range r.feedbacks {
		if feedback.DeletedAt != nil {
			continue
		}

		match := true

		if filter.ID > 0 && filter.ID != feedback.ID {
			match = false
		}
		if filter.Rate > 0 && feedback.Rate != filter.Rate {
			match = false
		}
		if filter.ProductID > 0 && feedback.ProductID != filter.ProductID {
			match = false
		}
		if filter.ClientID > 0 && feedback.ClientID != filter.ClientID {
			match = false
		}
		if match {
			result = append(result, feedback)
		}
	}

	return result, nil
}
