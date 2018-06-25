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