package saler

import (
//	"errors"
)

func RegisterSaler(openID string, pw string, rt_name string) {
	temp := newSaler(openID, pw, rt_name)

	service.Insert(temp)
}

func GetSalerPassword(openID string) string {
	return service.FindPwByID(openID)
}
