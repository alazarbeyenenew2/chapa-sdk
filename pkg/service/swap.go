package service

import (
	"context"

	"github.com/alazarbeyenenew2/chapasdk/pkg/constants"
	"github.com/alazarbeyenenew2/chapasdk/pkg/model"
	"go.uber.org/zap"
)

func (s *Chapa) SwapCurrency(cxt context.Context, request model.SwapRequest) (model.SwapResponse, error) {

	var response model.SwapResponse
	if err := s.HttpClient.DoGet(
		constants.SWAP_URL,
		map[string]string{"Authorization": s.authorization},
		&response,
	); err != nil {
		s.logger.Error(err.Error(), zap.Any("request", request))
		return model.SwapResponse{}, err
	}

	return response, nil
}
