/********************************
服务端结构，用于json返回
********************************/

package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"errors"

	simplejson "github.com/bitly/go-simplejson"
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

func parsePost(r *http.Request) map[string]string {
	// 解析参数
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
                panic(errors.New("203|解析url参数错误"))
	}
	//CheckNewErr(err, "203|解析json错误")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
                fmt.Fprintln(os.Stderr, err)
                panic(errors.New("203|解析url参数错误"))
        }
	//CheckNewErr(err, "203|解析json错误")
	defer r.Body.Close()

	// 解析json，转换成map
	temp, err := simplejson.NewJson(body)
	if err != nil {
                fmt.Fprintln(os.Stderr, err)
                panic(errors.New("203|解析url参数错误"))
        }
	//CheckNewErr(err, "203|解析json错误")
	m, _ := temp.Map()
	if err != nil {
                fmt.Fprintln(os.Stderr, err)
                panic(errors.New("203|解析url参数错误"))
        }
	//CheckNewErr(err, "203|解析json错误")
	rtnmap := make(map[string]string)
	for k, v := range m {
		rtnmap[k] = fmt.Sprint(v)
	}
	return rtnmap
}
