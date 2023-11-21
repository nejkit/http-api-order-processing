package api

import (
	"example/mymodule/requests"
	"example/mymodule/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger  *logrus.Logger
	balServ services.BalanceService
}

func NewHandler(logger *logrus.Logger, bs services.BalanceService) Handler {
	return Handler{
		logger:  logger,
		balServ: bs,
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
	h.balServ.EmmitBalance(ctx, &emBalRequest)
	ctx.Status(http.StatusOK)
}

func (h *Handler) WalletInfoHandle(ctx *gin.Context) {
	var walinfoRequest requests.GetBalance
	val := validator.New()
	err := ctx.BindJSON(walinfoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
	}
	err = val.Struct(walinfoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
	}
	response := h.balServ.WalletInfo(ctx, &walinfoRequest)
	ctx.JSON(http.StatusOK, response)
}
