package proc

import (
	"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	"market-server/server/common/types"
	"market-server/server/model"
	"net/http"
	"strconv"
)

/*func makeTestResp() []types.Category {
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
}*/

func getItemCategoryList(shop *types.Shop, ctx *types.Context) ([]types.ItemCategory, error) {
	itemCategoryList, _ := model.GetItemCategory(shop, ctx)
	itemList, _ := model.GetAllItems(shop, ctx)
	for _, e := range itemList {
		for i, c := range itemCategoryList {
			if c.Id == e.ItemCategoryId {
				itemCategoryList[i].ItemList = append(itemCategoryList[i].ItemList, e)
			}
		}
	}
	return itemCategoryList, nil
}

func GetNearShops(pos *types.Position, ctx *types.Context) ([]types.Category, error) {
	//todo : add dist filter
	categoryList, _ := model.GetCategory(ctx)
	shopList, _ := model.GetAllShop(ctx)
	for _, e := range shopList {
		e.ItemCategoryList, _ = getItemCategoryList(&e, ctx)
		for j, c := range categoryList {
			if c.Id == e.CategoryId {
				categoryList[j].ShopList = append(categoryList[j].ShopList, e)
			}
		}
	}
	return categoryList, nil
}

func GetMarkets(c *gin.Context) {
	var pos types.Position
	pos.Lat, _ = strconv.ParseFloat(c.Query("lat"), 64)
	pos.Lon, _ = strconv.ParseFloat(c.Query("lon"), 64)
	res, _ := GetNearShops(&pos, nil)
	data, _ := json.Marshal(res)
	c.String(http.StatusOK, string(data))
}

/*func Order(c *gin.Context) {
	orderJson := c.PostForm("order")
	var order types.Order
	err := json.Unmarshal([]byte(orderJson), &order)
	if err != nil {
		c.String(http.StatusOK, "param error")
		return
	}
	AddOrder(&order, nil)
	//todo 发往抢单队列
}*/
