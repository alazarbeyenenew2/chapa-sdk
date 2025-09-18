package service

import (
	"context"
	"fmt"

	"github.com/alazarbeyenenew2/chapasdk/pkg/constants"
	"github.com/alazarbeyenenew2/chapasdk/pkg/model"
	"go.uber.org/zap"
)

func (s *Chapa) InitAcceptPayment(ctx context.Context, req model.InitPaymentReq) (model.InitPaymentResp, error) {
	var response model.InitPaymentResp
	if err := s.HttpClient.DoPost(constants.INIT_ACCEPT_PAYMENT_URL, req, map[string]string{"Authorization": s.authorization}, &response); err != nil {
		//log error
		s.logger.Error(err.Error(), zap.Any("req", req))
		return model.InitPaymentResp{}, err
	}
	return response, nil
}

func (s *Chapa) VerifyAcceptPayment(cxt context.Context, PaymentRef string) (model.VerifyAcceptPaymentResp, error) {

	var response model.VerifyAcceptPaymentResp
	if err := s.HttpClient.DoGet(
		fmt.Sprintf(constants.VERIFY_ACCEPT_PAYMENT_URL, PaymentRef),
		map[string]string{"Authorization": s.authorization},
		&response,
	); err != nil {
		s.logger.Error(err.Error(), zap.Any("payment ref", PaymentRef))
		return model.VerifyAcceptPaymentResp{}, err
	}

	return response, nil
}

func (s *Chapa) CreateSubAccount(ctx context.Context, req model.SubAccount) (model.SubaccountResp, error) {
	var response model.SubaccountResp

	if err := s.HttpClient.DoPost(constants.CREATE_SUB_ACCOUNT_URL, req, map[string]string{"Authorization": s.authorization}, &response); err != nil {
		//log error
		s.logger.Error(err.Error(), zap.Any("req", req))
		return model.SubaccountResp{}, err
	}
	return response, nil
}
