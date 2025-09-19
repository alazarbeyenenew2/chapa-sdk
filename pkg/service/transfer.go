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

func (s *Chapa) BulkTransfer(ctx context.Context, req model.BulkTransferRequest) (model.BulkTransferResponse, error) {
	var response model.BulkTransferResponse
	if err := s.HttpClient.DoPost(constants.TRANSFER_URL, req, map[string]string{"Authorization": s.authorization}, &response); err != nil {
		//log error
		s.logger.Error(err.Error(), zap.Any("req", req))
		return model.BulkTransferResponse{}, err
	}
	return response, nil
}

func (s *Chapa) CheckBulkTransferStatus(ctx context.Context, request model.CheckBulkTransferRequest) (model.CheckTransferResponse, error) {
	var response model.CheckTransferResponse
	if err := s.HttpClient.DoGet(
		fmt.Sprintf(constants.CHECK_BULK_TRANSFER_URL, request.BatchID, request.Page, request.PerPage),
		map[string]string{"Authorization": s.authorization},
		&response,
	); err != nil {
		s.logger.Error(err.Error(), zap.Any("request", request))
		return model.CheckTransferResponse{}, err
	}

	return response, nil
}

func (s *Chapa) CheckTransferStatus(ctx context.Context, request model.CheckTransferRequest) (model.CheckTransferResponse, error) {
	var response model.CheckTransferResponse
	if err := s.HttpClient.DoGet(
		fmt.Sprintf(constants.CHECK_TRANSFER_URL, request.Page, request.PerPage),
		map[string]string{"Authorization": s.authorization},
		&response,
	); err != nil {
		s.logger.Error(err.Error(), zap.Any("request", request))
		return model.CheckTransferResponse{}, err
	}

	return response, nil
}
