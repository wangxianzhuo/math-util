package determinant

import "bytes"
import "fmt"

type Determinant struct {
	Elements [][]float64
	RowLen   int
	CowLen   int
	Factor   float64
	Value    float64
}

func NewDeterminant(elements [][]float64) (Determinant, error) {
	d := Determinant{
		Elements: elements,
		Factor:   1,
		RowLen:   len(elements),
		CowLen:   len(elements[0]),
	}

	return d, nil
}

func (d Determinant) String() string {
	buf := bytes.NewBufferString("")
	for row := 0; row < d.RowLen; row++ {
		buf.WriteString("|")
		for cow := 0; cow < d.CowLen; cow++ {
			buf.WriteString(fmt.Sprintf(" %2v ", d.Elements[row][cow]))
		}
		buf.WriteString("|\n")
	}
	return buf.String()
}

func (d Determinant) MultiplyFactor(factor float64) Determinant {
	d.Factor *= factor
	return d
}

func (d Determinant) CalculateValue() (float64, error) {
	return 0, nil
}

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

func (d Determinant) Expansion(row int) []Determinant {
	return nil
}
