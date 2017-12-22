package linear

import "testing"

func Test_AdditionSubtractionElimination(t *testing.T) {
	e1, err := NewEquation(3, []float64{1, 2, 1}, 1)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	e2, err := NewEquation(3, []float64{2, 3, 3}, 2)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	e3, err := NewEquation(3, []float64{3, 3, 5}, 3)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	es := Equations{
		EquationList: []Equation{e1, e2, e3},
		UnknownNum:   3,
	}

	solution, err := AdditionSubtractionElimination(es)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	t.Logf("solutions are: %v", solution)
}
