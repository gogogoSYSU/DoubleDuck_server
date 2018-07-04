package service

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gogogoSYSU/DoubleDuck_server/entity/saler"
	"github.com/gogogoSYSU/DoubleDuck_server/entity/rt"
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

func uploadInfoHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("uploadInfoHandle!")

		//解析参数
		param := parsePost(r)

		i := rt.UploadRtInfo(param["rtname"], param["rtlocation"], param["rtdes"],
			 param["rtlogo"], param["rtphone"])

		if i == false {
			fmt.Println("上传用户信息失败")
			formatter.JSON(w, http.StatusInternalServerError, SalerRtnJson{
				State:"UploadRTInfoFail",
				})
		} else {
			fmt.Println("上传用户信息成功")
			formatter.JSON(w, http.StatusOK, SalerRtnJson{
				State:"UploadRTInfoSuccess",
				})
		}
	}
}
//上传菜品，这里注意，上传菜品的种类肯定在上传菜品前就建立了
func uploadDishHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("uploadDinshHandle!")

		param := parsePost(r)
		price, err := strconv.ParseFloat(param["price"], 64)
		if err != nil {
			panic(err)
		}
		sale, err := strconv.Atoi(param["sale"])
		if err != nil {
			panic(err)
		}
		i := rt.UploadDish(param["name"], param["des"], price, param["pic"], sale, param["cate"],param["rt"])
		if i == false {
			fmt.Println("上传菜品信息失败")
			formatter.JSON(w, http.StatusInternalServerError, SalerRtnJson{
				State:"UploadDishFail",
				})
		} else {
			fmt.Println("上传用户信息成功")
			formatter.JSON(w, http.StatusOK, SalerRtnJson{
				State:"UploadDishSuccess",
			})
		}
	}
}
//新建菜品种类
func createCateHandle(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errResponse(w, formatter)
		fmt.Println("createCateHandle!")

		param := parsePost(r)
		rtname := param["rtname"]
		cate := param["cate"]
		
		succeed := rt.CreateCate(rtname, cate)
		if succeed == false {
			fmt.Println("添加菜品种类失败")
			formatter.JSON(w, http.StatusInternalServerError, SalerRtnJson{
				State:"createCateFail",
			})
		} else {
			fmt.Println("添加菜品种类成功")
			formatter.JSON(w, http.StatusOK, SalerRtnJson{
				State:"createCateSuccess",
			})
		}
	}
}





