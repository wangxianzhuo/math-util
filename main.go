package main

import (
	"github.com/astaxie/beego"
	"github.com/wangxianzhuo/math-util/controllers"
)

func main() {
	beego.Router("/det/value", &controllers.DetController{}, "get:Value")
	beego.Router("/det/expanse", &controllers.DetController{}, "get:Expanse")
	beego.Router("/equations/solve", &controllers.EquationsController{}, "get:Solve")
	beego.Router("/util/inversionnumber", &controllers.UtilController{}, "get:InversionNumber")
	beego.ErrorController(&ErrorController{})
	beego.SetLogFuncCall(true)
	beego.Run()
}

type ErrorController struct {
	beego.Controller
}

type ErrMessage struct {
	Message string
}

func (e ErrorController) Error500() {
	e.Data["json"] = ErrMessage{
		Message: "internal server error",
	}
	e.ServeJSON()
}

func (e ErrorController) Error404() {
	e.Data["json"] = ErrMessage{
		Message: "not found",
	}
	e.ServeJSON()
}
