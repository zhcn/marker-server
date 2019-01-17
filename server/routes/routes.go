package routes

import (
	"github.com/gin-gonic/gin"
	"market-server/server/proc"
)

var Router *gin.Engine

func Init() {
	Router = gin.New()
	Router.GET("/get_markets", proc.GetMarkets)
	//Router.Get("/get_market_items", proc.GetMarketItems)
}
