package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/wangxianzhuo/math-util/determinant"
)

type DetController struct {
	beego.Controller
}

func (this *DetController) Value() {
	var elements [][]float64

	beego.Debug("elements: ", this.GetString("elements"))
	err := json.Unmarshal([]byte(this.GetString("elements")), &elements)
	if err != nil {
		errHandler(ErrMessage{
			Status:  http.StatusInternalServerError,
			Message: "Calculate determinant parse form value error",
			Err:     err,
		}, this.Ctx.ResponseWriter)
		return
	}
	det, err := determinant.NewDeterminantWithValueParallel(elements)
	if err != nil {
		errHandler(ErrMessage{
			Status:  http.StatusInternalServerError,
			Message: "Calculate determinant parse form value error",
			Err:     err,
		}, this.Ctx.ResponseWriter)
		return
	}

	this.Data["json"] = CalcResult{Value: det.Value}
	this.ServeJSON()
}

func (this *DetController) Expanse() {
	var elements [][]float64

	beego.Debug("elements: ", this.GetString("elements"))
	err := json.Unmarshal([]byte(this.GetString("elements")), &elements)
	if err != nil {
		errHandler(ErrMessage{
			Status:  http.StatusInternalServerError,
			Message: "Expanse determinant parse form value error",
			Err:     err,
		}, this.Ctx.ResponseWriter)
		return
	}
	order, err := this.GetInt("order")
	if err != nil {
		errHandler(ErrMessage{
			Status:  http.StatusInternalServerError,
			Message: "Expanse determinant parse form value error",
			Err:     err,
		}, this.Ctx.ResponseWriter)
		return
	}
	rowCowType, _ := this.GetInt("rowCowType")

	det, err := determinant.NewDeterminant(elements)
	if err != nil {
		errHandler(ErrMessage{
			Status:  http.StatusInternalServerError,
			Message: "Expanse determinant parse form value error",
			Err:     err,
		}, this.Ctx.ResponseWriter)
		return
	}

	dets, err := det.Expanse(order, determinant.RowCowType(rowCowType))
	if err != nil {
		errHandler(ErrMessage{
			Status:  http.StatusInternalServerError,
			Message: "Expanse determinant error",
			Err:     err,
		}, this.Ctx.ResponseWriter)
		return
	}

	this.Data["json"] = dets
	this.ServeJSON()
}
