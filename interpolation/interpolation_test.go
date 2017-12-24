package interpolation

import (
	"testing"

	"github.com/wangxianzhuo/math-util/polynomial"
)

func Test_Interpolation(t *testing.T) {
	pointX := []float64{1, 2, 3, 4, 5}
	pointY := []float64{1, 2, 3, 4, 5}

	x := 9.4

	y, err := Interpolation(x, pointX, pointY, DefaultInterpolation)
	if err != nil {
		t.Errorf("interpolation calc error: %v", err)
	}
	t.Logf("x = %f, y = %f", x, y)
}

func Test_CubicInterpolation(t *testing.T) {
	p, err := polynomial.NewCubicPolynomial(4, 3.5, 0, 9, 0, 0, false, false, true, true)
	if err != nil {
		t.Error(err)
	}

	pointsX := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	pointsY := make([]float64, 9)

	for index := 0; index < 9; index++ {
		pointsY[index], err = p.F(pointsX[index])
		if err != nil {
			t.Error(err)
		}
	}

	x := 2.28
	y, err := p.F(x)
	if err != nil {
		t.Error(err)
	}

	y1, err := CubicSpline(x, pointsX, pointsY)
	if err != nil {
		t.Error(err)
	}

	t.Logf("original (%v, %v), interpolation (%v, %v)\n", x, y, x, y1)
}

func Benchmark_CubicInterpolation(b *testing.B) {
	p, err := polynomial.NewCubicPolynomial(4, 3.5, 0, 9, 0, 0, false, false, true, true)
	if err != nil {
		b.Error(err)
	}

	pointsX := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	pointsY := make([]float64, 9)

	for index := 0; index < 9; index++ {
		pointsY[index], err = p.F(pointsX[index])
		if err != nil {
			b.Error(err)
		}
	}

	for i := 0; i < b.N; i++ {
		x := 2.28
		CubicSpline(x, pointsX, pointsY)
	}
}
