package feedback_in_memory

import (
	"fmt"

	"github.com/Kiveri/rating/internal/adapter/storage"

	"github.com/samber/lo"
)

func (r *Repo) Delete(id int64) error {
	feedback, exist := r.feedbacks[id]
	if !exist {
		return fmt.Errorf("%w with id = %d", storage.ErrFeedbackNotFound, id)
	}

	feedback.DeletedAt = lo.ToPtr(r.timer.NowMoscow())

	r.feedbacks[id] = feedback

	return nil
}
