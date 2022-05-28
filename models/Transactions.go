package models

type Transaction struct {
	Id       int
	Currency string
	Price    float64
	Issuer    string
	Receiver string
	Date     string
}
