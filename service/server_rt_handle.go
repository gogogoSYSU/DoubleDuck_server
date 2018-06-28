package service

import (
	//"errors"
	"fmt"
	"net/http"
	
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
