package models

import "time"

type OrderItem struct {
	Product Product `json:"Product"`
	OrderID uint64  `json:"OrderID"`
	Amount  uint64  `json:"Amount"`
	Price   float64 `json:"Price"`
}

type Order struct {
	OrderID    uint64      `json:"OrderID"`
	UserID     uint64      `json:"UserID"`
	OrderItems []OrderItem `json:"OrderItems"`
	Date       time.Time   `json:"Date"`
	Status     string      `json:"Status"`
}
