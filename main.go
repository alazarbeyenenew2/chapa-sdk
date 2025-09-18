package chapasdk

import "github.com/alazarbeyenenew2/chapa-sdk/pkg/service"

func NewClient(userName, password string) *service.Chapa {
	return service.NewClient(userName, password)
}
