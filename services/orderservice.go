package services

import (
	"context"
	"errors"
	"example/mymodule/external/orders"
	"example/mymodule/requests"
	"example/mymodule/rmq"
	"example/mymodule/util"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type OrderService struct {
	createSender   rmq.AmqpSender
	getSender      rmq.AmqpSender
	createListener rmq.Listener[orders.CreateOrderResponse]
	getListener    rmq.Listener[orders.GetOrderResponse]
}

func NewOrderService(createListener rmq.Listener[orders.CreateOrderResponse], getListener rmq.Listener[orders.GetOrderResponse], cs rmq.AmqpSender, gs rmq.AmqpSender) OrderService {
	return OrderService{createListener: createListener, getListener: getListener, createSender: cs, getSender: gs}
}

func (s *OrderService) CreateOrder(ctx context.Context, request *requests.CreateOrderRequest) (*requests.CreateOrderResponse, error) {
	id := uuid.NewString()
	event := orders.CreateOrderRequest{
		Id:             id,
		CurrencyPair:   request.CurrencyPair,
		Direction:      request.Direction,
		InitPrice:      request.InitialPrice,
		InitVolume:     request.InitialVolume,
		OrderType:      request.OrderType,
		ExchangeWallet: request.ExchangeWallet,
	}

	logger.Infoln("Event: ", event.String())
	s.createSender.SendMessage(ctx, &event)
	response := s.createListener.ConsumeById(ctx, id)
	logger.Infoln("Receive response: ", response.String())
	if response.GetError() != nil {
		return nil, util.GetOrdersError(response.GetError().GetErorCode())
	}
	return &requests.CreateOrderResponse{Id: response.GetOrderId()}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, orderId string) (*requests.OrderInfo, error) {
	id := uuid.NewString()
	event := orders.GetOrderRequest{
		Id:      id,
		OrderId: orderId,
	}

	logger.Infoln("Event: ", event.String())
	s.getSender.SendMessage(ctx, &event)
	response := s.getListener.ConsumeById(ctx, id)
	return parseOrderInfo(response)
}

func parseOrderInfo(response *orders.GetOrderResponse) (*requests.OrderInfo, error) {
	if response.GetOrderData() == nil {
		return nil, errors.New("NotFound")
	}
	orderData := response.GetOrderData()

	orderInfo := &requests.OrderInfo{
		Id:             orderData.GetId(),
		CurrencyPair:   orderData.GetCurrencyPair(),
		Direction:      int(orderData.GetDirection()),
		InitPrice:      orderData.GetInitPrice(),
		FillPrice:      orderData.FillPrice,
		FillVolume:     orderData.FillVolume,
		MatchingDate:   orderData.Date.AsTime(),
		InitVolume:     orderData.GetInitVolume(),
		ExchangeWallet: orderData.GetExchangeWallet(),
		CreationDate:   orderData.GetCreationDate().AsTime(),
		UpdatedDate:    orderData.GetUpdatedDate().AsTime(),
		ExpirationDate: orderData.GetExpirationDate().AsTime(),
		OrderState:     int(orderData.GetOrderState()),
		OrderType:      int(orderData.GetOrderType()),
	}
	return orderInfo, nil
}
