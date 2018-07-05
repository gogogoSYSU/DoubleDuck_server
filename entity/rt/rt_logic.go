/*****************************************
RT的逻辑层，使用dao层的接口，为server提供接口
*****************************************/

package rt

import (
	//"errors"
)

// GetRtBasicInfo 提供餐厅基础信息
func GetRtBasicInfo(rtname string) (string,string,string,string) {
	return service.FindDesLocByRT(rtname)
}

// GetRtDishes 提供餐厅所有菜品的信息
func GetRtDishes(rtname string) []DishInfos {
	cate := service.FindCateByRT(rtname)
	rtdishes := make([]DishInfos, len(cate))
	//对每一类菜品分析
	for i := 0; i < len(cate); i++ {
		rtdishes[i].Category = cate[i]
		rtdishes[i].DishesList, rtdishes[i].DishesNum = service.FindByCategory(cate[i], rtname)
	}
	return rtdishes
}

// UploadRtInfo 卖家上传餐厅信息
func UploadRtInfo(name string, loc string, des string, logo string, phone string) bool{
	exist := service.CheckRT(name)
	ifok := true
	if exist != true {
		cate := []string{}
		rt := newRT(name, loc, des, logo, phone, cate)
		ifok = service.InsertRT(rt)
	} else {
		ifok = service.UpdateRTInfo(name, loc, des, logo, phone)
	}
	return ifok
}

// UploadDish 卖家上传单个菜品信息
func UploadDish(name string, des string, price float64, pic string, sale int, cat string, rt string) bool {
	exist := service.CheckDish(rt, name)
	ifok := false
	if exist == false  {
		dish := newDish(name, des, price, pic, sale, cat, rt)
		ifok = service.InsertDish(dish)
	}
	return ifok
}

// CreateCate 卖家添加菜品种类
func CreateCate(rt string, cate string) bool {
	ifok := true
	//先检查菜品种类是否重名
	cates := service.FindCateByRT(rt)
	for i := 0; i < len(cates); i++ {
		if cates[i] == cate {
			return false
		}
	}
	ifok = service.CreateCateForRT(rt, cate)
	return ifok
}