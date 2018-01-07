package determinant

import (
	"bytes"
	"errors"
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

// NewDeterminantWithValueParallel new a Determinant with Value parallel
func NewDeterminantWithValueParallel(elements [][]float64) (*Determinant, error) {
	det, err := NewDeterminant(elements)
	if err != nil {
		return nil, fmt.Errorf("Create new determinant with value error: %v", err)
	}
	// det.Value, err = det.CalculateValueParalle()
	// if err != nil {
	// 	return nil, fmt.Errorf("Create new determinant with value calculate value error: %v", err)
	// }

	valueChan := make(chan CalcValue)
	defer close(valueChan)
	go det.CalculateValueParalle(valueChan, det.Rank)
	select {
	case value := <-valueChan:
		if v, ok := value.value.(float64); !ok {
			return nil, fmt.Errorf("Create new determinant with value error: get error value")
		} else {
			det.Value = v
		}
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
	d.Mu.RLock()
	defer d.Mu.RUnlock()
	return d.Elements[i][j], nil
}

// SetElement 修改行列式元素，线程安全
func (d *Determinant) SetElement(value float64, i, j int) error {
	if i > d.Rank-1 {
		return fmt.Errorf("row %v exceed bound %v", i, d.Rank)
	}
	if j > d.Rank-1 {
		return fmt.Errorf("cow %v exceed bound %v", i, d.Rank)
	}
	if i < 0 || j < 0 {
		return fmt.Errorf("row %v or cow %v lower than 0", i, j)
	}
	d.Mu.Lock()
	defer d.Mu.Unlock()
	d.Elements[i][j] = value
	return nil
}

// MultiplyFactor 行列式乘系数
func (d *Determinant) MultiplyFactor(factor float64) {
	d.Mu.Lock()
	defer d.Mu.Unlock()
	d.Factor *= factor
}

// GetValue 返回行列式的值
func (d *Determinant) GetValue() (float64, error) {
	var value float64
	valueChan := make(chan CalcValue)
	defer close(valueChan)
	go d.CalculateValueParalle(valueChan, d.Rank)
	select {
	case result := <-valueChan:
		if v, ok := result.value.(float64); !ok {
			return 0.0, fmt.Errorf("Get determinant's value error: get error value")
		} else {
			value = v
		}
	}
	return value, nil
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

// CalcValue 带error的数据
type CalcValue struct {
	value interface{}
	err   error
}

// CalculateValueParalle 计算行列式的值（并行版）
func (d *Determinant) CalculateValueParalle(valueChan chan CalcValue, rootRank int) {
	var value float64

	if d.Rank == 1 {
		value, err := d.GetElement(0, 0)
		if err != nil {
			valueChan <- CalcValue{
				err: fmt.Errorf("Calculate determinant value error: %v", err),
			}
			return
		}
		valueChan <- CalcValue{
			value: value * d.Factor,
		}
		return
	}

	ds, err := d.Expanse(0, UseRow)
	if err != nil {
		valueChan <- CalcValue{
			err: fmt.Errorf("Calculate determinant value error: %v", err),
		}
	}
	// fmt.Printf("determinants: %v", ds)

	if d.Rank == rootRank {
		subValueChan := make(chan CalcValue)
		defer close(subValueChan)

		for _, det := range ds {
			go det.CalculateValueParalle(subValueChan, rootRank)
		}

		counter := d.Rank
		var errList []error
		for subValue := range subValueChan {
			if subValue.err != nil {
				errList = append(errList, subValue.err)
			}
			if v, ok := subValue.value.(float64); !ok {
				errList = append(errList, fmt.Errorf("unknown value of cofactor"))
			} else {
				value += v
			}
			counter--
			if counter == 0 {
				break
			}
		}
	} else {
		ds, err := d.Expanse(0, UseRow)
		if err != nil {
			valueChan <- CalcValue{
				err: fmt.Errorf("Calculate determinant value error: %v", err),
			}
		}
		// fmt.Printf("determinants: %v", ds)
		var errList []error
		for _, det := range ds {
			v, err := det.CalculateValue()
			if err != nil {
				errList = append(errList, err)
			}
			value += v
		}
		if len(errList) != 0 {
			valueChan <- CalcValue{
				err: fmt.Errorf("Calculate determinant value error: %v", errList),
			}
			return
		}
	}
	valueChan <- CalcValue{
		value: value * d.Factor,
	}
}

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

// ErrEquationNoSolves 方程组无解
var ErrEquationNoSolves = errors.New("equations have no results")

// KramerRuleSolveEquation 克拉默法则解方程组
func KramerRuleSolveEquation(equations linear.Equations) ([]float64, error) {
	valueList := equations.EquationsValueVector()
	det, err := equationsToDeterminant(equations)
	if err != nil {
		return nil, fmt.Errorf("sovle equations error: %v", err)
	}
	D, err := det.GetValue()
	if err != nil {
		return nil, fmt.Errorf("sovle equations error: %v", err)
	}
	if D == 0 {
		return nil, ErrEquationNoSolves
	}
	Dn := make([]float64, equations.UnknownNum)

	for cow := 0; cow < det.Rank; cow++ {
		newDet, _ := replaceRowOrCow(det, valueList, cow, UseCow)
		Dn[cow], err = newDet.GetValue()
		if err != nil {
			return nil, fmt.Errorf("sovle equations error: %v", err)
		}
	}

	resultList := make([]float64, equations.UnknownNum)
	for index := 0; index < equations.UnknownNum; index++ {
		resultList[index] = Dn[index] / D
	}

	return resultList, nil
}

// KramerRuleSolveEquationParallel 克拉默法则解方程组，并行算法
// TODO: 分析并行运算效率
func KramerRuleSolveEquationParallel(equations linear.Equations) ([]float64, error) {
	valueList := equations.EquationsValueVector()
	det, err := equationsToDeterminant(equations)
	if err != nil {
		return nil, fmt.Errorf("sovle equations error: %v", err)
	}
	D, err := det.GetValue()
	if err != nil {
		return nil, fmt.Errorf("sovle equations error: %v", err)
	}
	if D == 0 {
		return nil, ErrEquationNoSolves
	}

	resultList := make([]float64, equations.UnknownNum)
	dataChan := make(chan replacedDetValue)
	defer close(dataChan)

	for index := 0; index < equations.UnknownNum; index++ {
		go replacedDetParallel(det, valueList, index, dataChan, index)
	}

	counter := equations.UnknownNum
	var errList []error
	for data := range dataChan {
		if data.err != nil {
			errList = append(errList, data.err)
		} else {
			resultList[data.index] = data.value / D
		}
		counter--
		if counter == 0 {
			break
		}
	}
	if len(errList) != 0 {
		return nil, fmt.Errorf("sovle equations error: %v", errList)
	}

	return resultList, nil
}

type replacedDetValue struct {
	index int
	value float64
	err   error
}

func replacedDetParallel(det *Determinant, valueList []float64, cow int, dataChan chan replacedDetValue, index int) {
	newDet, _ := replaceRowOrCow(det, valueList, cow, UseCow)
	value, err := newDet.GetValue()
	if err != nil {
		dataChan <- replacedDetValue{
			err: fmt.Errorf("Replace determinant value error: %v", err),
		}
		return
	}
	dataChan <- replacedDetValue{
		index: index,
		value: value,
	}
}

func replaceRowOrCow(det *Determinant, list []float64, order int, eType RowCowType) (*Determinant, error) {
	if order > det.Rank {
		c := ""
		switch eType {
		case UseCow:
			c = "cow"
		default:
			c = "row"
		}
		return nil, fmt.Errorf("replace %s error: %v exceed det's rank %v", c, order, det.Rank)
	}
	if len(list) != det.Rank {
		return nil, fmt.Errorf("replace determinant row or cow error: the give list %v can't match the det's rank %v", list, det.Rank)
	}

	elements := make([][]float64, det.Rank)
	switch eType {
	case UseCow:
		for row := 0; row < det.Rank; row++ {
			elements[row] = make([]float64, det.Rank)
			for cow := 0; cow < det.Rank; cow++ {
				if cow == order {
					elements[row][cow] = list[row]
				} else {
					elements[row][cow], _ = det.GetElement(row, cow)
				}
			}
		}
	default:
		for row := 0; row < det.Rank; row++ {
			elements[row] = make([]float64, det.Rank)
			for cow := 0; cow < det.Rank; cow++ {
				if row == order {
					elements[row][cow] = list[row]
				} else {
					elements[row][cow], _ = det.GetElement(row, cow)
				}
			}
		}
	}
	return NewDeterminant(elements)
}

// equationsToDeterminant 返回一个未结算值的行列式
func equationsToDeterminant(equations linear.Equations) (*Determinant, error) {
	unknownNumber := equations.UnknownNum
	elements := make([][]float64, unknownNumber)
	for i, equation := range equations.EquationList {
		if unknownNumber != equation.UnknownNum {
			return nil, fmt.Errorf("equation [%v]'s unknown number can't match equations'(%v)", equation, unknownNumber)
		}
		line := make([]float64, equation.UnknownNum)
		for j, value := range equation.Factors {
			line[j] = value
		}
		elements[i] = line
	}
	return NewDeterminant(elements)
}
