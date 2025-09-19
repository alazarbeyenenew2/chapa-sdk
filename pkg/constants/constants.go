package constants

const (
	INIT_ACCEPT_PAYMENT_URL   = "https://api.chapa.co/v1/transaction/initialize"
	VERIFY_ACCEPT_PAYMENT_URL = "https://api.chapa.co/v1/transaction/verify/%s"
	BANK_LIST_URL             = "https://api.chapa.co/v1/banks"
	CREATE_SUB_ACCOUNT_URL    = "https://api.chapa.co/v1/subaccount"
	VIEW_TRANSACTIONS_URL     = "https://api.chapa.co/v1/transactions/?per_page=%d&page=%d"
	TRANSACTION_LOGS_URL      = "https://api.chapa.co/v1/transaction/events/%s"
	TRANSFER_URL              = "https://api.chapa.co/v1/transfers"
	VERIFY_TRANSFER_URL       = "https://api.chapa.co/v1/transfers/verify/%s"
	CHECK_BULK_TRANSFER_URL   = "https://api.chapa.co/v1/transfers?batch_id=%d&page=%d&per_page=%d"
	CHECK_TRANSFER_URL        = "https://api.chapa.co/v1/transfers?page=%d&per_page=%d"
	BALANCE_URL               = "https://api.chapa.co/v1/balances"
	SWAP_URL                  = "https://api.chapa.co/v1/swap"
)
