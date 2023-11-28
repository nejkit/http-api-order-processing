package requests

import "example/mymodule/external/orders"

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
	InitialPrice   float32          `json:"initialPrice"`
	InitialVolume  float32          `json:"initialVolume"`
	OrderType      orders.OrderType `json:"orderType"`
	ExchangeWallet string           `json:"exchangeWallet"`
}

type CreateOrderResponse struct {
	Id string `json:"orderId"`
}

type MatchingData struct {
	FillVolume float64           `json:"filledVolume"`
	FillPrice  float64           `json:"filledPrice"`
	Date       uint64            `json:"creationDate"`
	State      orders.MatchState `json:"dealState"`
}

type OrderInfo struct {
	Id string `json:"orderId"`

	CurrencyPair   string         `json:"currencyPair"`
	Direction      int            `json:"direction"`
	InitPrice      float64        `json:"initialPrice"`
	MatchInfo      []MatchingData `json:"dealInfos"`
	InitVolume     float64        `json:"initialVolume"`
	ExchangeWallet string         `json:"exchangeWallet"`

	CreationDate   uint64 `json:"creationDate"`
	UpdatedDate    uint64 `json:"updatedDate"`
	ExpirationDate uint64 `json:"expirationDate"`

	OrderState int `json:"orderState"`
	OrderType  int `json:"orderType"`
}
