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
