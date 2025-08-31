package feedback_usecase

import "fmt"

type EditRateAndCommentReq struct {
	ID      int64
	Rate    uint8
	Comment *string
}

func (u *UseCase) EditRateAndComment(req EditRateAndCommentReq) error {
	feedback, err := u.feedbackRepo.GetByID(req.ID)
	if err != nil {
		return fmt.Errorf("feedbackRepo.GetByID: %w", err)
	}
	feedback.Rate = req.Rate
	feedback.Comment = req.Comment

	err = u.feedbackRepo.Update(feedback)
	if err != nil {
		return fmt.Errorf("feedbackRepo.Update: %w", err)
	}

	product, err := u.productRepo.GetByID(feedback.ProductID)
	if err != nil {
		return fmt.Errorf("productRepo.GetByID: %w", err)
	}
	product.CalculateAverageRateByDeleteOld(feedback.Rate, product.FeedbacksCount)
	product.CalculateAverageRateByAddNew(req.Rate, product.FeedbacksCount)

	err = u.productRepo.Update(product)
	if err != nil {
		return fmt.Errorf("productRepo.Update: %w", err)
	}

	return nil
}
