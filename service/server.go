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
}

//饭店部分
func initRTRoutes(mx *mux.Router, formatter *render.Render){
	//显示饭店信息
	mx.HandleFunc("/v1/", showRTinfoHandle(formatter)).Methods("GET")
}