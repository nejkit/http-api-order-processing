package main

import (
	"example/mymodule/api"
	"example/mymodule/rmq"

	"github.com/sirupsen/logrus"
)

func main() {
	var logger *logrus.Logger
	logger.SetLevel(logrus.InfoLevel)
	channelRmq := rmq.GetChannel("", logger)
	listenerWalletInfo := rmq.CreateLisWalletInfoResponce(channelRmq, logger)
	api.StartServer(logger, "", channelRmq, listenerWalletInfo)

}
