package saler

import (
//	"errors"
)

func RegisterSaler(openID string, pw string, rt_name string) (bool) {
	temp := newSaler(openID, pw, rt_name)

	isexit := service.Checkid(openID)

	if isexit == false {
		service.Insert(temp)
		return true
	} else {
		return false
	}
}

func GetSalerPassword(openID string) string {

	isexit := service.Checkid(openID)

	if isexit == false {
		return "fail"
	} else {
		return service.FindPwByID(openID)
	}
}

func GetSalerRtname(openID string) string {
	return service.FindRtnameByID(openID)
}
