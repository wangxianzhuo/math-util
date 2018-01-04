package determinant

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/wangxianzhuo/math-util/equation/linear"
)

// RowCowType 行列式行、列使用类型
type RowCowType int

// UseRow 按行使用
// UseCow 按列使用
var (
	UseRow RowCowType
	UseCow RowCowType = 1
)

// Determinant 行列式，Factor 线程安全
type Determinant struct {
	Elements [][]float64
	Rank     int
	Factor   float64
	Value    float64
	Mu       sync.RWMutex
}

// NewDeterminant new a Determinant
func NewDeterminant(elements [][]float64) (*Determinant, error) {
	// var err error
	row := len(elements)
	cow := len(elements[0])
	if row != cow || row == 0 || cow == 0 {
		return nil, fmt.Errorf("Create new determinant error: the elements(%v) with row(%v) cow(%v) are illegal", elements, row, cow)
	}
	d := Determinant{
		Elements: elements,
		Factor:   1,
		Rank:     len(elements),
	}

	// d.Value, err = d.CalculateValue()
	// if err != nil {
	// 	return nil, fmt.Errorf("Create new determinant error: %v", err)
	// }
	return &d, nil
}

// NewDeterminantWithValue new a Determinant with Value
func NewDeterminantWithValue(elements [][]float64) (*Determinant, error) {
	det, err := NewDeterminant(elements)
	if err != nil {
		return nil, fmt.Errorf("Create new determinant with value error: %v", err)
	}
	det.Value, err = det.CalculateValue()
	if err != nil {
		return nil, fmt.Errorf("Create new determinant with value calculate value error: %v", err)
	}
	return det, nil
}

func (d *Determinant) String() string {
	buf := bytes.NewBufferString("factor ")
	buf.WriteString(fmt.Sprintf("%v\n", d.Factor))
	for row := 0; row < d.Rank; row++ {
		buf.WriteString("|")
		for cow := 0; cow < d.Rank; cow++ {
			buf.WriteString(fmt.Sprintf(" %2v ", d.Elements[row][cow]))
		}
		buf.WriteString("|\n")
	}
	return buf.String()
}

// GetElement 获取单个元素，线程安全
func (d *Determinant) GetElement(i, j int) (float64, error) {
	if i > d.Rank-1 {
		return 0, fmt.Errorf("row %v exceed bound %v", i, d.Rank)
	}
	if j > d.Rank-1 {
		return 0, fmt.Errorf("cow %v exceed bound %v", i, d.Rank)
	}
	if i < 0 || j < 0 {
		return 0, fmt.Errorf("row %v or cow %v lower than 0", i, j)
	}

	return d.Elements[i][j], nil
}

// MultiplyFactor 行列式乘系数
func (d *Determinant) MultiplyFactor(factor float64) {
	d.Mu.Lock()
	defer d.Mu.Unlock()
	d.Factor *= factor
}

// CalculateValue 计算行列式的值
func (d *Determinant) CalculateValue() (float64, error) {
	// d.Mu.Lock()
	// defer d.Mu.RUnlock()
	var value float64
	// fmt.Printf("det: \n%v\n", d)

	if d.Rank == 1 {
		value, err := d.GetElement(0, 0)
		if err != nil {
			return 0, fmt.Errorf("Calculate determinant value error: %v", err)
		}
		return value * d.Factor, nil
	}

	ds, err := d.Expanse(0, UseRow)
	if err != nil {
		return 0, fmt.Errorf("Calculate determinant value error: %v", err)
	}
	// fmt.Printf("determinants: %v", ds)
	for _, det := range ds {
		v, err := det.CalculateValue()
		if err != nil {
			return 0, fmt.Errorf("Calculate determinant value error: %v", err)
		}
		value += v
	}
	return value * d.Factor, nil
}

// TODO:计算行列式的值（并行版）

// Expanse 行列式按行（列）展开，默认按行展开
func (d *Determinant) Expanse(order int, eType RowCowType) ([]*Determinant, error) {
	determinants := make([]*Determinant, d.Rank)
	switch eType {
	case UseCow:
		for row := 0; row < d.Rank; row++ {
			e, err := d.GetElement(row, order)
			if err != nil {
				return nil, fmt.Errorf("determinant Expanse row %v cow %v error：%v", row, order, err)
			}
			ac := d.AlgebraicCofactor(row, order)
			ac.MultiplyFactor(e)
			determinants[row] = ac
		}
	default:
		for cow := 0; cow < d.Rank; cow++ {
			e, err := d.GetElement(order, cow)
			if err != nil {
				return nil, fmt.Errorf("determinant Expanse row %v cow %v error：%v", order, cow, err)
			}
			ac := d.AlgebraicCofactor(order, cow)
			ac.MultiplyFactor(e)
			determinants[cow] = ac
		}
	}
	return determinants, nil
}

// CalculateInversionNumber 计算逆序数
func CalculateInversionNumber(numbers []int) int {
	result := 0
	for i := 1; i < len(numbers); i++ {
		for j := 0; j < i; j++ {
			if numbers[j] > numbers[i] {
				result++
			}
		}
	}
	return result
}

// KramerRuleSolveEquation 克拉默法则解方程组
// TODO:
func KramerRuleSolveEquation(equations linear.Equations) ([]float64, error) {
	return nil, nil
}
