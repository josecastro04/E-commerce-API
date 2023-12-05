package models

import "time"

type Product struct {
	ID          uint64    `json:"ID"`
	Name        string    `json:"Name"`
	Description string    `json:"Desc"`
	Price       float64   `json:"Price"`
	Stock       uint64    `json:"Stock"`
	AddedIn     time.Time `json:"AddedIn"`
}
