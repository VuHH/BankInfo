package dto

type RequestTransaction struct {
	AccountId       int     `json:"account_id" binding:"required,gte=1"`
	Amount          float64 `json:"amount" binding:"required,gte=1"`
	TransactionType string  `json:"transaction_type" binding:"required"`
}
