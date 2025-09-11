package filters

type FeedbackFindAllFilter struct {
	ID        int64
	Rate      uint8
	ProductID int64
	ClientID  int64
}
