package main

import (
	"example/mymodule/requests"
	"example/mymodule/rmq"
	"fmt"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid"
	proto "github.com/nejkit/processing-proto/balances"
)

func main() {
	router := gin.Default()
	rmq.InitRabbit()
	router.POST("emmit-balance", func(ctx *gin.Context) {
		var emitBalance requests.EmitBalance
		if err := ctx.ShouldBindJSON(&emitBalance); err != nil {
			ctx.JSON(400, gin.H{"Wrong request ": err.Error()})
		}
		err := ctx.BindJSON(emitBalance)
		fmt.Println(emitBalance.Amount)
		id, err := gonanoid.ID(10)
		if err != nil {
			ctx.JSON(500, gin.H{"Internal": err.Error()})
		}
		emmitBalanceEvent := proto.EmmitBalanceRequest{
			Id:       id,
			Address:  emitBalance.Address,
			Amount:   int32(emitBalance.Amount),
			Currency: emitBalance.Currency,
		}

		rmq.PublishMessage(emmitBalanceEvent)

		ctx.JSON(200, gin.H{"message": "Emmit success"})

	})

	router.Run(":" + "8000")

}
