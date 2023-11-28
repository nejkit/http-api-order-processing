package statics

const (
	BalanceExchangeName        = "e.balances.forward"
	OrdersExchangeName         = "e.orders.forward"
	RkEmmitBalanceRequest      = "r.balances.#.EmmitBalanceRequest.#"
	RkGetWalletInfoRequest     = "r.balances.#.GetWalletInfoRequest.#"
	RkCreateOrderRequest       = "r.request.#.CreateOrderRequest.#"
	QueueGetWalletInfoResponse = "q.balances.response.GetWalletInfoResponse"
	QueueCreateOrderResponse   = "q.orders.response.CreateOrderResponse"
	QueueGetOrderResponse      = "q.orders.response.GetOrderResponse"
	RkGetOrderRequest          = "r.request.#.GetOrderRequest.#"
	ConfigPath                 = "./app_config.yaml"
	ErrorNotEnoughBalance      = "NotEnoughBalance"
	ErrorNotExistsBalance      = "NotExistsBalance"
	InternalError              = "InternalError"
)
