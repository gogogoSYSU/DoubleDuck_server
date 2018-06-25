package service

import (
	//"errors"
	"fmt"
	"net/http"
	
	"github.com/gogogoSYSU/DoubleDuck_server/entity/rt"
	"github.com/gogogoSYSU/DoubleDuck_server/util"
	"github.com/unrolled/render"
)
// ErrorRtnJson 包含错误码和错误信息
type ErrorRtnJson struct {
	// 错误码
	Errorcode int `json:"errorcode"`
	// 错误信息
	Errorinformation string `json:"errorinformation"`
}
// 返回错误表单
func errResponse(w http.ResponseWriter, formatter *render.Render) {
	if err := recover(); err != nil {
		fmt.Println(err)
		var rtn ErrorRtnJson
		rtn.Errorcode, rtn.Errorinformation = util.HandleError(err)
		formatter.JSON(w, 500, rtn)
	}
}

// RtInfoRtJson 返回餐厅信息
type RtInfoRtJson struct {
	RtName string `json:"rtname"`
	RtDes string `json:rtdes`
	RtLoc string `json:rtloc`
	DishedInfo []rt.Dish `json:rtdishes`
}

func showRTinfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("showRtInfoHandle")

		//解析url参数数据
		param := parseURL(r)
		rtname := param["rtname"]
		fmt.Println(rtname)
		rtjson := RtInfoRtJson{}
		rtjson.RtName = rtname
		rtjson.RtDes, rtjson.RtLoc = rt.GetRtDesLoc(rtname)
	
		formatter.JSON(w, http.StatusOK, rtjson)
	}
}