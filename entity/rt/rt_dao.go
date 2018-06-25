/********************************
restaurant的dao层,使用数据库接口，为上层逻辑提供接口
*********************************/

package rt 

import (
	//"errors"
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/gogogoSYSU/DoubleDuck_server/mangodb"

)

// rt 记录每个餐厅的菜品使用同一个数据库，不同表
var database *mgo.Database
// RTService 空类型
type RTService struct{}
//
var service = RTService{}

func init() {
	//切换到rt数据库
	database = mangodb.Mydb.DB("rt")
	fmt.Println("rt database init")
	array := [] string{"su","rou"}
	service.InsertRT(newRT("HAOCHI","menkou","hhh","url",array))
}


/***************************************
rt数据库，对每个餐厅的菜品table查询修改
***************************************/

// InsertDish 添加一条菜品信息至所属餐厅的菜品table
func (*RTService) InsertDish(dish *Dish) {
	//切换到所属集合，包含该餐厅菜品的表
	c := database.C(dish.DishBelong)
	err := c.Insert(dish)
	if err != nil {
		panic(err)
	}
}

// FindByCatogory 通过所属种类找到菜品，为前端方便？？？
func (*RTService) FindByCategory(cate string) []Dish {
	c := database.C(cate)
	dishes := []Dish{}
	err := c.Find(bson.M{"category":cate}).All(&dishes)
	if err != nil {
		panic(err)
	}
	return dishes
}


/***************************************
rt数据库，对餐厅信息table查询修改
***************************************/

// InsertRT 添加一条餐厅信息到储存餐厅信息的table
func (*RTService) InsertRT(rt *Rt) {
	c := database.C("rt_table")
	err := c.Insert(rt)
	if err != nil {
		panic(err)
	}
}
// FindCateByRT 根据餐厅名字查询种类
func (*RTService) FindCateByRT(rtname string) []string {
	c := database.C("rt_table")
	rt := Rt{}
	err := c.Find(bson.M{"_id":rtname}).One(&rt)
	if err != nil {
		panic(err)
	}
	return rt.RtCategories
}
// FindDesLocByRT 根据餐厅名字查询餐厅简介
func (*RTService) FindDesLocByRT(rtname string) (string,string) {
	c := database.C("rt_table")
	rt := Rt{}
	err := c.Find(bson.M{"_id":rtname}).One(&rt)
	if err != nil {
		panic(err)
	}
	return rt.RtDes, rt.RtLocation
} 
//更多操作待完善