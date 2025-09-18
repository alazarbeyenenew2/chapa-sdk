package chapasdk

import "github.com/alazarbeyenenew2/chapasdk/pkg/service"

func NewClient(userName, password string) *service.Chapa {
	return service.NewClient(userName, password)
}
