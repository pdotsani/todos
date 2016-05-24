package controllers

import (
	"github.com/astaxie/beego"
	"todos/services"
)

type MainController struct {
	beego.Controller
	services.Service
<<<<<<< HEAD
}

type ActorsController struct {
	beego.Controller
	services.Service
}

type MoviesController struct {
	beego.Controller
	services.Service
}

type DirectorsController struct {
	beego.Controller
	serices.Service
=======
>>>>>>> master
}

type ActorsController struct {
	beego.Controller
	services.Service
}

func (mainController *MainController) Prepare() {
	if err := mainController.Service.Prepare(); err != nil {
		return
	}
}

func (actorsController *ActorsController) Prepare() {
	if err := actorsController.Service.Prepare(); err != nil {
		return
	}
}

func (mainController *MainController) Get() {
	mainController.TplName = "index.tpl"
}

func (actorsController *ActorsController) Get() {
	actorsController.TplName = "index.tpl"
}

func (c *ActorsController) Get() {
	c.TplName = "index.tpl"
}

func (c *MoviesController) Get() {
	c.TplName = "index.tpl"
}

func (c *DirectorsController) Get() {
	c.TplName = "index.tpl"
}
