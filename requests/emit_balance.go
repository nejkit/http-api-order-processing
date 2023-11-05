package requests

type EmitBalance struct {
	Address  string `json:"address"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}
