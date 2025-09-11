package model

import "time"

type ProductType uint8

const (
	PRODUCT_TYPE_RAW  ProductType = 1
	PRODUCT_TYPE_PROD ProductType = 2
)

type Product struct {
	ID             int64
	Name           string
	Article        string
	Price          int64
	ProductType    ProductType
	AverageRate    float64
	FeedbacksCount uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewProduct(name, article string, price int64, productType ProductType) *Product {
	return &Product{
		Name:        name,
		Article:     article,
		Price:       price,
		ProductType: productType,
	}
}

// CalculateAverageRateByAddNew пересчитывает среднюю оценку при добавлении отзыва.
func (p *Product) CalculateAverageRateByAddNew(newRate uint8, currentFeedBacksCount uint64) {
	p.AverageRate = (p.AverageRate*float64(p.FeedbacksCount) + float64(newRate)) /
		float64(currentFeedBacksCount+1)
	p.FeedbacksCount++
}

// CalculateAverageRateByDeleteOld пересчитывает среднюю оценку при удалении отзыва.
func (p *Product) CalculateAverageRateByDeleteOld(oldRate uint8, currentFeedBacksCount uint64) {
	if p.FeedbacksCount <= 1 {
		p.AverageRate = 0
		p.FeedbacksCount = 0
	} else {
		p.AverageRate = (p.AverageRate*float64(p.FeedbacksCount) - float64(oldRate)) /
			float64(currentFeedBacksCount-1)
		p.FeedbacksCount--
	}
}
