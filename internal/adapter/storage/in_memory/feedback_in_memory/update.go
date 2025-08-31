package feedback_in_memory

import (
	"fmt"

	"github.com/Kiveri/rating/internal/domain/model"
)

func (r *Repo) Update(feedback *model.Feedback) error {
	if _, exist := r.feedbacks[feedback.ID]; !exist {
		return fmt.Errorf("%w with id = %d", errFeedbackNotFound, feedback.ID)
	}

	feedback.UpdatedAt = r.timer.NowMoscow()

	r.feedbacks[feedback.ID] = feedback

	return nil
}
