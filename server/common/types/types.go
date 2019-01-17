package types

type Position struct {
	Lat float64
	Lon float64
}

type Shop struct {
	Id         int64   `db:"id"`
	Name       string  `db:"name"`
	Lat        float64 `db:"lat"`
	Lon        float64 `db:"lon"`
	Img        string  `db:"img"`
	ItemList   []Item
	CategoryId uint32 `db:"categoryid"`
	Star       uint32 `db:"star"`
	SellerId   uint32 `db:"sellerid"`
}

type Category struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	ShopList []Shop
}

type Item struct {
	Id     int64   `db:"id"`
	Name   string  `db:"name"`
	Price  float32 `db:"price"`
	Img    string  `db:"img"`
	ShopId uint32  `db:"shopid"`
}

type Seller struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	AuthImg  string `db:"authimg"`
}

type Context struct{}
