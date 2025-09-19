package model

import (
	"time"

	"github.com/shopspring/decimal"
)

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

// verify transfer

type VerifyTransferResponseData struct {
	AccountName         string    `json:"account_name"`
	AccountNumber       string    `json:"account_number"`
	Mobile              string    `json:"mobile"`
	Currency            string    `json:"currency"`
	Amount              int       `json:"amount"`
	Charge              int       `json:"charge"`
	Mode                string    `json:"mode"`
	TransferMethod      string    `json:"transfer_method"`
	Narration           string    `json:"narration"`
	ChapaTransferID     string    `json:"chapa_transfer_id"`
	BankCode            int       `json:"bank_code"`
	BankName            string    `json:"bank_name"`
	CrossPartyReference string    `json:"cross_party_reference"`
	IPAddress           string    `json:"ip_address"`
	Status              string    `json:"status"`
	TxRef               string    `json:"tx_ref"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
type VerifyTransferResponse struct {
	Message string                     `json:"message"`
	Status  string                     `json:"status"`
	Data    VerifyTransferResponseData `json:"data"`
}

// bulk transfer
type BulkTransferRequest struct {
	Title    string             `json:"title"`
	Currency string             `json:"currency"`
	BulkData []BulkTransferData `json:"bulk_data"`
}

type BulkTransferData struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	Reference     string `json:"reference"`
	BankCode      int    `json:"bank_code"`
}

type BulkTransferResponse struct {
	Message string                   `json:"message"`
	Status  string                   `json:"status"`
	Data    BulkTransferResponseData `json:"data"`
}

type BulkTransferResponseData struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
