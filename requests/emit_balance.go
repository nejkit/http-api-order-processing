package requests

type EmitBalanceRequest struct {
	Address  string  `json:"address"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type EmmitBalanceResponse struct {
	Address string `json:"address"`
}

type GetBalance struct {
	Address string `json:"address"`
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
