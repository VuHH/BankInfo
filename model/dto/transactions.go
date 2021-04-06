package dto

import "time"

type Transaction struct {
	IdTransaction   int       `json:"id_transaction"`
	IdAccount       int       `json:"id_account"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
	Bank            string    `json:"bank"`
}
