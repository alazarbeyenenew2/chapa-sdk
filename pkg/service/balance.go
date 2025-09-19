package service

import (
	"context"

	"github.com/alazarbeyenenew2/chapasdk/pkg/constants"
	"github.com/alazarbeyenenew2/chapasdk/pkg/model"
)

func (s *Chapa) GetBalance(ctx context.Context) (model.BalanceResponse, error) {
	var response model.BalanceResponse
	if err := s.HttpClient.DoGet(
		constants.BALANCE_URL,
		map[string]string{"Authorization": s.authorization},
		&response,
	); err != nil {
		s.logger.Error(err.Error())
		return model.BalanceResponse{}, err
	}

	return response, nil
}
