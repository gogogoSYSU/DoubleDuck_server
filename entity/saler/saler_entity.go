/*********************************
saler的数据层，商家端要用
*********************************/

package saler

type SalerInfo struct {
	ID string `json:"openID" bson:"_id"`
	Password string 
}/*********************************
saler的数据层，商家端要用
*********************************/

package saler
// SalerInfo 商家基本信息
type SalerInfo struct {
	// 用户id，唯一的
	ID string `json:"openID" bson:"_id"`
	// 用户密码
	Password string `json:"pw"`
	// 用户的餐厅
	Restaurant string `json:"rt_name"`
}

func newSaler(ID string, pw string, rt string) *SalerInfo {
	newSaler := new(SalerInfo)
	newSaler.ID = ID
	newSaler.Password = pw
	newSaler.Restaurant = rt
	return newSaler
}