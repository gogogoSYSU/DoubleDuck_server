/*****************************************
saler的dao层，使用数据库接口，为上层逻辑提供接口
******************************************/

package saler

import (
	"fmt"
	"errors"
	"os"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/gogogoSYSU/DoubleDuck_server/mangodb"
)

//saler 使用一个数据库，所有saler存储在同一张table
var collector *mgo.Collection

type SalerService struct{}
var service = SalerService{}

func init() {
	collector = mangodb.Mydb.DB("saler").C("saler_table")
	fmt.Println("saler database init")
}

// Insert 插入新的saler信息
func (*SalerService) Insert(saler *SalerInfo) {
	err := collector.Insert(saler)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(errors.New("insert saler info出错"))
	}
}

// FindPwByID 根据id找到密码，用来验证输入密码是否正确
func (*SalerService) FindPwByID(ID string) (string,bool) {
	saler := &SalerInfo{}
	err := collector.Find(bson.M{"_id":ID}).One(saler)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(errors.New("数据库根据id查询pw出错|没有这个id"))
		return "", false
	}
	return saler.Password, true
}

//更多操作待之后需要再写