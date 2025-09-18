package model

type SubaccountResp struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		SubaccountID string `json:"subaccounts[id]"`
	} `json:"data"`
}
