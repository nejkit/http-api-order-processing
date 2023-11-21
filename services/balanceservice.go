package services

import (
	"context"
	"example/mymodule/external/balances"
	"example/mymodule/requests"
	"example/mymodule/rmq"
	"example/mymodule/statics"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type BalanceService struct {
	logger         *logrus.Logger
	senders        map[int]rmq.AmqpSender
	walletListener rmq.Listener[balances.GetWalletInfoResponse]
}

func NewBalanceService(logger *logrus.Logger, senders map[int]rmq.AmqpSender, walLis rmq.Listener[balances.GetWalletInfoResponse]) BalanceService {
	return BalanceService{logger: logger, senders: senders, walletListener: walLis}
}

func (s *BalanceService) EmmitBalance(ctx context.Context, request *requests.EmitBalanceRequest) {
	id := uuid.NewString()
	event := &balances.EmmitBalanceRequest{
		Id:       id,
		Address:  request.Address,
		Currency: request.Currency,
		Amount:   request.Amount,
	}

	sender, ok := s.senders[statics.SendEmmitBalanceRequest]
	if ok {
		sender.SendMessage(ctx, event)
	}
}

func (s *BalanceService) WalletInfo(ctx context.Context, request *requests.GetBalance) *requests.GetBalanceResponse {
	id := uuid.NewString()

	event := &balances.GetWalletInfoRequest{
		Id:      id,
		Address: request.Address,
	}

	sender, ok := s.senders[statics.SendWalletInfoRequest]
	if !ok {
		s.logger.Errorln("Fail get sender! ")
		return nil
	}

	sender.SendMessage(ctx, event)
	response := s.walletListener.ConsumeById(ctx, id)

	return statics.Map(response)
}
