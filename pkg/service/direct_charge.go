package service

import (
	"context"
	"fmt"

	"github.com/alazarbeyenenew2/chapasdk/pkg/constants"
	"github.com/alazarbeyenenew2/chapasdk/pkg/model"
	"go.uber.org/zap"
)

func (s *Chapa) InitiateDirectPayments(ctx context.Context, req model.DirectChargeRequest) (model.DirectChargeResponse, error) {

	if !constants.DIRECT_CHARGE_TYPES[string(req.Type)] {
		err := fmt.Errorf("invalid payment type")
		s.logger.Error(err.Error(), zap.Any("req", req))
		return model.DirectChargeResponse{}, err
	}

	var response model.DirectChargeResponse
	if err := s.HttpClient.DoPost(fmt.Sprintf(constants.DIRECT_CHARGE_URL, req.Type), req, map[string]string{"Authorization": s.authorization}, &response); err != nil {
		//log error
		s.logger.Error(err.Error(), zap.Any("req", req))
		return model.DirectChargeResponse{}, err
	}
	return response, nil
}
