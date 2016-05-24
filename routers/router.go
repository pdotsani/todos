package routers

import (
	"github.com/astaxie/beego"
	"todos/controllers"
)

func init() {
	// Set Static Paths to bower and npm
	beego.SetStaticPath("/bower", "bower_components")
	beego.SetStaticPath("/npm", "node_modules")

	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/explore",
		beego.NSNamespace("/actors",
			beego.NSInclude(
				&controllers.ActorsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
