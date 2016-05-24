package routers

import (
	"github.com/astaxie/beego"
	"todos/controllers"
)

func init() {
	// Set Static Paths to bower and npm
	beego.SetStaticPath("/bower", "bower_components")
	beego.SetStaticPath("/npm", "node_modules")

	// Create server api namespace
	// https://github.com/shlomizadok/beego_react/blob/master/routers/router.go
	api := beego.NewNamespace("/api",
		beego.NSNamespace("/images",
			beego.NSInclude(
				&controllers.ImageController{},
			),
		),
	)

	// Make api routes available for requests in app
	beego.AddNamespace(api)
	// Route to index.tpl
	beego.Router("/", &controllers.MainController{})
}
