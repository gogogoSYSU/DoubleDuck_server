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
	State string `json:"state"`
//	ErrorRtnJson
}

//注册创建一个新用户
func registerSalerHandle(formatter *render.Render) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("registerSalerHandle!")

		//解析参数
		param := parsePost(r)
		saler.RegisterSaler(param["openID"], param["pw"], param["rtname"])

		//发回消息
		formatter.JSON(w, http.StatusOK, ErrorRtnJSON{})
	}
}

//登陆
func loginSalerHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("loginHandle!")

		//解析参数
		param := parsePost(r)

		temp := saler.GetSalerPassword(param["openID"])
		fmt.Println(temp + " " + param["pw"])
		if temp == param["pw"] {
			temp = "OK"
		} else {
			temp = "Fail"
		}

		fmt.Println(temp)

		formatter.JSON(w, http.StatusOK, SalerRtnJson{
			State:temp,
			})
	}
}
