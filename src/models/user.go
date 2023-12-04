package models

import "time"

type User struct {
	ID        uint64    `json:"ID"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	Password  string    `json:"Password"`
	CreatedIn time.Time `json:"CreatedIn"`
}
