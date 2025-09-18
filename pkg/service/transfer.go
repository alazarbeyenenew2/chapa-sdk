package service

import (
	"context"

	"github.com/alazarbeyenenew2/chapasdk/pkg/constants"
	"github.com/alazarbeyenenew2/chapasdk/pkg/model"
	"go.uber.org/zap"
)

func (s *Chapa) InitTransaction(ctx context.Context, req model.InitTransactionReq) (model.InitTransactionResp, error) {
	var response model.InitTransactionResp
	if err := s.HttpClient.DoPost(constants.INIT_TRANSACTION_URL, req, map[string]string{"Authorization": s.authorization}, &response); err != nil {
		//log error
		s.logger.Error(err.Error(), zap.Any("req", req))
		return model.InitTransactionResp{}, err
	}
	return response, nil
}
