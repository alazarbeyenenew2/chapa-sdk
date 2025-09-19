package service

import (
	"context"

	"github.com/alazarbeyenenew2/chapasdk/pkg/constants"
	"github.com/alazarbeyenenew2/chapasdk/pkg/model"
	"go.uber.org/zap"
)

func (s *Chapa) Transfer(ctx context.Context, req model.TransferRequest) (model.TransferResponse, error) {
	var response model.TransferResponse
	if err := s.HttpClient.DoPost(constants.TRANSFER_URL, req, map[string]string{"Authorization": s.authorization}, &response); err != nil {
		//log error
		s.logger.Error(err.Error(), zap.Any("req", req))
		return model.TransferResponse{}, err
	}
	return response, nil
}
