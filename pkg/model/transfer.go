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

// check transfer
type CheckTransferRequest struct {
	BatchID int `json:"batch_id"`
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}
type CheckTransferResponse struct {
	Message string              `json:"message"`
	Meta    CheckTransferMeta   `json:"meta"`
	Data    []CheckTransferData `json:"data"`
}

type CheckTransferMeta struct {
	CurrentPage  int           `json:"current_page"`
	FirstPageURL string        `json:"first_page_url"`
	LastPage     int           `json:"last_page"`
	LastPageURL  string        `json:"last_page_url"`
	NextPageURL  string        `json:"next_page_url"`
	Path         string        `json:"path"`
	PerPage      int           `json:"per_page"`
	PrevPageURL  string        `json:"prev_page_url"`
	To           int           `json:"to"`
	Total        int           `json:"total"`
	Error        []interface{} `json:"error"`
}

type CheckTransferData struct {
	AccountName    string  `json:"account_name"`
	AccountNumber  *string `json:"account_number"`
	Currency       string  `json:"currency"`
	Amount         float64 `json:"amount"`
	Charge         float64 `json:"charge"`
	TransferType   string  `json:"transfer_type"`
	ChapaReference string  `json:"chapa_reference"`
	BankCode       int     `json:"bank_code"`
	BankName       string  `json:"bank_name"`
	BankReference  string  `json:"bank_reference"`
	Status         string  `json:"status"`
	Reference      string  `json:"reference"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}
