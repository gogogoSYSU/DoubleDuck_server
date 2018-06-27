package service

import (
//	"errors"
	"fmt"
	"net/http"

	"github.com/gogogoSYSU/DoubleDuck_server/entity/saler"
//	"github.com/gogogoSYSU/DoubleDuck_server/util"
	"github.com/unrolled/render"
)


//SalerReturnJSON 用于返回商家信息
type SalerRtnJson struct {
	Password string `json:"pw"`
//	ErrorRtnJson
}

//注册创建一个新用户
func registerSalerHandle(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("registerSalerHandle!")

		//解析参数
		param := parseURL(r)
		saler.RegisterSaler(param["openID"], param["pw"], param["rtname"])

		//发回消息
		formatter.JSON(w, http.StatusOK, ErrorRtnJSON{})
	}
}

//显示商家信息
func listSalerInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("listSalerInfoHandle!")

		//解析参数
		param := parseURL(r)

		temp := saler.GetSalerPassword(param["openID"])

		if temp == param["password"] {
			temp = "OK"
		} else {
			temp = "Fail"
		}

		formatter.JSON(w, http.StatusOK, SalerRtnJson{
			Password:temp,
			})
	}
}
