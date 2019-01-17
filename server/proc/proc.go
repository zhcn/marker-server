package proc

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"market-server/server/common/types"
	"net/http"
	"strconv"
)

func makeTestResp() []types.Category {
	var categoryList []types.Category
	var marketList []types.Market
	for i := 0; i < 20; i++ {
		var market types.Market
		market.Name = "测试" + strconv.Itoa(i)
		marketList = append(marketList, market)
	}
	for i := 0; i < 10; i++ {
		var category types.Category
		category.Name = "测试" + strconv.Itoa(i)
		category.MarketList = marketList
		categoryList = append(categoryList, category)
	}
	fmt.Println("test")
	return categoryList
}

func GetMarkets(c *gin.Context) {
	//var categoryList []Category
	res := makeTestResp()
	data, _ := json.Marshal(res)
	c.String(http.StatusOK, string(data))
}
