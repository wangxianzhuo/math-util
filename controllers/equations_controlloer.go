package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/wangxianzhuo/math-util/determinant"
	"github.com/wangxianzhuo/math-util/equation/linear"
)

// EquationsController 方程组相关接口控制器
type EquationsController struct {
	beego.Controller
}

// EquationsSolveMode 解方程组方法
type EquationsSolveMode int

// UseKramerRule 解方程组使用克拉默法则
var UseKramerRule EquationsSolveMode

// Solve 解方程组，默认solveMode是UseKramerRule
//
// URL:		/equations/solve
// Method:	GET
// Param:	factors=[[2,1,-5,1],[1,-3,0,-6],[0,2,-1,2],[1,4,-7,6]]&values=[8,9,-5,0]&unknownNum=4
// Return:	{"value":[3,-4,-1,1],"des":"Onlyonesolution"}
func (ec *EquationsController) Solve() {
	solveMode, _ := ec.GetInt("solveMode")
	var factors [][]float64
	err := json.Unmarshal([]byte(ec.GetString("factors")), &factors)
	if err != nil {
		beego.Error("Solve equations error ", err)
		ec.Abort("500")
	}
	var values []float64
	err = json.Unmarshal([]byte(ec.GetString("values")), &values)
	if err != nil {
		beego.Error("Solve equations error ", err)
		ec.Abort("500")
	}
	unknownNum, err := ec.GetInt("unknownNum")
	if err != nil {
		beego.Error("Solve equations error ", err)
		ec.Abort("500")
	}

	equations, err := linear.NewEquationsWithElements(factors, values, unknownNum)
	if err != nil {
		beego.Error("Solve equations error ", err)
		ec.Abort("500")
	}

	var result []float64
	switch solveMode {
	default:
		result, err = determinant.KramerRuleSolveEquationParallel(equations)
		if err != nil && err == determinant.ErrEquationNoSolves {
			ec.Data["json"] = CalcResult{
				Value:       nil,
				Description: "No solution",
			}
			ec.ServeJSON()
			return
		} else if err != nil {
			beego.Error("Solve equations error ", err)
			ec.Abort("500")
		}
	}

	ec.Data["json"] = CalcResult{
		Value:       result,
		Description: "Only one solution",
	}
	ec.ServeJSON()
}
