package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/wangxianzhuo/math-util/determinant"
)

type UtilController struct {
	beego.Controller
}

// InversionNumber 求解序列的逆序数
//
// URL:		/util/inversionnumber
// Method:	GET
// Param:	list=[4,3,5,6,23,7,2,54,9]
// Return:	{"value": 10,"des": ""}
func (uc *UtilController) InversionNumber() {
	var list []int
	err := json.Unmarshal([]byte(uc.GetString("list")), &list)
	if err != nil {
		beego.Error("Calculate list", uc.GetString("list"), " inversion number error: ", err)
		uc.Abort("500")
	}

	value := determinant.CalculateInversionNumber(list)
	uc.Data["json"] = CalcResult{Value: value}
	uc.ServeJSON()
}
