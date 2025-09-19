package service

import (
	"context"
	"fmt"

	"github.com/alazarbeyenenew2/chapasdk/pkg/constants"
	"github.com/alazarbeyenenew2/chapasdk/pkg/model"
	"go.uber.org/zap"
)

func (s *Chapa) ViewTransactions(ctx context.Context, req model.ViewTransactionsRequest) (model.ViewTransactionsResponse, error) {
	var response model.ViewTransactionsResponse

	if req.Page < 1 {
		req.Page = 1
	}

	if req.PerPage < 1 {
		req.PerPage = 10
	}

	if err := s.HttpClient.DoPost(fmt.Sprintf(constants.VIEW_TRANSACTIONS_URL, req.PerPage, req.Page), req, map[string]string{"Authorization": s.authorization}, &response); err != nil {
		//log error
		s.logger.Error(err.Error(), zap.Any("req", req))
		return model.ViewTransactionsResponse{}, err
	}
	return response, nil
}

func (s *Chapa) GetTransactionLogs(ctx context.Context, PaymentRef string) (model.TransactionLogsResponse, error) {
	var response model.TransactionLogsResponse
	if err := s.HttpClient.DoGet(
		fmt.Sprintf(constants.TRANSACTION_LOGS_URL, PaymentRef),
		map[string]string{"Authorization": s.authorization},
		&response,
	); err != nil {
		s.logger.Error(err.Error(), zap.Any("payment ref", PaymentRef))
		return model.TransactionLogsResponse{}, err
	}

	return response, nil
}
