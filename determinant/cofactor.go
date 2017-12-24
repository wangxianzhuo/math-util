package determinant

func (d Determinant) AlgebraicCofactor(row, cow int) Determinant {
	sign := 1
	p := row + cow
	if p%2 != 0 {
		sign = -1
	}
	return d.Cofactor(row, cow).MultiplyFactor(float64(sign))
}

func (d Determinant) Cofactor(row, cow int) Determinant {
	elements := make([][]float64, d.RowLen-1)
	for index := 0; index < d.RowLen-1; index++ {
		elements[index] = make([]float64, d.CowLen-1)
	}

	for i := 0; i < d.RowLen; i++ {
		for j := 0; j < d.CowLen; j++ {
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
