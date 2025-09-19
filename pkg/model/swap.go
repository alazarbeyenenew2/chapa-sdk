package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type SwapRequest struct {
	Amount       decimal.Decimal `json:"amount"`
	FromCurrency string          `json:"from"`
	ToCurrency   string          `json:"to"`
}

type SwapResponse struct {
	Message string           `json:"message"`
	Status  string           `json:"status"`
	Data    SwapResponseData `json:"data"`
}

type SwapResponseData struct {
	Status          string    `json:"status"`
	RefID           string    `json:"ref_id"`
	FromCurrency    string    `json:"from_currency"`
	ToCurrency      string    `json:"to_currency"`
	Amount          int       `json:"amount"`
	ExchangedAmount int       `json:"exchanged_amount"`
	Charge          int       `json:"charge"`
	Rate            int       `json:"rate"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
