package interpolation

import "testing"

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
