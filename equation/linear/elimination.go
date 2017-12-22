package linear

import (
	"fmt"
)

// AdditionSubtractionElimination 加减消元法
func AdditionSubtractionElimination(es Equations) ([]float64, error) {
	if err := HasSolution(es); err != nil {
		return nil, fmt.Errorf("addition-subtraction elimination error: %v", err)
	}
	metrix, err := GetMatrixFromEquations(es)
	if err != nil {
		return nil, fmt.Errorf("addition-subtraction elimination error: %v", err)
	}
	// fmt.Println(metrix)
	currentRow := 0
	currentCow := 0
	for times := 0; times < len(es.EquationList); times++ {
		for row := currentRow; row < len(es.EquationList); row++ {
			sentry := metrix[row][currentCow]
			// fmt.Printf("sentry is %v\n", sentry)
			// fmt.Println(metrix[row])
			for cow := currentCow; cow < len(metrix[0]); cow++ {
				if row == currentRow {
					metrix[row][cow] /= sentry
				} else {
					metrix[row][cow] /= sentry
					metrix[row][cow] -= metrix[currentRow][cow]
				}
				// fmt.Println(metrix[row])
			}
			// fmt.Println()
		}
		// fmt.Println(metrix)
		currentRow++
		currentCow++
	}

	// fmt.Println(metrix)

	// 上三角阵求解
	// S[i] = M[i][n+1] - ∑(i≤j≤n-1)(M[i][j+1]*S[j+1]), 0 ≤ i ≤ n-1
	// S[i] = M[i][n+1], i = n
	solution := make([]float64, es.UnknownNum)
	for i := es.UnknownNum - 1; i >= 0; i-- {
		if i == es.UnknownNum-1 {
			solution[i] = metrix[i][es.UnknownNum-1]
		}
		solution[i] = metrix[i][es.UnknownNum]
		for j := i; j < es.UnknownNum-2; j++ {
			solution[i] -= metrix[i][j+1] * solution[j+1]
		}
	}

	return solution, nil
}

func HasSolution(es Equations) error {
	return nil
}

// GetMatrixFromEquations 方程组转换成矩阵
func GetMatrixFromEquations(es Equations) ([][]float64, error) {
	matrix := make([][]float64, len(es.EquationList))
	for index := 0; index < len(es.EquationList); index++ {
		matrix[index] = make([]float64, es.UnknownNum+1)
	}
	unknowNum := es.UnknownNum
	for row, e := range es.EquationList {
		if unknowNum != e.UnknownNum {
			return nil, fmt.Errorf("equation %v, do not have %v unknows", e, unknowNum)
		}
		for cow, f := range e.Factors {
			matrix[row][cow] = f
		}
		matrix[row][e.UnknownNum] = e.Value
	}
	return matrix, nil
}
