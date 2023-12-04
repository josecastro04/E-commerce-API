package models

type Password struct {
	Current string `json:"Current"`
	New     string `json:"New"`
}
