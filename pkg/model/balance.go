package model

import "github.com/shopspring/decimal"

type BalanceResponseData struct {
	Currency         string          `json:"currency"`
	AvailableBalance decimal.Decimal `json:"available_balance"`
	LedgerBalance    decimal.Decimal `json:"ledger_balance"`
}

type BalanceResponse struct {
	Message string                `json:"message"`
	Status  string                `json:"status"`
	Data    []BalanceResponseData `json:"data"`
}
