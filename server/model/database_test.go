package model

import (
	"fmt"
	"market-server/server/common/types"
	"testing"
	"time"
)

var ctx types.Context

func TestAddCategory(t *testing.T) {
	var category types.Category
	category.Name = "测试golang"
	err := AddCategory(&category, &ctx)
	if err != nil {
		t.Errorf("AddCategory error")
	}
}

func TestGetCategory(t *testing.T) {
	_, err := GetCategory(&ctx)
	if err != nil {
		t.Errorf("GetCategory error")
	}
}

func TestUpdateCategory(t *testing.T) {
	var category types.Category
	category.Id = 6
	category.Name = "测试Update"
	err := UpdateCategory(&category, &ctx)
	if err != nil {
		t.Errorf("UpdateCategory error")
	}
}

func TestDelCategory(t *testing.T) {
	var category types.Category
	categoryList, _ := GetCategory(&ctx)
	category.Id = categoryList[0].Id
	jsonPrint(categoryList)
	err := DelCategory(&category, &ctx)
	if err != nil {
		t.Errorf("DelCategory error")
	}
}

func TestAddShop(t *testing.T) {
	var shop types.Shop
	shop.Name = "shop测试"
	shop.Lat = 1
	shop.Lon = 2
	shop.Img = "img"
	shop.CategoryId = 1
	shop.Star = 100
	shop.SellerId = 1
	shop.DispatchMin = 0
	shop.DispatchPrice = 5
	AddShop(&shop, &ctx)
}

func TestGetAllShop(t *testing.T) {
	GetAllShop(&ctx)
}

func TestUpdateShop(t *testing.T) {
	var shop types.Shop
	shop.Name = "shop测试Update"
	shop.Lat = 1
	shop.Lon = 2
	shop.Img = "img"
	shop.CategoryId = 1
	shop.Star = 200
	shop.SellerId = 2
	shop.DispatchMin = 1
	shop.DispatchPrice = 10
	AddShop(&shop, &ctx)
	shop.Lat = 100
	jsonPrint(shop)
	UpdateShop(&shop, &ctx)
	jsonPrint(shop)
}
func TestDelShop(t *testing.T) {
	var shop types.Shop
	shop.Name = "shop测试Del"
	shop.Lat = 1
	shop.Lon = 2
	shop.Img = "img"
	shop.CategoryId = 1
	shop.Star = 200
	shop.SellerId = 2
	AddShop(&shop, &ctx)
	DelShop(&shop, &ctx)
}
func TestAddItem(t *testing.T) {
	var item types.Item
	item.Name = "item测试add"
	item.ShopId = 1
	item.Price = 100.123456
	item.Img = "img"
	AddItem(&item, &ctx)
}
func TestUpdateItem(t *testing.T) {
	var item types.Item
	item.Name = "item测试Update"
	item.ShopId = 2
	item.Price = 100.123456
	item.Img = "img"
	AddItem(&item, &ctx)
	item.Img = "img-update"
	UpdateItem(&item, &ctx)

}
func TestDelItem(t *testing.T) {
	var item types.Item
	item.Name = "item测试Del"
	item.ShopId = 3
	item.Price = 100.123456
	item.Img = "img"
	AddItem(&item, &ctx)
	DelItem(&item, &ctx)
}
func TestGetAllItems(t *testing.T) {
	var shop types.Shop
	shop.Id = 1
	res, _ := GetAllItems(&shop, &ctx)
	jsonPrint(res)
}
func TestAddSeller(t *testing.T) {
	var seller types.Seller
	seller.Name = "seller测试Add"
	seller.Password = "123456"
	seller.AuthImg = "authimg"
	AddSeller(&seller, &ctx)
}
func TestUpdateSeller(t *testing.T) {
	var seller types.Seller
	seller.Name = "seller测试update"
	seller.Password = "123456"
	seller.AuthImg = "authimg"
	AddSeller(&seller, &ctx)
	seller.Password = "654321"
	UpdateSeller(&seller, &ctx)
}
func TestDelSeller(t *testing.T) {
	var seller types.Seller
	seller.Name = "seller测试del"
	seller.Password = "123456"
	seller.AuthImg = "authimg"
	AddSeller(&seller, &ctx)
	DelSeller(&seller, &ctx)
}
func TestAddData(t *testing.T) {
	var category types.Category
	var shop types.Shop
	var item types.Item
	category.Name = "测试golang"
	AddCategory(&category, &ctx)
	shop.Name = "shop测试Del"
	shop.Lat = 1
	shop.Lon = 2
	shop.Img = "img"
	shop.CategoryId = category.Id
	shop.Star = 200
	shop.SellerId = 2
	AddShop(&shop, &ctx)
	item.Name = "item测试Del"
	item.ShopId = shop.Id
	item.Price = 100.123456
	item.Img = "img"
	AddItem(&item, &ctx)
}

func TestAddOrder(t *testing.T) {
	var order types.Order
	order.Price = 100.123456
	order.UserId = 1
	order.Ts = time.Now().Unix()
	fmt.Println(order.Ts)
	order.Cancel = 0
	AddOrder(&order, nil)
}

func TestAddItemCategory(t *testing.T) {
	var itemCategory types.ItemCategory
	itemCategory.Name = "测试itemCategory"
	itemCategory.ShopId = 1
	AddItemCategory(&itemCategory, nil)
}
func TestUpdateItemCategory(t *testing.T) {
	var itemCategory types.ItemCategory
	itemCategory.Name = "测试update itemCategory"
	itemCategory.ShopId = 1
	AddItemCategory(&itemCategory, nil)
	itemCategory.ShopId = 2
	UpdateItemCategory(&itemCategory, nil)
}
func TestDelItemCategory(t *testing.T) {
	var itemCategory types.ItemCategory
	itemCategory.Name = "测试del itemCategory"
	itemCategory.ShopId = 1
	AddItemCategory(&itemCategory, nil)
	DelItemCategory(&itemCategory, nil)
}
