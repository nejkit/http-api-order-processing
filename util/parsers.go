package util

import (
	"example/mymodule/external/balances"
	"example/mymodule/external/orders"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func GetParserForWalletInfoListener() func([]byte) (*balances.GetWalletInfoResponse, error) {
	return func(b []byte) (*balances.GetWalletInfoResponse, error) {
		fmt.Println(b)
		var message balances.GetWalletInfoResponse
		err := proto.Unmarshal(b, &message)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		return &message, nil
	}
}

func GetParserForCreateOrderListener() func([]byte) (*orders.CreateOrderResponse, error) {
	return func(b []byte) (*orders.CreateOrderResponse, error) {
		var message orders.CreateOrderResponse
		err := proto.Unmarshal(b, &message)
		if err != nil {
			return nil, err
		}
		return &message, nil
	}
}

func GetParserForGetOrderListener() func([]byte) (*orders.GetOrderResponse, error) {
	return func(b []byte) (*orders.GetOrderResponse, error) {
		var message orders.GetOrderResponse
		err := proto.Unmarshal(b, &message)
		if err != nil {
			return nil, err
		}
		return &message, nil
	}
}

func GetIdFromWalletInfoResponse() func(*balances.GetWalletInfoResponse) string {
	return func(message *balances.GetWalletInfoResponse) string {
		return message.GetId()
	}
}

func GetIdFromCreateOrderResponse() func(*orders.CreateOrderResponse) string {
	return func(cor *orders.CreateOrderResponse) string {
		return cor.GetId()
	}
}

func GetIdFromGetOrderResponse() func(*orders.GetOrderResponse) string {
	return func(gor *orders.GetOrderResponse) string {
		return gor.GetId()
	}
}
