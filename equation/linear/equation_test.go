package linear

import "testing"

func Test_EquationString(t *testing.T) {
	_, err := NewEquation(10, []float64{}, 2.3)
	if err != nil {
		t.Logf("need 10 unknown elements but find 0, error: %v", err)
	}

	e, err := NewEquation(10, []float64{-1.2, 1.4, 1.5, 3.4, 5.4, 6.77, 4.46, 67.23, -12.1, 12.0}, 2.3)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("equation %v", e)

	e1, err := NewEquation(10, []float64{1.2, 0, 1.5, 3.4, 5.4, 0, 4.46, 67.23, -12.1, 12.0}, 2.3)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("equation %v", e1)
}
