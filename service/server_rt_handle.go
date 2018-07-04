package service

import (
	//"errors"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gogogoSYSU/DoubleDuck_server/entity/rt"
	"github.com/gogogoSYSU/DoubleDuck_server/util"
	"github.com/unrolled/render"
)
// ErrorRtnJSON 包含错误码和错误信息
type ErrorRtnJSON struct {
	// 错误码
	Errorcode int `json:"errorcode"`
	// 错误信息
	Errorinformation string `json:"errorinformation"`
}
// 返回错误表单
func errResponse(w http.ResponseWriter, formatter *render.Render) {
	if err := recover(); err != nil {
		fmt.Println(err)
		var rtn ErrorRtnJSON
		rtn.Errorcode, rtn.Errorinformation = util.HandleError(err)
		formatter.JSON(w, 500, rtn)
	}
}

// RtInfoRtJSON 返回餐厅信息
type RtInfoRtJSON struct {
	RtName string `json:"rtname"`
	RtDes string `json:"rtdes"`
	RtLoc string `json:"rtloc"`
	RtPhone string `json:"rtphone"`
	RtLogo string `json:"rtlogo"`
	//DishedInfo []rt.Dish `json:rtdishes`
}

// DishesJSON 返回全部菜品的信息
type DishesJSON struct {
	Alldishes []rt.DishInfos `json:"alldishes"`
}
// showRTinfoHandle 返回餐厅基本信息
func showRTinfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("showRtInfoHandle")

		//解析url参数数据
		param := parseURL(r)
		rtname := param["rtname"]
		fmt.Println(rtname)
		rtjson := RtInfoRtJSON{}
		rtjson.RtName = rtname
		rtjson.RtDes, rtjson.RtLoc, rtjson.RtLogo, rtjson.RtPhone = rt.GetRtBasicInfo(rtname)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("content-type", "application/json")
		formatter.JSON(w, http.StatusOK, rtjson)
	}
}
// showRTdishHandle 返回餐厅菜品信息
func showRTdishHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("showRTdishHandle")

		//解析url参数数据
		param := parseURL(r)
		rtname := param["rtname"]
		fmt.Println(rtname)
		rtjson := DishesJSON{}
		rtjson.Alldishes = rt.GetRtDishes(rtname)
		w.Header().Set("Access-Control-Allow-Origin", "*")
                w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
                w.Header().Set("content-type", "application/json")

		formatter.JSON(w, http.StatusOK, rtjson)
	}
}

//用来解析订单， 单个菜品
type DishItem struct {
	Name string `json:"dish_name"`
	Pic string `json:"dish_pict"`
	Sale int `json:"dish_sale"`
	Price float64 `json:"dish_price"`
	Dis string `json:"dish_discription"`
}
//用来解析订单， 整个订单
type OrderItem struct {
	RTname string `json:"storeName"`
	TotalPrice float64 `json:"totalPrice"`
	Dishes []DishItem `json:"selectedDish"`
}
//RtnJSON 用于返回商家信息
type RtnJSON struct {
	State string `json:"state"`
}
// 还没写进数据库，post数据需要再改，先验证
func postOrderHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		//定义好结构体直接去解析
		var temp OrderItem
		json.Unmarshal(parseOrder(r), &temp)
		fmt.Println("post订单详情：")	
		fmt.Println(temp.RTname)
		fmt.Println(temp.TotalPrice)
		fmt.Println(temp.Dishes)
		formatter.JSON(w, http.StatusCreated, RtnJSON{
			State:"post order success",
		})
	}	
}