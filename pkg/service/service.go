package service

import (
	httpclient "github.com/alazarbeyenenew2/chapasdk/platform/httpClient"
	"github.com/alazarbeyenenew2/chapasdk/platform/logger"
)

type Chapa struct {
	PublicKey          string
	Secretkey          string
	Encryptionkey      string
	InitTransactionUrl string
	HttpClient         *httpclient.Client
	logger             *logger.ZapLogger
}

func NewClient(
	PublicKey string,
	Secretkey string,
	Encryptionkey string,
	InitTransactionUrl string,
	HttpClient *httpclient.Client,
	logger *logger.ZapLogger,

) *Chapa {
	// test response
	return &Chapa{
		PublicKey:          PublicKey,
		Secretkey:          Secretkey,
		Encryptionkey:      Encryptionkey,
		InitTransactionUrl: InitTransactionUrl,
		HttpClient:         HttpClient,
		logger:             logger,
	}

}
