package determinant

import (
	"testing"

	"github.com/wangxianzhuo/math-util/equation/linear"
)

func Benchmark_CalculateValue_1(b *testing.B) {
	det := generateDeteminantByRank(1)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_1(b *testing.B) {
	det := generateDeteminantByRank(1)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}
func Benchmark_CalculateValue_2(b *testing.B) {
	det := generateDeteminantByRank(2)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_2(b *testing.B) {
	det := generateDeteminantByRank(2)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}
func Benchmark_CalculateValue_3(b *testing.B) {
	det := generateDeteminantByRank(3)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_3(b *testing.B) {
	det := generateDeteminantByRank(3)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}
func Benchmark_CalculateValue_4(b *testing.B) {
	det := generateDeteminantByRank(4)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_4(b *testing.B) {
	det := generateDeteminantByRank(4)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}
func Benchmark_CalculateValue_5(b *testing.B) {
	det := generateDeteminantByRank(5)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_5(b *testing.B) {
	det := generateDeteminantByRank(5)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}
func Benchmark_CalculateValue_6(b *testing.B) {
	det := generateDeteminantByRank(6)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_6(b *testing.B) {
	det := generateDeteminantByRank(6)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}
func Benchmark_CalculateValue_7(b *testing.B) {
	det := generateDeteminantByRank(7)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_7(b *testing.B) {
	det := generateDeteminantByRank(7)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}
func Benchmark_CalculateValue_8(b *testing.B) {
	det := generateDeteminantByRank(8)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_8(b *testing.B) {
	det := generateDeteminantByRank(8)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}
func Benchmark_CalculateValue_9(b *testing.B) {
	det := generateDeteminantByRank(9)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_9(b *testing.B) {
	det := generateDeteminantByRank(9)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValue_10(b *testing.B) {
	det := generateDeteminantByRank(10)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_10(b *testing.B) {
	det := generateDeteminantByRank(10)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValue_11(b *testing.B) {
	det := generateDeteminantByRank(11)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValue(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_CalculateValueParallel_11(b *testing.B) {
	det := generateDeteminantByRank(11)
	for i := 0; i < b.N; i++ {
		_, err := NewDeterminantWithValueParallel(det.Elements)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}
func generateDeteminantByRank(rank int) *Determinant {
	num := 0.0
	element := make([][]float64, rank)
	for index := 0; index < rank; index++ {
		element[index] = make([]float64, rank)
		for order := 0; order < rank; order++ {
			element[index][order] = num
			num++
		}
	}

	d, _ := NewDeterminant(element)
	return d
}

func Benchmark_KramerRuleSolveEquation(b *testing.B) {
	e1, _ := linear.NewEquation(4, []float64{2, 1, -5, 1}, 8)
	e2, _ := linear.NewEquation(4, []float64{1, -3, 0, -6}, 9)
	e3, _ := linear.NewEquation(4, []float64{0, 2, -1, 2}, -5)
	e4, _ := linear.NewEquation(4, []float64{1, 4, -7, 6}, 0)

	es, err := linear.NewEquations(4, []linear.Equation{e1, e2, e3, e4})
	if err != nil {
		b.Errorf("%v\n", err)
	}
	for i := 0; i < b.N; i++ {
		_, err := KramerRuleSolveEquation(es)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_KramerRuleSolveEquationParallel(b *testing.B) {
	e1, _ := linear.NewEquation(4, []float64{2, 1, -5, 1}, 8)
	e2, _ := linear.NewEquation(4, []float64{1, -3, 0, -6}, 9)
	e3, _ := linear.NewEquation(4, []float64{0, 2, -1, 2}, -5)
	e4, _ := linear.NewEquation(4, []float64{1, 4, -7, 6}, 0)

	es, err := linear.NewEquations(4, []linear.Equation{e1, e2, e3, e4})
	if err != nil {
		b.Errorf("%v\n", err)
	}
	for i := 0; i < b.N; i++ {
		_, err := KramerRuleSolveEquationParallel(es)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}

func Benchmark_replaceRowOrCow(b *testing.B) {
	det := generateDeteminantByRank(4)
	list1 := []float64{1.5, 1.5, 1.5, 1.5}
	for i := 0; i < b.N; i++ {
		_, err := replaceRowOrCow(det, list1, 0, UseCow)
		if err != nil {
			b.Errorf("%v\n", err)
		}
	}
}
