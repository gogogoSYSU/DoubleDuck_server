//设置路由

package service

import (
	//"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON:true,
	})
	//negroni.Classic() provides some default middleware
	n := negroni.Classic()
	//register a couple of URL paths and handlers	
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	//negroni use mx to handle different routes
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	//handle url for homepage
	initRTRoutes(mx, formatter)
	initSalerRoutes(mx, formatter)
}

//饭店部分
func initRTRoutes(mx *mux.Router, formatter *render.Render){
	//显示饭店信息
	mx.HandleFunc("/v1/rt/", showRTinfoHandle(formatter)).Methods("GET")
	mx.HandleFunc("/v1/dish/", showRTdishHandle(formatter)).Methods("GET")
	//上传订单
	mx.HandleFunc("/v1/order", postOrderHandle(formatter)).Methods("Post")
}

//商家部分
func initSalerRoutes(mx *mux.Router, formatter *render.Render) {
	//创建新用户
	mx.HandleFunc("/v1/salers", registerSalerHandle(formatter)).Methods("Post")

	//用户登陆
	mx.HandleFunc("/v1/salers/login", loginSalerHandle(formatter)).Methods("Post")

	//上传饭店信息
	mx.HandleFunc("/v1/salers/info", uploadInfoHandle(formatter)).Methods("Post")
	//上传菜品信息
	mx.HandleFunc("/v1/salers/dish", uploadDishHandle(formatter)).Methods("Post")
	//新建菜品种类
	mx.HandleFunc("/v1/salers/cate", createCateHandle(formatter)).Methods("Post")
}
