//设置路由

package service

import (
	"net/http"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	//"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	//negroni.Classic() provides some default middleware
	n := negroni.Classic()
	//register a couple of URL paths and handlers	
	mx := mux.NewRouter()
	initRoutes(mx)
	//negroni use mx to handle different routes
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router) {
	//handle url for homepage
	mx.HandleFunc("/",func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "Welcome to the home page!\n")	
	})
 	//handle paths that have variables, defined using the format{name}
	mx.HandleFunc("/student/{name}", StudentHandler)
	mx.HandleFunc("/teacher/{name}", TeacherHandler)
}
//下面是之前作业的示例
func StudentHandler(w http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Student: %v\n", vars["name"])
}
func TeacherHandler(w http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Teacher: %v\n", vars["name"])
}
