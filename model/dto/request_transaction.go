package dto

type RequestTransaction struct {
	AccountId       int     `json:"account_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
	TransactionType string  `json:"transaction_type" binding:"required"`
}
