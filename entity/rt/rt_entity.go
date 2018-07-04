/********************************
restaurant的数据层，保存基本信息
*********************************/

package rt 

import (
)

// Dish 单个菜品的信息
type Dish struct {
	//菜品名称 不会有相同的
	DishName string `json:"dishname" bson:"_id"`
	//菜品描述
	DishDes string `json:"dishdes"`
	//菜品价格
	DishPrice float64 `json:"dishprice"`
	//size? 
	//菜品照片 url
	DishPicture string `json:"dishpicture_url"`
	//菜品已售出数量
	DishSales int `json:"dishsales"`
	//菜品所属种类
	DishCategory string `json:"category" bson:"category"` 
	//所属餐厅
	DishBelong string `json:"belongrt"`
}
func newDish(name string, des string, price float64, pic string, sale int, cat string, rt string) *Dish{
	newdish := new(Dish)
	newdish.DishBelong = rt
	newdish.DishCategory = cat
	newdish.DishDes = des
	newdish.DishName = name
	newdish.DishPicture = pic
	newdish.DishPrice = price
	newdish.DishSales = sale
	return newdish
}

// Rt 每个餐厅的信息
type Rt struct {
	RtName string `json:"rtname" bson:"_id"` 
	RtLocation string `json:"rtlocation" bson:"rtlocation"`
	RtDes string `json:"rtdes" bson:"rtdes"`
	RtLogo string `json:"rtlogo_url" bson:"rtlogo"`
	RtPhone string `json:"rtphone" bson:"rtphone"`
	//RtDishes []Dish `json:"rtdishes"`
	RtCategories []string `json:"rtcategories" bson:"rtcate"`
	
}

func newRT(name string, loc string, des string, logo string, phone string, cate []string) *Rt{
	newrt := new(Rt)
	newrt.RtName = name
	newrt.RtLocation = loc
	newrt.RtDes = des
	newrt.RtLogo = logo
	newrt.RtCategories = cate
	newrt.RtPhone = phone
	return newrt
}
// Orderitem 每个订单菜品的信息
type Orderitem struct {
	ItemName string `json:"itemname"`
	//ItemSize?
	ItemAmount int `json:"itemamount"`
}

// Order 每个订单的信息
type Order struct {
	OrderID int `json:"orderid"`
	OrderDesk int `json:"orderdesk"`
	//order_number?
	OrderPrice float64 `json:"orderprice"`
	OrderIspayed bool `json:"orderispayed"`
	OrderItems []Orderitem `json:"orderitems"`
}

// DishInfo 菜品信息
type DishInfo struct {
	DishName string `json:"dishname"`
	DishLogo string `json:"dishpict"`
	DishSale int `json:"dishsale"`
	DishPrice float64 `json:"dishprice"`
	DishDis string `json:"dishdis"`
}
// DishInfos 一类菜品的信息
type DishInfos struct {
	Category string `json:"category"`
	DishesNum int `json:"dishesnum"`
	DishesList []DishInfo `json:"disheslist"`
}