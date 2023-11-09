package api

import (
	"example/mymodule/requests"
	"example/mymodule/rmq"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func EmmitBalance(ctx *gin.Context, logger *logrus.Logger, channel *amqp091.Channel) {
	logger.Info("Received request: ", ctx.Request.Body)
	var emmitBalanceRequest *requests.EmitBalanceRequest
	err := ctx.ShouldBindJSON(&emmitBalanceRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request:": err.Error()})
		return
	}
	rmq.SendEmmitBalanceRequest(emmitBalanceRequest, channel, logger)

	ctx.Status(http.StatusOK)
}

func GetWalletInfo(ctx *gin.Context, logger *logrus.Logger, channel *amqp091.Channel, msgs <-chan amqp091.Delivery) {
	logger.Info("Received request: ", ctx.Request.Body)
	var request *requests.GetBalance
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request:": err.Error()})
	}
	id := rmq.SendGetWalletInfoRequest(request, channel, logger)
	var response balances.GetWalletInfoResponse
	for msg := range msgs {
		proto.Unmarshal(msg.Body, &response)
		if response.Id == id {
			msg.Ack(false)
			break
		} else {
			msg.Nack(false, true)
		}
	}
	var balanceInfos []requests.BalanceInfo
	for _, value := range response.GetWalletInfo().GetBalanceInfos() {
		balanceInfos = append(balanceInfos, requests.BalanceInfo{
			Currency:      value.GetCurrency(),
			ActualBalance: value.GetActualBalance(),
			FreezeBalance: value.GetFreezeBalance()})
	}

	ctx.JSON(http.StatusOK, &requests.GetBalanceResponse{
		Address:  response.GetWalletInfo().GetAddress(),
		Created:  response.GetWalletInfo().GetCreated(),
		Balances: balanceInfos,
	})

}
