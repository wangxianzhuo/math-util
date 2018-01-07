package determinant

import "testing"
import "fmt"
import "github.com/wangxianzhuo/math-util/equation/linear"

func Test_CalculateInversionNumber(t *testing.T) {
	n1 := []int{1, 2, 3, 4, 5}
	n2 := []int{4, 3, 5, 6, 23, 7, 2, 54, 9}

	if r1 := CalculateInversionNumber(n1); r1 != 0 {
		t.Errorf("calculate %v result %v error", n1, r1)
	}

	if r2 := CalculateInversionNumber(n2); r2 != 10 {
		t.Errorf("calculate %v result %v error", n2, r2)
	}
}

func Benchmark_CalculateInversionNumber(b *testing.B) {
	n2 := []int{4, 3, 5, 6, 23, 7, 2, 54, 9}
	for i := 0; i < b.N; i++ {
		CalculateInversionNumber(n2)
	}
}

func Test_Cofactor(t *testing.T) {
	e := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	det, _ := NewDeterminant(e)
	fmt.Printf("%v\n", det)

	fmt.Printf("%v\n", det.Cofactor(0, 0))
	fmt.Printf("%v\n", det.Cofactor(1, 0))
	fmt.Printf("%v\n", det.Cofactor(0, 2))
}

func Benchmark_Cofactor(b *testing.B) {
	e := [][]float64{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	det, _ := NewDeterminant(e)
	for i := 0; i < b.N; i++ {
		det.Cofactor(0, 0)
	}
}

func Test_CalculateValue(t *testing.T) {
	// e := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	e := [][]float64{
		{3, 1, -1, 2},
		{-5, 1, 3, -4},
		{2, 0, 1, -1},
		{1, -5, 3, -3},
	}
	det, err := NewDeterminantWithValue(e)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	fmt.Printf("%v\nvalue = %v\n", det, det.Value)
}

func Test_CalculateValueParallel(t *testing.T) {
	e := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	det, err := NewDeterminantWithValueParallel(e)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	fmt.Printf("%v\nvalue = %v\n", det, det.Value)
}

func Test_Expanse(t *testing.T) {
	e := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	det, err := NewDeterminant(e)
	ds, err := det.Expanse(0, UseRow)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	fmt.Printf("ds %v\n", ds)
}

func Benchmark_Expanse(b *testing.B) {
	// e := [][]float64{
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// 	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	// }
	// e := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// e := [][]float64{{1, 2}, {4, 5}}
	// det, _ := NewDeterminant(e)
	det := generateDeteminantByRank(10)
	// fmt.Println(det)
	for i := 0; i < b.N; i++ {
		_, err := det.Expanse(0, UseRow)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Test_replaceRowOrCow(t *testing.T) {
	det := generateDeteminantByRank(3)
	list1 := []float64{1.5, 1.5, 1.5}
	list2 := []float64{2.5, 2.5}

	t.Logf("## before replace, the det is : %v\n", det)
	newDet, err := replaceRowOrCow(det, list1, 0, UseCow)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Logf("## after replace cow 0, the det is : %v\n", newDet)
	_, err = replaceRowOrCow(det, list1, 4, UseCow)
	if err != nil {
		t.Logf("expect error: %v\n", err)
	}
	_, err = replaceRowOrCow(det, list2, 0, UseCow)
	if err != nil {
		t.Logf("expect error: %v\n", err)
	}
}

func Test_equationsToDeterminant(t *testing.T) {
	e1, _ := linear.NewEquation(4, []float64{2, 1, -5, 1}, 8)
	e2, _ := linear.NewEquation(4, []float64{1, -3, 0, -6}, 9)
	e3, _ := linear.NewEquation(4, []float64{0, 2, -1, 2}, -5)
	e4, _ := linear.NewEquation(4, []float64{1, 4, -7, 6}, 0)

	es, err := linear.NewEquations(4, []linear.Equation{e1, e2, e3, e4})
	if err != nil {
		t.Errorf("%v\n", err)
	}
	det, err := equationsToDeterminant(es)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	fmt.Printf("equations: %v\n", es)
	fmt.Printf("det: %v\n", det)
}

func Test_KramerRuleSolveEquation(t *testing.T) {
	e1, _ := linear.NewEquation(4, []float64{2, 1, -5, 1}, 8)
	e2, _ := linear.NewEquation(4, []float64{1, -3, 0, -6}, 9)
	e3, _ := linear.NewEquation(4, []float64{0, 2, -1, 2}, -5)
	e4, _ := linear.NewEquation(4, []float64{1, 4, -7, 6}, 0)

	es, err := linear.NewEquations(4, []linear.Equation{e1, e2, e3, e4})
	if err != nil {
		t.Errorf("%v\n", err)
	}

	results, err := KramerRuleSolveEquation(es)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	fmt.Printf("equations : %v\n solve : %v\n", es, results)
}
