package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

func StartServer(logger *logrus.Logger, port string, channel *amqp091.Channel, msgs <-chan amqp091.Delivery) *http.Server {
	router := gin.Default()

	router.POST("/emmit-balance", func(ctx *gin.Context) {
		EmmitBalance(ctx, logger, channel)
	})

	router.GET("/wallet-info", func(ctx *gin.Context) {
		GetWalletInfo(ctx, logger, channel, msgs)
	})

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	server.ListenAndServe()

	return server

}
