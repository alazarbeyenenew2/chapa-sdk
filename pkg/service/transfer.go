package service

import (
	"context"
	"fmt"

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

func (s *Chapa) VerifyTransfer(ctx context.Context, PaymentRef string) (model.VerifyTransferResponse, error) {
	var response model.VerifyTransferResponse
	if err := s.HttpClient.DoGet(
		fmt.Sprintf(constants.VERIFY_TRANSFER_URL, PaymentRef),
		map[string]string{"Authorization": s.authorization},
		&response,
	); err != nil {
		s.logger.Error(err.Error(), zap.Any("payment ref", PaymentRef))
		return model.VerifyTransferResponse{}, err
	}

	return response, nil
}
