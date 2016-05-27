package baseController

import (
	"github.com/astaxie/beego"
	"todos/models"
)

type (
	// BaseController composes all required types and behavior.
	BaseController struct {
		beego.Controller
		models.Service
	}
)

//** INTERCEPT FUNCTIONS

// Prepare is called prior to the baseController method.
func (baseController *BaseController) Prepare() {
	if err := baseController.Service.Prepare(); err != nil {
		return
	}
}

//** AJAX SUPPORT

// AjaxResponse returns a standard ajax response.
//func (baseController *BaseController) AjaxResponse(resultCode int, resultString string, data interface{}) {
//	response := struct {
//		Result       int
//		ResultString string
//		ResultObject interface{}
//	}{
//		Result:       resultCode,
//		ResultString: resultString,
//		ResultObject: data,
//	}

//	baseController.Data["json"] = response
//	baseController.ServeJson()
//}
