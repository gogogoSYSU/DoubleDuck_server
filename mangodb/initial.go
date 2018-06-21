/*****************************
初始化数据库，和dao层对接
*****************************/

package mangodb

import (
	"labix.org/v2/mgo"
)

// Mydb 数据库指针 Session.DB()切换数据库
var Mydb * mgo.Session

func init() {
	//通过Dial()来与mongodb服务器建立连接，默认端口27017
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	
	//Optional.Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	Mydb = session
}