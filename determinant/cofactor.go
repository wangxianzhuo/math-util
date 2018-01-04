package determinant

// AlgebraicCofactor 计算代数余子式
func (d *Determinant) AlgebraicCofactor(row, cow int) *Determinant {
	sign := 1
	p := row + cow
	if p%2 != 0 {
		sign = -1
	}
	cofactor := d.Cofactor(row, cow)
	cofactor.MultiplyFactor(float64(sign))
	// fmt.Printf("algebraic cofactor: \n%v", cofactor)
	return cofactor
}

// Cofactor 计算余子式
func (d *Determinant) Cofactor(row, cow int) *Determinant {
	elements := make([][]float64, d.Rank-1)
	for index := 0; index < d.Rank-1; index++ {
		elements[index] = make([]float64, d.Rank-1)
	}

	for i := 0; i < d.Rank; i++ {
		for j := 0; j < d.Rank; j++ {
			e := d.Elements[i][j]
			r := i
			c := j
			if i == row {
				continue
			} else if i > row {
				r--
			}
			if j == cow {
				continue
			} else if j > cow {
				c--
			}
			elements[r][c] = e
		}
	}

	cofactor, _ := NewDeterminant(elements)
	return cofactor
}
