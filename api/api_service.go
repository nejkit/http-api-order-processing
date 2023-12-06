package api

import (
	"example/mymodule/requests"
	"example/mymodule/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	balServ services.BalanceService
	orServ  services.OrderService
}

func NewHandler(bs services.BalanceService, os services.OrderService) Handler {
	return Handler{
		balServ: bs,
		orServ:  os,
	}
}

func (h *Handler) EmmitBalanceHandle(ctx *gin.Context) {
	var emBalRequest requests.EmitBalanceRequest
	val := validator.New()
	err := ctx.BindJSON(&emBalRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
		return
	}
	err = val.Struct(emBalRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
		return
	}
	h.balServ.EmmitBalance(ctx, &emBalRequest)
	ctx.Status(http.StatusOK)
}

func (h *Handler) WalletInfoHandle(ctx *gin.Context) {
	var walinfoRequest requests.GetBalance
	val := validator.New()
	err := ctx.BindJSON(&walinfoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
		return
	}
	err = val.Struct(walinfoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
		return
	}
	response := h.balServ.WalletInfo(ctx, &walinfoRequest)
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) CreateOrderHandle(ctx *gin.Context) {
	var createOrderRequest requests.CreateOrderRequest
	val := validator.New()
	err := ctx.BindJSON(&createOrderRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
		return
	}
	err = val.Struct(createOrderRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
		return
	}
	response, err := h.orServ.CreateOrder(ctx, &createOrderRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Wrong request ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetOrderHangle(ctx *gin.Context) {
	orderid := ctx.Param("orderid")
	if orderid == "" {
		ctx.JSON(http.StatusBadRequest, "OrderNotFound")
		return
	}
	response, err := h.orServ.GetOrder(ctx, orderid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "OrderNotFound")
		return
	}
	ctx.JSON(http.StatusOK, response)
}
