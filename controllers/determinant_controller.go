package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/wangxianzhuo/math-util/determinant"
)

type DetController struct {
	beego.Controller
}

// Value 计算行列式值
//
// URL:		/det/value
// Method:	GET
// Param:	elements=[[1,2,3],[4,5,6],[7,8,9]]
// Return:	{"value":0,"des": ""}
func (dc *DetController) Value() {
	var elements [][]float64

	beego.Debug("elements: ", dc.GetString("elements"))
	err := json.Unmarshal([]byte(dc.GetString("elements")), &elements)
	if err != nil {
		beego.Error("Calculate determinant parse form value error ", err)
		dc.Abort("500")
		return
	}
	det, err := determinant.NewDeterminantWithValueParallel(elements)

	if err != nil {
		beego.Error("Calculate determinant parse form value error ", err)
		dc.Abort("500")
		return
	}
	dc.Data["json"] = CalcResult{Value: det.Value}
	dc.ServeJSON()
}

// Expanse 行列式展开
//
// URL:		/det/expanse
// Method:	GET
// Param:	elements=[[1,2,3],[4,5,6],[7,8,9]]&order=0&cowRowType=0
// Return:	{"value":[{"Elements":[[5,6],[8,9]],"Rank":2,"Factor":1},{"Elements":[[4,6],[7,9]],"Rank":2,"Factor":-2},{"Elements":[[4,5],[7,8]],"Rank":2,"Factor":3}],"des":""}
func (dc *DetController) Expanse() {
	var elements [][]float64

	beego.Debug("elements: ", dc.GetString("elements"))
	err := json.Unmarshal([]byte(dc.GetString("elements")), &elements)
	if err != nil {
		beego.Error("Expanse determinant parse form value error ", err)
		dc.Abort("500")
		return
	}
	order, err := dc.GetInt("order")
	if err != nil {
		beego.Error("Expanse determinant parse form value error ", err)
		dc.Abort("500")
		return
	}
	rowCowType, _ := dc.GetInt("rowCowType")

	det, err := determinant.NewDeterminant(elements)
	if err != nil {
		beego.Error("Expanse determinant parse form value error ", err)
		dc.Abort("500")
		return
	}

	dets, err := det.Expanse(order, determinant.RowCowType(rowCowType))
	if err != nil {
		beego.Error("Expanse determinant error ", err)
		dc.Abort("500")
		return
	}

	dc.Data["json"] = CalcResult{Value: dets}
	dc.ServeJSON()
}
