package model

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"market-server/server/common/types"
)

type DataBase struct {
	SqlConn *dbr.Connection
}

func (database *DataBase) GetSession() *dbr.Session {
	return database.SqlConn.NewSession(nil)
}

var Database DataBase

func init() {
	conn, _ := dbr.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/market", nil)
	conn.SetMaxOpenConns(100)
	Database.SqlConn = conn
}

//get all near shop info, include items info, because
//data is not too big, so we can get these by one xfer
/*func GetNearShops(pos *types.Position, dist int, limit int, ctx *types.Context) []types.Category {

}*/

func jsonPrint(t interface{}) {
	data, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		fmt.Printf("jsonPrint err [%v]", err)
		return
	}
	fmt.Println(string(data))
}

func GetCategory(ctx *types.Context) ([]types.Category, error) {
	var categoryList []types.Category
	sess := Database.GetSession()
	sess.Select("*").From("category").Load(&categoryList)
	jsonPrint(categoryList)
	return categoryList, nil
}

func AddCategory(category *types.Category, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.InsertInto("category").Columns("name").Record(category).Exec()
	return nil
}

func UpdateCategory(category *types.Category, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.Update("category").Set("name", category.Name).Where("id=?", category.Id).Exec()
	return nil
}

func DelCategory(category *types.Category, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.DeleteFrom("category").Where("id=?", category.Id).Exec()
	return nil
}

func GetAllShop(ctx *types.Context) ([]types.Shop, error) {
	var shopList []types.Shop
	sess := Database.GetSession()
	sess.Select("*").From("shop").Load(&shopList)
	jsonPrint(shopList)
	return shopList, nil
}

func AddShop(shop *types.Shop, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.InsertInto("shop").Columns("name", "star", "categoryid", "img", "lat", "lon", "sellerid").
		Record(shop).
		Exec()
	jsonPrint(shop)
	return nil
}

func UpdateShop(shop *types.Shop, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.Update("shop").
		Set("name", shop.Name).
		Set("star", shop.Star).
		Set("categoryid", shop.CategoryId).
		Set("img", shop.Img).
		Set("lat", shop.Lat).
		Set("lon", shop.Lon).
		Set("sellerid", shop.SellerId).
		Where("id=?", shop.Id).
		Exec()

	jsonPrint(shop)
	return nil
}

func DelShop(shop *types.Shop, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.DeleteFrom("shop").Where("id=?", shop.Id).Exec()
	return nil
}

func GetAllItems(shop *types.Shop, ctx *types.Context) ([]types.Item, error) {
	var itemList []types.Item
	sess := Database.GetSession()
	sess.Select("*").From("item").Where("shopid=?", shop.Id).Load(&itemList)
	return itemList, nil
}

func AddItem(item *types.Item, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.InsertInto("item").
		Columns("name", "shopid", "price", "img").
		Record(item).
		Exec()
	return nil
}
func UpdateItem(item *types.Item, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.Update("item").
		Set("name", item.Name).
		Set("shopid", item.ShopId).
		Set("price", item.Price).
		Set("img", item.Img).
		Where("id=?", item.Id).
		Exec()
	return nil
}
func DelItem(item *types.Item, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.DeleteFrom("item").
		Where("id=?", item.Id).
		Exec()
	return nil
}
func AddSeller(seller *types.Seller, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.InsertInto("seller").
		Columns("name", "password", "authimg").
		Record(seller).
		Exec()
	return nil
}
func UpdateSeller(seller *types.Seller, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.Update("seller").
		Set("name", seller.Name).
		Set("password", seller.Password).
		Set("authimg", seller.AuthImg).
		Where("id=?", seller.Id).
		Exec()
	return nil
}
func DelSeller(seller *types.Seller, ctx *types.Context) error {
	sess := Database.GetSession()
	sess.DeleteFrom("seller").
		Where("id=?", seller.Id).
		Exec()
	return nil
}
