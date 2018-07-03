package service

import (
	"fmt"
	"net/http"

	"github.com/gogogoSYSU/DoubleDuck_server/entity/saler"
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
		result := saler.RegisterSaler(param["openID"], param["pw"], param["rtname"])

		//发回消息
		if result == false {
			fmt.Println("注册失败，该用户已经被注册")
			formatter.JSON(w, http.StatusOK, ErrorRtnJSON{
				Errorcode:100,
				Errorinformation:"用户已注册",
			})
		} else {
			formatter.JSON(w, http.StatusOK, ErrorRtnJSON{})
		}
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
		//fmt.Println(temp + " " + param["pw"])

		if temp == "fail" {
			fmt.Println("该用户不存在")
			formatter.JSON(w, http.StatusOK, SalerRtnJson{
				State:"noexit",
				})
		} else {
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
}
