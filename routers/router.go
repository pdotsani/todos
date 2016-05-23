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
	beego.Router("/actors", &controllers.ActorsController{})
	beego.Router("/movies", &controllers.MoviesController{})
	beego.Router("/directors", &controllers.DirectorsController{})
}
