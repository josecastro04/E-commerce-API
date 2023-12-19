package models

import "time"

type OrderItem struct {
	Product Product `json:"Product"`
	OrderID string  `json:"OrderID"`
	Amount  uint64  `json:"Amount"`
	Price   float64 `json:"Price"`
}

type Order struct {
	OrderID    string      `json:"OrderID"`
	UserID     string      `json:"UserID"`
	OrderItems []OrderItem `json:"OrderItems"`
	Date       time.Time   `json:"Date"`
	Status     string      `json:"Status"`
}
