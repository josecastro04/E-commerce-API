package models

import "time"

type ProductReviews struct {
	ReviewID  uint64    `json:"ReviewID"`
	ProductID uint64    `json:"ProductID"`
	UserID    uint64    `json:"UserID"`
	Review    string    `json:"Review"`
	Stars     uint64    `json:"Stars"`
	Date      time.Time `json:"Date"`
}

type Images struct {
	ImageID  uint64 `json:"ImageID"`
	Filename string `json:"Filename"`
	Path     string `json:"Path"`
}

type Product struct {
	ID          string    `json:"ID"`
	Name        string    `json:"Name"`
	Description string    `json:"Desc"`
	Price       float64   `json:"Price"`
	Stock       uint64    `json:"Stock"`
	Image       Images    `json:"Image"`
	AddedIn     time.Time `json:"AddedIn"`
}
