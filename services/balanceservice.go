package services

import (
	"context"
	"example/mymodule/external/balances"
	"example/mymodule/requests"
	"example/mymodule/rmq"
	"example/mymodule/statics"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type BalanceService struct {
	emmitSender    rmq.AmqpSender
	walInfoSender  rmq.AmqpSender
	walletListener rmq.Listener[balances.GetWalletInfoResponse]
}

func NewBalanceService(walLis rmq.Listener[balances.GetWalletInfoResponse], wls rmq.AmqpSender, ebs rmq.AmqpSender) BalanceService {
	return BalanceService{walletListener: walLis, walInfoSender: wls, emmitSender: ebs}
}

func (s *BalanceService) EmmitBalance(ctx context.Context, request *requests.EmitBalanceRequest) {
	id := uuid.NewString()
	event := &balances.EmmitBalanceRequest{
		Id:       id,
		Address:  request.Address,
		Currency: request.Currency,
		Amount:   request.Amount,
	}
	s.emmitSender.SendMessage(ctx, event)
}

func (s *BalanceService) WalletInfo(ctx context.Context, request *requests.GetBalance) *requests.GetBalanceResponse {
	id := uuid.NewString()

	event := &balances.GetWalletInfoRequest{
		Id:      id,
		Address: request.Address,
	}
	logger.Infoln("Event: ", event.String())

	s.walInfoSender.SendMessage(ctx, event)
	response := s.walletListener.ConsumeById(ctx, id)
	logger.Infoln("Receive response: ", response.String())
	return statics.Map(response)
}
