package routers

import (
	"github.com/astaxie/beego"
	"todos/controllers"
)

func init() {
	// Set Static Paths to bower and npm
	beego.SetStaticPath("/bower", "bower_components")

	// Images
	beego.Router("/api/images", &controllers.ImageController{})
	// Actors as test of DB queries
	beego.Router("/api/actors", &controllers.ActorsController{})
	// Route to index.tpl
	beego.Router("/", &controllers.MainController{})
}
