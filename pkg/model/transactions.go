package model

import "time"

type ViewTransactionsRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type ViewTransactionsResponse struct {
	Status string               `json:"status"`
	Data   ViewTransactionsData `json:"data"`
}

type ViewTransactionsData struct {
	Transactions []Transaction `json:"transactions"`
	Pagination   Pagination    `json:"pagination"`
}
type Transaction struct {
	Status        string   `json:"status"`
	RefID         string   `json:"ref_id"`
	Type          string   `json:"type"`
	CreatedAt     string   `json:"created_at"`
	Currency      string   `json:"currency"`
	Amount        string   `json:"amount"`
	Charge        string   `json:"charge"`
	TransID       string   `json:"trans_id"`
	PaymentMethod string   `json:"payment_method"`
	Customer      Customer `json:"customer"`
}

type Customer struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mobile    string `json:"mobile"`
}

type Pagination struct {
	PerPage      int    `json:"per_page"`
	CurrentPage  int    `json:"current_page"`
	FirstPageURL string `json:"first_page_url"`
	NextPageURL  string `json:"next_page_url"`
	PrevPageURL  string `json:"prev_page_url"`
}

// transaction logs

type TransactionLogsResponse struct {
	Message string                `json:"message"`
	Status  string                `json:"status"`
	Data    []TransactionLogsData `json:"data"`
}

type TransactionLogsData struct {
	Item      int       `json:"item"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
