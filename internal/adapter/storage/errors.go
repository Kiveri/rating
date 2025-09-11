package storage

import "errors"

var (
	ErrClientNotFound   = errors.New("client not found")
	ErrFeedbackNotFound = errors.New("feedback not found")
	ErrProductNotFound  = errors.New("product not found")
)
