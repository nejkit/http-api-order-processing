package statics

import (
	"example/mymodule/requests"

	"github.com/nejkit/processing-proto/balances"
)

func Map(protoResponse *balances.GetWalletInfoResponse) *requests.GetBalanceResponse {
	var response requests.GetBalanceResponse
	data := *protoResponse.GetWalletInfo()
	response.Address = data.GetAddress()
	response.Created = data.GetCreated()
	for _, balanceInfo := range data.BalanceInfos {
		response.Balances = append(response.Balances, requests.BalanceInfo{
			Currency:      balanceInfo.GetCurrency(),
			ActualBalance: balanceInfo.GetActualBalance(),
			FreezeBalance: balanceInfo.GetFreezeBalance()})
	}

	return &response
}
