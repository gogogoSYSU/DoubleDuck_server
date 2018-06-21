/********************************
restaurant的dao层,使用数据库接口，为上层逻辑提供接口
*********************************/

package rt 

import (
	"errors"
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/gogogoSYSU/DoubleDuck_server/mangodb"

)

// rt 记录每个餐厅的菜品使用同一个数据库，不同表
var database *mgo.Database
// RT Service
type RTService struct{}
//
var service = RTService{}

func init() {
	//切换到rt数据库
	database = mangodb.Mydb.DB("rt")
	fmt.Println("rt database init")
}

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
func (*RTService) FindByCatogory(cato string) []Dish {
	c := database.C(cato)
	dishes := []Dish{}
	err := c.Find(bson.M{"category":cato}).All(&dishes)
	if err != nil {
		panic(err)
	}
	return dishes
}

// InsertRT 添加一条餐厅信息到储存餐厅信息的table
