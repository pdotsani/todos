package routers

import (
	"github.com/astaxie/beego"
	"todos/controllers"
)

func init() {
	// Set Static Paths to bower and npm
	beego.SetStaticPath("/bower", "bower_components")
	beego.SetStaticPath("/npm", "node_modules")

	// Images
	beego.Router("/api/images", &controllers.ImageController{})
	// Route to index.tpl
	beego.Router("/", &controllers.MainController{})
}
