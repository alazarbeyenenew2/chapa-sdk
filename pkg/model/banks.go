package model

import "time"

type Bank struct {
	ID            int       `json:"id"`
	Slug          string    `json:"slug"`
	Swift         string    `json:"swift"`
	Name          string    `json:"name"`
	AcctLength    int       `json:"acct_length"`
	CountryID     int       `json:"country_id"`
	IsMobileMoney *bool     `json:"is_mobilemoney"`
	IsActive      int       `json:"is_active"`
	IsRtgs        int       `json:"is_rtgs"`
	Active        int       `json:"active"`
	Is24Hrs       *bool     `json:"is_24hrs"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Currency      string    `json:"currency"`
}

type BankLists struct {
	Message string `json:"message"`
	Data    []Bank `json:"data"`
	Status  string `json:"status"`
}
