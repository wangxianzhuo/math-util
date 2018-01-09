package main

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/wangxianzhuo/math-util/controllers"
)

func main() {
	beego.Router("/det/value", &controllers.DetController{}, "get:Value")
	beego.Router("/det/expanse", &controllers.DetController{}, "get:Expanse")
	beego.ErrorHandler("404", err404)
	beego.SetLogFuncCall(true)
	beego.Run()
}

func err404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	http.Error(w, "{\"message\":\"not found\"}", http.StatusNotFound)
}
