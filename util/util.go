/**************************
常用函数
**************************/

package util

import (
	//"errors"
	"strings"
	"strconv"
)

// HandleError 处理错误，并且返回成int，string形式
func HandleError(err interface{}) (int, string) {
	strs := strings.Split(err.(error).Error(), "|")
	if len(strs) != 2 {
		return 200, "未定义错误"
	}
	errcode, err := strconv.Atoi(strs[0])
	if err != nil {
		return 200, "未定义错误"
	}
	return errcode, strs[1]
}
