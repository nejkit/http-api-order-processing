package api

import (
	"context"
	"example/mymodule/rmq"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

func StartServer(logger *logrus.Logger, port string, connection *amqp091.Connection, ctx context.Context) *http.Server {
	router := gin.Default()
	emmitBalanceReqChannel := rmq.GetChannel(connection, logger)
	walletInfoReqChannel := rmq.GetChannel(connection, logger)
	walletInfoRespChannel := rmq.GetChannel(connection, logger)
	router.POST("/emmit-balance", func(ctx *gin.Context) {
		EmmitBalance(ctx, logger, emmitBalanceReqChannel)
	})

	listenerWalletInfoResponse := rmq.CreateLisWalletInfoResponce(walletInfoRespChannel, logger)

	router.GET("/wallet-info", func(ctx *gin.Context) {
		GetWalletInfo(ctx, logger, walletInfoReqChannel, listenerWalletInfoResponse)
	})

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go server.ListenAndServe()

	for {
		select {
		case <-ctx.Done():
			server.Close()
			emmitBalanceReqChannel.Close()
			walletInfoReqChannel.Close()
			walletInfoRespChannel.Close()
			break
		}
	}

}
