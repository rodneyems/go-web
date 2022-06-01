package transactions

type transaction struct {
	Id       int     `json:"id"`
	Currency string  `json:"currency" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Issuer   string  `json:"issuer" binding:"required"`
	Receiver string  `json:"receiver" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}
