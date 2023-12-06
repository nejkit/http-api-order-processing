package requests

import (
	"example/mymodule/external/orders"
	"time"
)

type EmitBalanceRequest struct {
	Address  string  `json:"address" validate:"required,min=10,max=20"`
	Currency string  `json:"currency" validate:"required,len=3"`
	Amount   float64 `json:"amount" validate:"required,gt=0.00"`
}

type EmmitBalanceResponse struct {
	Address string `json:"address"`
}

type GetBalance struct {
	Address string `json:"address" validate:"required"`
}

type BalanceInfo struct {
	Currency      string  `json:"currency"`
	ActualBalance float64 `json:"actualBalance"`
	FreezeBalance float64 `json:"freezeBalance"`
}

type GetBalanceResponse struct {
	Address  string        `json:"address"`
	Created  uint64        `json:"created"`
	Balances []BalanceInfo `json:"balanceInfos"`
}

type CreateOrderRequest struct {
	CurrencyPair   string           `json:"currencyPair"`
	Direction      orders.Direction `json:"direction"`
	InitialPrice   float64          `json:"initialPrice"`
	InitialVolume  float64          `json:"initialVolume"`
	OrderType      orders.OrderType `json:"orderType"`
	ExchangeWallet string           `json:"exchangeWallet"`
}

type CreateOrderResponse struct {
	Id string `json:"orderId"`
}

type OrderInfo struct {
	Id string `json:"orderId"`

	CurrencyPair   string    `json:"currencyPair"`
	Direction      int       `json:"direction"`
	InitPrice      float64   `json:"initialPrice"`
	FillPrice      float64   `json:"fillPrice"`
	FillVolume     float64   `json:"fillVolume"`
	MatchingDate   time.Time `json:"matchingDate"`
	InitVolume     float64   `json:"initialVolume"`
	ExchangeWallet string    `json:"exchangeWallet"`

	CreationDate   time.Time `json:"creationDate"`
	UpdatedDate    time.Time `json:"updatedDate"`
	ExpirationDate time.Time `json:"expirationDate"`

	OrderState int `json:"orderState"`
	OrderType  int `json:"orderType"`
}
