package main

import (
	"context"
	"example/mymodule/api"
	"example/mymodule/external/balances"
	"example/mymodule/rmq"
	"example/mymodule/services"
	"example/mymodule/statics"
	"example/mymodule/util"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(statics.ConfigPath)
	viper.ReadInConfig()
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	amqpFactory := rmq.NewFactory(logger, viper.GetString("connection.rabbitmq"))
	walletInfoLis := rmq.NewListener[balances.GetWalletInfoResponse](amqpFactory, statics.QueueGetWalletInfoResponse, util.GetParserForWalletInfoListener(), util.GetIdFromWalletInfoResponse())
	senders := map[int]rmq.AmqpSender{
		statics.SendEmmitBalanceRequest: amqpFactory.NewSender(statics.BalanceExchangeName, statics.RkEmmitBalanceRequest),
		statics.SendWalletInfoRequest:   amqpFactory.NewSender(statics.BalanceExchangeName, statics.RkGetWalletInfoRequest),
	}

	balServ := services.NewBalanceService(logger, senders, walletInfoLis)
	handler := api.NewHandler(logger, balServ)
	server := api.NewServer(logger, handler, viper.GetString("http-server.port"))

	go server.StartServe(ctx)

	exit := make(chan os.Signal, 1)
	for {
		select {
		case <-exit:
			{
				cancel()

			}
		}
	}

}
