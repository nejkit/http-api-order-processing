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
	logger.Info("Received request: ", ctx.Request.Body)
	var emmitBalanceRequest *requests.EmitBalanceRequest
	err := ctx.ShouldBindJSON(&emmitBalanceRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request:": err.Error()})
		return
	}
	go rmq.SendEmmitBalanceRequest(emmitBalanceRequest, channel, logger)

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

	response := rmq.ConsumeWalletInfoResponse(id, msgs)

	httpResponse := statics.Map(response)

	ctx.JSON(http.StatusOK, httpResponse)

}
