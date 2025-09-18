package chapasdk

import (
	"log"

	"github.com/alazarbeyenenew2/chapasdk/pkg/service"
	httpclient "github.com/alazarbeyenenew2/chapasdk/platform/httpClient"
	"github.com/alazarbeyenenew2/chapasdk/platform/logger"
	"github.com/spf13/viper"
)

func NewClient(
	PublicKey string,
	Secretkey string,
	Encryptionkey string,
	IsProduction bool,
) *service.Chapa {
	httpClient := httpclient.NewClient()
	logger, err := logger.NewLogger(IsProduction)
	if err != nil {
		log.Println("unable to initialize logger")
		log.Fatal(1)
	}

	return service.NewClient(
		PublicKey, Secretkey, Encryptionkey, viper.GetString("chapa.initTransactionURL"), httpClient, logger,
	)
}
