/*****************************************
saler的dao层，使用数据库接口，为上层逻辑提供接口
******************************************/

package saler

import (
	"fmt"
	"errors"
//	"sync"
	"os"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/gogogoSYSU/DoubleDuck_server/mangodb"
)

//saler 使用一个数据库，所有saler存储在同一张table
var collector *mgo.Collection

//空类型
type SalerService struct{}

//service 空类型的指针，使用函数
var service = SalerService{}

//用于临界区的同步互斥
//var lock = *sync.RWMutex

func init() {
	collector = mangodb.Mydb.DB("saler").C("saler_table")
	//lock = new(sync.RWMutex)
	fmt.Println("saler database init")
//	service.Insert(newSaler("lianqy", "lianqy", "xixi"))
}

// Insert 插入新的saler信息
func (*SalerService) Insert(saler *SalerInfo) {
	//lock.Lock()
	err := collector.Insert(saler)
	//lock.Unlock()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(errors.New("insert saler info出错"))
	}
}

// FindPwByID 根据id找到密码，用来验证输入密码是否正确
func (*SalerService) FindPwByID(ID string) (string) {
	saler := &SalerInfo{}

	//lock.Lock()

	err := collector.Find(bson.M{"_id":ID}).One(saler)

	//lock.Unlock()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(errors.New("数据库根据id查询pw出错|没有这个id"))
		return ""
	}
	return saler.Password
}

//更多操作待之后需要再写