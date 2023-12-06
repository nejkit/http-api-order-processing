package main

import (
	"context"
	"example/mymodule/api"
	"example/mymodule/external/balances"
	"example/mymodule/external/orders"
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

	logrus.SetLevel(logrus.InfoLevel)

	amqpFactory := rmq.NewFactory(viper.GetString("connection.rabbitmq"))
	walletInfoLis := rmq.NewListener[balances.GetWalletInfoResponse](amqpFactory, statics.QueueGetWalletInfoResponse, util.GetParserForWalletInfoListener(), util.GetIdFromWalletInfoResponse())
	createOrderLis := rmq.NewListener[orders.CreateOrderResponse](amqpFactory, statics.QueueCreateOrderResponse, util.GetParserForCreateOrderListener(), util.GetIdFromCreateOrderResponse())
	getOrderLis := rmq.NewListener[orders.GetOrderResponse](amqpFactory, statics.QueueGetOrderResponse, util.GetParserForGetOrderListener(), util.GetIdFromGetOrderResponse())
	emmitSender := amqpFactory.NewSender(statics.BalanceExchangeName, statics.RkEmmitBalanceRequest)
	getWalletSender := amqpFactory.NewSender(statics.BalanceExchangeName, statics.RkGetWalletInfoRequest)
	createOrderSender := amqpFactory.NewSender(statics.OrdersExchangeName, statics.RkCreateOrderRequest)
	getOrderSender := amqpFactory.NewSender(statics.OrdersExchangeName, statics.RkGetOrderRequest)
	oServ := services.NewOrderService(createOrderLis, getOrderLis, createOrderSender, getOrderSender)
	balServ := services.NewBalanceService(walletInfoLis, getWalletSender, emmitSender)
	handler := api.NewHandler(balServ, oServ)
	server := api.NewServer(handler, viper.GetString("http-server.port"))

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
