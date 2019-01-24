package types

type Position struct {
	Lat float64
	Lon float64
}

type Shop struct {
	Id               int64   `db:"id"`
	Name             string  `db:"name"`
	Lat              float64 `db:"lat"`
	Lon              float64 `db:"lon"`
	GeoHash          uint64  `db:"geohash"`
	Img              string  `db:"img"`
	CategoryId       int64   `db:"categoryid"`
	Star             uint32  `db:"star"`
	SellerId         int64   `db:"sellerid"`
	DispatchMin      float64 `db:"dispatch_min"`
	DispatchPrice    float64 `db:"dispatch_price"`
	ItemCategoryList []ItemCategory
}

type Category struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	ShopList []Shop
}

type ItemCategory struct {
	Id       int64  `db:"id"`
	Name     string `db:name`
	ShopId   int64  `db:"shopid"`
	ItemList []Item
}
type Item struct {
	Id             int64   `db:"id"`
	Name           string  `db:"name"`
	Price          float32 `db:"price"`
	Img            string  `db:"img"`
	ShopId         int64   `db:"shopid"`
	BarCode        string  `db:"barcode"`
	ItemCategoryId int64   `db:item_category_id`
}

type Seller struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	AuthImg  string `db:"authimg"`
}

type Order struct {
	Id       int64   `db:"id"`
	ItemsStr string  `db:"items"`
	Price    float32 `db:"price"`
	UserId   int64   `db:"userid"`
	Ts       int64   `db:"ts"`
	ItemList []Item
	Cancel   int `db:"cancel"`
}

type Context struct{}
