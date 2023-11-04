package requests

type EmitBalance struct {
	Address  string `json:"address"`
	Currency string `json:"currency"`
	Amount   int32 `json:"amount"`
}
