package api

import (
	"example/mymodule/external/balances"
	"example/mymodule/external/orders"
	"example/mymodule/requests"
	"example/mymodule/rmq"
	"example/mymodule/statics"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger              *logrus.Logger
	emmitSender         rmq.AmqpSender
	walletInfoSender    rmq.AmqpSender
	createOrderSender   rmq.AmqpSender
	walletInfoListener  rmq.Listener[balances.GetWalletInfoResponse]
	createOrderListener rmq.Listener[orders.CreateOrderResponse]
}

func NewHandler(logger *logrus.Logger, emSend rmq.AmqpSender, walInfSend rmq.AmqpSender, crOrderSend rmq.AmqpSender, walletLis rmq.Listener[balances.GetWalletInfoResponse], crOrderLis rmq.Listener[orders.CreateOrderResponse]) Handler {
	return Handler{
		logger:              logger,
		emmitSender:         emSend,
		walletInfoSender:    walInfSend,
		createOrderSender:   crOrderSend,
		walletInfoListener:  walletLis,
		createOrderListener: crOrderLis,
	}
}

func (h *Handler) EmmitBalanceHandle(ctx *gin.Context) {
	var emBalRequest requests.EmitBalanceRequest
	val := validator.New()
	err := ctx.BindJSON(emBalRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
	}
	err = val.Struct(emBalRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
	}
	//add service mb
	ctx.Status(http.StatusOK)
}

func EmmitBalance(ctx *gin.Context, logger *logrus.Logger, channel *amqp091.Channel) {
	var emmitBalanceRequest requests.EmitBalanceRequest
	validate := validator.New()
	err := ctx.BindJSON(&emmitBalanceRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request:": err.Error()})
		return
	}

	err = validate.Struct(emmitBalanceRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message: ": err.Error()})
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
	validate := validator.New()

	err = validate.Struct(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message: ": err.Error()})
		return
	}

	logger.Info("Received address: ", request.Address)

	id := rmq.SendGetWalletInfoRequest(request, channel, logger)

	response := rmq.ConsumeWalletInfoResponse(id, msgs, logger)

	httpResponse := statics.Map(response)

	ctx.JSON(http.StatusOK, httpResponse)

}
