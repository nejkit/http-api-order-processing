package main

import (
	"context"
	"example/mymodule/api"
	"example/mymodule/rmq"
	"example/mymodule/statics"
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
	conRmq, err := rmq.GetConnection(viper.GetString("connection.rabbitmq"), logger)
	defer conRmq.Close()
	if err != nil {
		panic(err.Error())
	}
	go api.StartServer(logger, viper.GetString("http-server.port"), conRmq, ctx)
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
