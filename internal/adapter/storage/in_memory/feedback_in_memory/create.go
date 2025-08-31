package feedback_in_memory

import (
	"github.com/Kiveri/rating/internal/domain/model"
)

func (r *Repo) Create(feedback *model.Feedback) error {
	if _, exist := r.feedbacks[feedback.ID]; exist {
		return nil
	}

	feedback.ID = r.getNextID()
	feedback.CreatedAt = r.timer.NowMoscow()
	feedback.UpdatedAt = r.timer.NowMoscow()

	r.feedbacks[feedback.ID] = feedback

	return nil
}
