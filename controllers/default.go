package controllers

import (
	"github.com/astaxie/beego"
	"todos/services"
)

type MainController struct {
	beego.Controller
	services.Service
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
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
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
