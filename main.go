package main

import (
	"example/mymodule/api"
	"example/mymodule/rmq"
	"example/mymodule/statics"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(statics.ConfigPath)
	viper.ReadInConfig()

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	channelRmq := rmq.GetChannel(viper.GetString("connection.rabbitmq"), logger)
	listenerWalletInfo := rmq.CreateLisWalletInfoResponce(channelRmq, logger)
	api.StartServer(logger, viper.GetString("http-server.port"), channelRmq, listenerWalletInfo)

}
