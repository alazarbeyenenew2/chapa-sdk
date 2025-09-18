package service

import (
	"context"

	"github.com/alazarbeyenenew2/chapasdk/pkg/constants"
	"github.com/alazarbeyenenew2/chapasdk/pkg/model"
	"go.uber.org/zap"
)

func (s *Chapa) CreateSubAccount(ctx context.Context, req model.SubAccount) (model.SubaccountResp, error) {
	var response model.SubaccountResp

	if err := s.HttpClient.DoPost(constants.CREATE_SUB_ACCOUNT_URL, req, map[string]string{"Authorization": s.authorization}, &response); err != nil {
		//log error
		s.logger.Error(err.Error(), zap.Any("req", req))
		return model.SubaccountResp{}, err
	}
	return response, nil
}
