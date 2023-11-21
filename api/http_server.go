package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ServerApi struct {
	logger  *logrus.Logger
	handler Handler
	host    string
}

func NewServer(logger *logrus.Logger, handler Handler, host string) ServerApi {
	return ServerApi{logger: logger, handler: handler, host: host}
}

func (s *ServerApi) StartServe(ctx context.Context) {
	router := gin.Default()

	router.POST("/emmit-balance", s.handler.EmmitBalanceHandle)

	router.GET("/wallet-info", s.handler.WalletInfoHandle)
	server := &http.Server{
		Addr:    s.host,
		Handler: router,
	}

	go server.ListenAndServe()

	for {
		select {
		case <-ctx.Done():
			server.Close()
			break
		}
	}
}
