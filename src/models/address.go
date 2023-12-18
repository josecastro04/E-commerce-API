package models

type Address struct {
	UserID      string `json:"UserID"`
	City        string `json:"City"`
	Country     string `json:"Country"`
	Address     string `json:"Address"`
	Postal_Code string `json:"Postal_Code"`
	State       string `json:"State"`
}
