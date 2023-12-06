package util

import (
	"errors"
	"example/mymodule/external/orders"
	"example/mymodule/statics"
)

func GetOrdersError(err orders.OrdersErrorCodes) error {
	switch err {
	case orders.OrdersErrorCodes_ORDERS_ERROR_CODE_NOT_ENOUGH_BALANCE:
		return errors.New(statics.ErrorNotEnoughBalance)
	case orders.OrdersErrorCodes_ORDERS_ERROR_CODE_NOT_EXISTS_CURRENCY:
		return errors.New(statics.ErrorNotExistsBalance)
	default:
		return errors.New(statics.InternalError)
	}

}
