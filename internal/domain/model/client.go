package model

import "time"

type Client struct {
	ID        int64
	Fio       string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(fio, phone, email string) *Client {
	return &Client{
		Fio:   fio,
		Phone: phone,
		Email: email,
	}
}
