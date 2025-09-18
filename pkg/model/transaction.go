package model

import (
	"github.com/shopspring/decimal"
)

type InitTransactionReq struct {
	Amount                   decimal.Decimal `json:"amount"`
	Currency                 string          `json:"currency"`
	Email                    string          `json:"email"`
	FirstName                string          `json:"first_name"`
	LastName                 string          `json:"last_name"`
	PhoneNumber              string          `json:"phone_number"`
	TxRef                    string          `json:"tx_ref"`
	CallbackURL              string          `json:"callback_url"`
	ReturnURL                string          `json:"return_url"`
	CustomizationTitle       string          `json:"customization[title]"`
	CustomizationDescription string          `json:"customization[description]"`
	Meta                     bool            `json:"meta[hide_receipt]"`
}

type InitTransactionRespData struct {
	CheckoutURL string `json:"checkout_url"`
}

type InitTransactionResp struct {
	Message string                  `json:"message"`
	Status  string                  `json:"status"`
	Data    InitTransactionRespData `json:"data"`
}
