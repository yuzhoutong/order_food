package main

import (
	"github.com/astaxie/beego"
	_ "order_food/routers"
	"order_food/service"
)

func main() {


	go service.AutoSaveToDate()
	beego.Run()
}
