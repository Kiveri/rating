package feedback_usecase

import "fmt"

type DeleteFeedbackReq struct {
	ID int64
}

func (u *UseCase) Delete(req DeleteFeedbackReq) error {
	feedBack, err := u.feedbackRepo.GetByID(req.ID)
	if err != nil {
		return fmt.Errorf("feedbackRepo.GetByID: %w", err)
	}

	err = u.feedbackRepo.Delete(req.ID)
	if err != nil {
		return fmt.Errorf("feedbackRepo.Delete: %w", err)
	}

	product, err := u.productRepo.GetByID(feedBack.ProductID)
	if err != nil {
		return fmt.Errorf("productRepo.GetByID: %w", err)
	}

	product.CalculateAverageRateByDeleteOld(feedBack.Rate, product.FeedbacksCount)

	err = u.productRepo.Update(product)
	if err != nil {
		return fmt.Errorf("productRepo.Update: %w", err)
	}

	return nil
}
