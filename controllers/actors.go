package controllers

import (
	bc "todos/controllers/baseController"
)

type ActorsController struct {
	bc.BaseController
}

func (actorsController *ActorsController) Post() {
	actor := actorsController.GetString("actor")
	if actor == "" {
		actorsController.Ctx.WriteString("no actor listed")
		return
	}
	actorsController.QueryAllNodes(actor)
}

func (actorsController *ActorsController) Get() {
	actorsController.QueryAllNodes("William Shatner")
}
