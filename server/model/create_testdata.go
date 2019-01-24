package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"market-server/server/common/types"
)

func CreateTestData() {
	var category types.Category
	var shop types.Shop
	var item types.Item
	category.Name = "超市"
	AddCategory(&category, nil)
	for i := 0; i < 2; i++ {
		shop.Name = "家乐福"
		shop.Lat = 1
		shop.Lon = 2
		shop.Img = "https://www.koudai.com/image/jlf.png"
		shop.CategoryId = category.Id
		shop.Star = 200
		shop.SellerId = 2
		shop.DispatchMin = 0
		shop.DispatchPrice = 5
		AddShop(&shop, nil)
		JsonPrint(shop)
		var itemCategory types.ItemCategory
		itemCategory.Name = "itemCategory"
		itemCategory.ShopId = shop.Id
		AddItemCategory(&itemCategory, nil)
		JsonPrint(itemCategory)
		item.Name = "item测试"
		item.ShopId = shop.Id
		item.Price = 100.123456
		item.Img = "img"
		item.BarCode = "123456"
		item.ItemCategoryId = itemCategory.Id
		AddItem(&item, nil)
		JsonPrint(item)
	}
}
func CreateTestDataFromFile(path string) {
	buf, err := ioutil.ReadFile(path)
	var categoryList []types.Category
	fmt.Println(string(buf))
	err = json.Unmarshal(buf, &categoryList)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range categoryList {
		AddCategory(&c, nil)
		for _, shop := range c.ShopList {
			shop.CategoryId = c.Id
			AddShop(&shop, nil)
			for _, itemCategory := range shop.ItemCategoryList {
				itemCategory.ShopId = shop.Id
				AddItemCategory(&itemCategory, nil)
				for _, item := range itemCategory.ItemList {
					item.ShopId = shop.Id
					item.ItemCategoryId = itemCategory.Id
					AddItem(&item, nil)
				}
			}
		}
	}
}
