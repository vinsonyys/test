package main

import (
	"github.com/astaxie/beego"
	_ "myproject/routers" // 默认是调用包里的init（）方法，不是调用包的函数
)

func main() {
	beego.Run()
}
