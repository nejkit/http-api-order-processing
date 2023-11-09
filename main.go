package main

import (
	"context"
	"example/mymodule/requests"
	"example/mymodule/rmq"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid"
	proto "github.com/nejkit/processing-proto/balances"
)

func main() {
	router := gin.Default()
	rmq.InitRabbit()
	router.POST("/emmit-balance", func(ctx *gin.Context) {
		var emitBalance *requests.EmitBalance

		err := ctx.BindJSON(&emitBalance)

		id, _ := gonanoid.ID(10)
		if err != nil {
			ctx.JSON(500, gin.H{"Internal": err.Error()})
			return
		}
		emmitBalanceEvent := proto.EmmitBalanceRequest{
			Id:       id,
			Address:  emitBalance.Address,
			Amount:   int32(emitBalance.Amount),
			Currency: emitBalance.Currency,
		}

		rmq.PublishMessage(&emmitBalanceEvent)

		ctx.JSON(200, gin.H{"message": "Emmit success"})

	})

	router.POST("/wallet-info", func(ctx *gin.Context) {
		fmt.Println("Api Called")
		var request requests.GetBalance
		err := ctx.BindJSON(&request)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(request.Address)
		id, _ := gonanoid.ID(10)
		getWalletEvent := proto.GetWalletInfoRequest{
			Id:      id,
			Address: request.Address,
		}
		rmq.SendEventGetWalletInfo(&getWalletEvent)
		response := rmq.CatchResponseWalletInfo(id)
		fmt.Print(response.String())
	})

	srv := http.Server{
		Addr:    "localhost:8035",
		Handler: router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Print(err.Error())
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	select {
	case <-ctx.Done():
		fmt.Print("Timeout 5 seconds")
	}

}
