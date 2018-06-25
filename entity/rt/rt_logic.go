/*****************************************
RT的逻辑层，使用dao层的接口，为server提供接口
*****************************************/

package rt

import (
	//"errors"
)

func GetRtDesLoc(rtname string) (string,string) {
	return service.FindDesLocByRT(rtname)
}