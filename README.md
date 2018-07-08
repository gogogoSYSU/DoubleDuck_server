# README

> 双鸭点餐服务端(Created by @Caroline1997、@lianqy)

服务端使用go语言编写，运行前需使用一下命令进行环境的配置

  - `go get github.com/codegangsta/negroni`
  - `go get github.com/gorilla/mux`
  - `go get github.com/unrolled/render`
  - `go get github.com/bitly/go-simplejson`
  - `go get github.com/spf13/pflag`
  
数据库使用mongodb,因此要配置mgo

  - `go get gopkg.in/mgo.v2`
  
环境配置好后先启动mongodb服务，而后直接运行main.go即可运行

`service mongod start`

`go run main.go`
