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
	DishPrice float32 `json:"dishprice"`
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

// Rt 每个餐厅的信息
type Rt struct {
	RtName string `json:"rtname" bson:"_id"` 
	RtLocation string `json:"rtlocation"`
	RtDes string `json:"rtdes"`
	RtLogo string `json:"rtlogo_url"`
	//RtDishes []Dish `json:"rtdishes"`
	RtCategories []string `json:"rtcategories"`
	
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