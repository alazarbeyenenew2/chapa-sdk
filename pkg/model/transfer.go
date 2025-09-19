package model

import "github.com/shopspring/decimal"

type TransferRequest struct {
	AccountName   string          `json:"account_name"`
	AccountNumber string          `json:"account_number"`
	Amount        decimal.Decimal `json:"amount"`
	Currency      string          `json:"currency"`
	Reference     string          `json:"reference"`
	BankCode      int             `json:"bank_code"`
}

type TransferResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    string `json:"data"`
}
