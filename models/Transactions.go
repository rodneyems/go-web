package models

type Transaction struct {
	Id       int     `json:"id"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
	Issuer   string  `json:"issuer"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}
