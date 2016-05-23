package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	_ "todos/routers"
)

func main() {
	beego.Run()
	log := logs.NewLogger(10000)
	log.SetLogger(logs.AdapterFile, `{"filename":"bg_mess.log"}`)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(3)
}
