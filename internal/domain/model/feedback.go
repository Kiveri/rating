package model

import "time"

type Feedback struct {
	ID        int64
	Rate      uint8
	Comment   *string
	ProductID int64
	ClientID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewFeedback(rate uint8, comment *string, productID int64, clientID int64) *Feedback {
	return &Feedback{
		Rate:      rate,
		Comment:   comment,
		ProductID: productID,
		ClientID:  clientID,
	}
}
