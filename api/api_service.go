package api

import (
	"example/mymodule/requests"
	"example/mymodule/rmq"
	"example/mymodule/statics"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

func EmmitBalance(ctx *gin.Context, logger *logrus.Logger, channel *amqp091.Channel) {
	var emmitBalanceRequest requests.EmitBalanceRequest
	err := ctx.BindJSON(&emmitBalanceRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request:": err.Error()})
		return
	}
	logger.Info("Received request: ", emmitBalanceRequest.Address, emmitBalanceRequest.Amount, emmitBalanceRequest.Currency)

	rmq.SendEmmitBalanceRequest(emmitBalanceRequest, channel, logger)

	ctx.Status(http.StatusOK)
}

func GetWalletInfo(ctx *gin.Context, logger *logrus.Logger, channel *amqp091.Channel, msgs <-chan amqp091.Delivery) {
	var request requests.GetBalance
	err := ctx.BindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request:": err.Error()})
	}
	logger.Info("Received address: ", request.Address)

	id := rmq.SendGetWalletInfoRequest(request, channel, logger)

	response := rmq.ConsumeWalletInfoResponse(id, msgs, logger)

	httpResponse := statics.Map(response)

	ctx.JSON(http.StatusOK, httpResponse)

}
