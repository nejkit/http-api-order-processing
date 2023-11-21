package util

import (
	"example/mymodule/external/balances"

	"google.golang.org/protobuf/proto"
)

func GetParserForWalletInfoListener() func([]byte) (*balances.GetWalletInfoResponse, error) {
	return func(b []byte) (*balances.GetWalletInfoResponse, error) {
		var message *balances.GetWalletInfoResponse
		err := proto.Unmarshal(b, message)
		if err != nil {
			return nil, err
		}
		return message, nil
	}
}

func GetIdFromWalletInfoResponse() func(*balances.GetWalletInfoResponse) string {
	return func(message *balances.GetWalletInfoResponse) string {
		return message.GetId()
	}
}
