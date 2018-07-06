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
	service.InsertRT(newRT("RT1","LOC1","DES1","LOGO1","phone1",array))
//	service.InsertDish(newDish("dish1","des1",22.8, "dishurl1", 1, "su", "RT1"))
//	service.InsertDish(newDish("dish2", "des2", 22.1, "url2", 2, "rou", "RT1"))
}


/***************************************
rt数据库，对每个餐厅的菜品table查询修改
***************************************/
// CheckDish 检查菜品是否重名了
func (*RTService) CheckDish(rtname string, dishname string) bool{
	c := database.C(rtname)
	dish := Dish{}
	err := c.Find(bson.M{"_id":dishname}).One(&dish)
	exist := true
	if err != nil {
		exist = false
		fmt.Println("上传菜品信息的名字没有重名,报错信息应该没有错")
		//panic(err)
	}
	return exist
}

// InsertDish 添加一条菜品信息至所属餐厅的菜品table
func (*RTService) InsertDish(dish *Dish) bool{
	//切换到所属集合，包含该餐厅菜品的表
	c := database.C(dish.DishBelong)
	err := c.Insert(dish)
	ifok := true
	if err != nil {
		fmt.Println("insert dish fail")
		ifok = false
		panic(err)
	}
	return ifok
}

// FindByCategory 通过所属种类找到菜品，为前端方便？？？
func (*RTService) FindByCategory(cate string, rt string) ([]DishInfo, int){
	c := database.C(rt)
	dishes := []Dish{}
	err := c.Find(bson.M{"category":cate}).All(&dishes)
	if err != nil {
		fmt.Println(cate)
		fmt.Println("该种类没有对应菜品，find的err应该不是空的")
		//panic(err)
	}

//	fmt.Println(dishes)

	dishesinfo := make([]DishInfo,len(dishes))
	for i := 0; i < len(dishes); i++ {
		dishesinfo[i].DishDis = dishes[i].DishDes
		dishesinfo[i].DishLogo = dishes[i].DishPicture
		dishesinfo[i].DishName = dishes[i].DishName
		dishesinfo[i].DishPrice = dishes[i].DishPrice
		dishesinfo[i].DishSale = dishes[i].DishSales
	}
	return dishesinfo, len(dishes)
}


/***************************************
rt数据库，对餐厅信息table查询修改
***************************************/
// CheckRT 检查餐厅是否已经被储存到餐厅信息的table，true表示在表中
func (*RTService) CheckRT(name string) bool {
	c := database.C("rt_table")
	rt := Rt{}
	err := c.Find(bson.M{"_id":name}).One(&rt)
	if err != nil {
		return false
	}
	return true
}
// InsertRT 添加一条餐厅信息到储存餐厅信息的table
func (*RTService) InsertRT(rt *Rt) bool{
	c := database.C("rt_table")
	err := c.Insert(rt)
	ifok := true
	if err != nil {
		ifok = false
		fmt.Println("insert rt info fail")
		panic(err)
	}
	fmt.Println("insert rt info success")
	return ifok
}
// UpdateRT 更改餐厅基本信息
func (*RTService) UpdateRTInfo(name string, loc string, des string, logo string, phone string) bool {
	c := database.C("rt_table")
	selector := bson.M{"_id": name}
	data := bson.M{"$set":bson.M{"rtlocation":loc, "rtdes":des, "rtlogo":logo, "rtphone":phone}}
	err := c.Update(selector, data)
	ifok := true
	if err != nil {
		fmt.Println("update rt info fail")
		ifok = false
		panic(err)
	}
	fmt.Println("update rt info success")
	return ifok
}
// FindCateByRT 根据餐厅名字查询种类
func (*RTService) FindCateByRT(rtname string) []string {
	c := database.C("rt_table")
	rt := Rt{}
	err := c.Find(bson.M{"_id":rtname}).One(&rt)
//	err := c.Find(nil).One(&rt)
	if err != nil {
		panic(err)
	}
	fmt.Println("find cate by rt")
	fmt.Println(rt)	
	return rt.RtCategories
}
// FindDesLocByRT 根据餐厅名字查询餐厅简介
func (*RTService) FindDesLocByRT(rtname string) (string,string,string,string) {
	c := database.C("rt_table")
	rt := Rt{}
	err := c.Find(bson.M{"_id":rtname}).One(&rt)
	if err != nil {
		panic(err)
	}
	return rt.RtDes, rt.RtLocation, rt.RtLogo, rt.RtPhone
} 

// CreateCateForRT 给餐厅添加参品种类
func (*RTService) CreateCateForRT(rtname string, cate string) bool {
	ifok := true
	c := database.C("rt_table")
	err := c.Update(
		bson.M{"_id":rtname},
		bson.M{"$push": bson.M{"rtcate":cate}})
	if err != nil {
		ifok = false
		panic(err)
	}
	return ifok
}
//更多操作待完善
