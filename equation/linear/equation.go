package linear

import (
	"bytes"
	"fmt"
	"strings"
)

// Equation 一次方程模型
// UnknownNum 方程的未知数个数
// Factors 方程的系数 len(Factors) == UnknownNum
// Value 方程的值
type Equation struct {
	UnknownNum int
	Factors    []float64
	Value      float64
}

func (e Equation) String() string {
	buf := bytes.NewBufferString("")
	for index, f := range e.Factors {
		if f == -1 {
			buf.WriteString(fmt.Sprintf("- x%d ", index))
		} else if f == 1 {
			buf.WriteString(fmt.Sprintf("+ x%d ", index))
		} else if f < 0 {
			buf.WriteString(fmt.Sprintf("- %vx%d ", f*-1, index))
		} else if f != 0 {
			buf.WriteString(fmt.Sprintf("+ %vx%d ", f, index))
		}
	}
	buf.WriteString(fmt.Sprintf("= %v", e.Value))
	s := buf.String()
	if strings.HasPrefix(s, "-") {
		return "-" + strings.TrimSpace(s[1:])
	} else {
		return strings.TrimSpace(s[1:])
	}
}

// NewEquation get a linear equation
func NewEquation(unknownNum int, factors []float64, value float64) (Equation, error) {
	if len(factors) != unknownNum {
		return Equation{}, fmt.Errorf("factor(%v) can't match unknown number(%v)", len(factors), unknownNum)
	}

	return Equation{
		UnknownNum: unknownNum,
		Factors:    factors,
		Value:      value,
	}, nil
}

// Equations 方程组模型
// EquationList 方程列表
// UnknownNum 未知数个数
type Equations struct {
	EquationList []Equation
	UnknownNum   int
}

func (e Equations) String() string {
	buf := bytes.NewBufferString(fmt.Sprintf("unknown number %v\n", e.UnknownNum))
	for _, equation := range e.EquationList {
		buf.WriteString(equation.String())
		buf.WriteString("\n")
	}
	return buf.String()
}

// NewEquations new an equations
func NewEquations(unknownNum int, equationList []Equation) (Equations, error) {
	for _, equation := range equationList {
		if equation.UnknownNum != unknownNum {
			return Equations{}, fmt.Errorf("Create new equtaions error: equation [%v] unknown number [%v] can't match the unknown number %v", equation, equation.UnknownNum, unknownNum)
		}
	}

	return Equations{
		UnknownNum:   unknownNum,
		EquationList: equationList,
	}, nil
}

// EquationsValueVector 返回方程组值向量
func (e Equations) EquationsValueVector() []float64 {
	values := make([]float64, e.UnknownNum)
	for index, equation := range e.EquationList {
		values[index] = equation.Value
	}
	return values
}
