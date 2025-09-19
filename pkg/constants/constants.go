package constants

const (
	INIT_ACCEPT_PAYMENT_URL   = "https://api.chapa.co/v1/transaction/initialize"
	VERIFY_ACCEPT_PAYMENT_URL = "https://api.chapa.co/v1/transaction/verify/%s"
	BANK_LIST_URL             = "https://api.chapa.co/v1/banks"
	CREATE_SUB_ACCOUNT_URL    = "https://api.chapa.co/v1/subaccount"
	VIEW_TRANSACTIONS_URL     = "https://api.chapa.co/v1/transactions/per_page=%d&page=%d"
	TRANSACTION_LOGS_URL      = "https://api.chapa.co/v1/transaction/events/%s"
	TRANSFER_URL              = "https://api.chapa.co/v1/transfers"
)
