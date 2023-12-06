package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerApi struct {
	handler Handler
	host    string
}

func NewServer(handler Handler, host string) ServerApi {
	return ServerApi{handler: handler, host: host}
}

func (s *ServerApi) StartServe(ctx context.Context) {
	router := gin.Default()

	router.POST("/emmit-balance", func(ctx *gin.Context) {
		s.handler.EmmitBalanceHandle(ctx)
	})

	router.GET("/wallet-info", func(ctx *gin.Context) {
		s.handler.WalletInfoHandle(ctx)
	})

	router.POST("/order", func(ctx *gin.Context) {
		s.handler.CreateOrderHandle(ctx)
	})

	router.GET("/order/:orderid", func(ctx *gin.Context) {
		s.handler.GetOrderHangle(ctx)
	})

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
