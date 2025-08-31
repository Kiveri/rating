package feedback_in_memory

import (
	"fmt"

	"github.com/samber/lo"
)

func (r *Repo) Delete(id int64) error {
	feedback, exist := r.feedbacks[id]
	if !exist {
		return fmt.Errorf("%w with id = %d", errFeedbackNotFound, id)
	}

	feedback.DeletedAt = lo.ToPtr(r.timer.NowMoscow())

	r.feedbacks[id] = feedback

	return nil
}
