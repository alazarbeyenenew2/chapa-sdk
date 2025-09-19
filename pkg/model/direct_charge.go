package model

import "github.com/shopspring/decimal"

type directChargeType string

const (
	TELEBIR      = "telebirr"
	MPESA        = "mpesa"
	CBE_BIRR     = "CBEBirr"
	COOPAY_EBIRR = "Coopay-Ebirr"
)

type DirectChargeRequest struct {
	Type      directChargeType `json:"type"`
	Amount    decimal.Decimal  `json:"amount"`
	Currency  string           `json:"currency"`
	TxRef     string           `json:"tx_ref"`
	Mobile    string           `json:"mobile"`
	Email     string           `json:"email"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
}

type DirectChargeResponse struct {
	Message string                   `json:"message"`
	Status  string                   `json:"status"`
	Data    DirectChargeResponseData `json:"data"`
}

type DirectChargeResponseData struct {
	AuthType  string                  `json:"auth_type"`
	RequestID string                  `json:"requestID"`
	Meta      DirecthargeResponseMeta `json:"meta"`
	Mode      string                  `json:"mode"`
}

type DirecthargeResponseMeta struct {
	Message       string `json:"message"`
	Status        string `json:"status"`
	Portal        string `json:"portal"`
	RefID         string `json:"ref_id"`
	PaymentStatus string `json:"payment_status"`
}
