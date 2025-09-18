package model

import (
	"github.com/shopspring/decimal"
)

type InitPaymentReq struct {
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

type InitPaymentRespData struct {
	CheckoutURL string `json:"checkout_url"`
}

type InitPaymentResp struct {
	Message string              `json:"message"`
	Status  string              `json:"status"`
	Data    InitPaymentRespData `json:"data"`
}

type VerifyAcceptPaymentRespData struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Email         string  `json:"email"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	Charge        float64 `json:"charge"`
	Mode          string  `json:"mode"`
	Method        string  `json:"method"`
	Type          string  `json:"type"`
	Status        string  `json:"status"`
	Reference     string  `json:"reference"`
	TxRef         string  `json:"tx_ref"`
	Customization struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Logo        *string `json:"logo"`
	} `json:"customization"`
	Meta      *interface{} `json:"meta"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
}

type VerifyAcceptPaymentResp struct {
	Message string                      `json:"message"`
	Status  string                      `json:"status"`
	Data    VerifyAcceptPaymentRespData `json:"data"`
}

type SubAccount struct {
	BusinessName  string  `json:"business_name"`
	AccountName   string  `json:"account_name"`
	BankCode      int     `json:"bank_code"`
	AccountNumber string  `json:"account_number"`
	SplitValue    float64 `json:"split_value"`
	SplitType     string  `json:"split_type"`
}

type SplitPaymentReq struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PhoneNumber   string `json:"phone_number"`
	TxRef         string `json:"tx_ref"`
	CallbackURL   string `json:"callback_url"`
	ReturnURL     string `json:"return_url"`
	Customization struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"customization"`
	Subaccounts struct {
		ID string `json:"id"`
	} `json:"subaccounts"`
}
