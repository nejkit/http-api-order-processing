package requests

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
