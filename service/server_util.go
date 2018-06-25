/********************************
服务端结构，用于json返回
********************************/

package service

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
	"errors"
)

// 解析传过来的Url /?rt_name=XXX
func parseURL(r *http.Request) map[string]string {
	// 解析参数
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(errors.New("204|解析url参数错误"))
	}

	// 解析ID
	rtnmap := make(map[string]string)
	for k, v := range r.Form {
		rtnmap[k] = v[0]
	}
	fmt.Println("Get:", rtnmap)
	return rtnmap
}