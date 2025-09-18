package service

import (
	"context"

	"github.com/alazarbeyenenew2/chapasdk/pkg/constants"
	"github.com/alazarbeyenenew2/chapasdk/pkg/model"
)

func (s *Chapa) GetListOfBanks(ctx context.Context) (model.BankLists, error) {
	var response model.BankLists

	if err := s.HttpClient.DoGet(
		constants.BANK_LIST_URL,
		map[string]string{"Authorization": s.authorization},
		&response,
	); err != nil {
		s.logger.Error(err.Error())
		return model.BankLists{}, err
	}

	return response, nil
}
