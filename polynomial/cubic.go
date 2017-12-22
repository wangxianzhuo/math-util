package polynomial

import (
	"fmt"
	"strings"
)

// CubicPolynomial 三次多项式
// f(x) = ax^3 + bx^2 + cx + d, min(∞) <(≤) x <(≤) max(∞)
type CubicPolynomial struct {
	A             float64
	B             float64
	C             float64
	D             float64
	Max           float64
	CanMax        bool
	IsMaxInfinite bool
	Min           float64
	CanMin        bool
	IsMinInfinite bool
}

func (p CubicPolynomial) String() string {
	var a, b, c, d string
	if p.A == 0 {
		a = ""
	} else {
		a = fmt.Sprintf("%vx^3", p.A)
	}
	if p.B == 0 {
		b = ""
	} else if p.B < 0 {
		b = fmt.Sprintf(" - %vx^2", p.B*-1)
	} else {
		b = fmt.Sprintf(" + %vx^2", p.B)
	}
	if p.C == 0 {
		c = ""
	} else if p.C < 0 {
		c = fmt.Sprintf(" - %vx", p.C*-1)
	} else {
		c = fmt.Sprintf(" + %vx", p.C)
	}
	if p.D == 0 {
		d = ""
	} else if p.D < 0 {
		d = fmt.Sprintf(" - %v", p.D*-1)
	} else {
		d = fmt.Sprintf(" + %v", p.D)
	}

	var condition string

	if p.IsMaxInfinite && p.IsMinInfinite {
		condition = "-∞ < x < +∞"
	} else if p.IsMaxInfinite {
		if p.CanMin {
			condition = fmt.Sprintf("x ≥ %v", p.Min)
		} else {
			condition = fmt.Sprintf("x > %v", p.Min)
		}
	} else if p.IsMinInfinite {
		if p.CanMax {
			condition = fmt.Sprintf("x ≤ %v", p.Max)
		} else {
			condition = fmt.Sprintf("x < %v", p.Max)
		}
	} else {
		var ls, rs string
		if p.CanMax {
			rs = "≤"
		} else {
			rs = "<"
		}

		if p.CanMin {
			ls = "≤"
		} else {
			ls = "<"
		}

		condition = fmt.Sprintf("%v %s x %s %v", p.Min, ls, rs, p.Max)
	}

	if a == "" && b == "" && c == "" && d == "" {
		return "F(x) = 0, " + condition
	}
	e := strings.TrimSpace(a + b + c + d)
	if strings.HasPrefix(e, "+") {
		e = e[1:]
	} else if strings.HasPrefix(e, "-") {
		e = "-" + strings.TrimSpace(e[1:])
	}
	return "F(x) = " + strings.TrimSpace(e) + ", " + condition

}

// NewCubicPolynomial get a CubicPolynomial
func NewCubicPolynomial(a, b, c, d, min, max float64, canMax, canMin, isMaxInfinite, isMinInfinite bool) (CubicPolynomial, error) {
	if !isMaxInfinite && !isMinInfinite {
		if canMax && canMin {
			if min > max {
				return CubicPolynomial{}, fmt.Errorf("the domain of definition[%v , %v] error", min, max)
			}
		} else {
			if min >= max {
				return CubicPolynomial{}, fmt.Errorf("the domain of definition(%v , %v) error", min, max)
			}
		}
	}

	return CubicPolynomial{
		A:             a,
		B:             b,
		C:             c,
		D:             d,
		Min:           min,
		Max:           max,
		CanMax:        canMax,
		CanMin:        canMin,
		IsMaxInfinite: isMaxInfinite,
		IsMinInfinite: isMinInfinite,
	}, nil
}

// DomainOfDefinition 获取定义域
func (p CubicPolynomial) DomainOfDefinition() string {
	var dod string

	if p.IsMaxInfinite && p.IsMinInfinite {
		dod = "-∞ < x < +∞"
	} else if p.IsMaxInfinite {
		if p.CanMin {
			dod = fmt.Sprintf("x ≥ %v", p.Min)
		} else {
			dod = fmt.Sprintf("x > %v", p.Min)
		}
	} else if p.IsMinInfinite {
		if p.CanMax {
			dod = fmt.Sprintf("x ≤ %v", p.Max)
		} else {
			dod = fmt.Sprintf("x < %v", p.Max)
		}
	} else {
		var ls, rs string
		if p.CanMax {
			rs = "≤"
		} else {
			rs = "<"
		}

		if p.CanMin {
			ls = "≤"
		} else {
			ls = "<"
		}

		dod = fmt.Sprintf("%v %s x %s %v", p.Min, ls, rs, p.Max)
	}
	return dod
}

// F 计算F(x)
func (p CubicPolynomial) F(x float64) (y float64, err error) {
	if err := p.IsInDomainOfDefinition(x); err != nil {
		return 0, fmt.Errorf("%v, F(%v) error: %v", p, x, err)
	}

	y = p.A*x*x*x + p.B*x*x + p.C*x + p.D
	return
}

// IsInDomainOfDefinition 判断x是否在定义域内
func (p CubicPolynomial) IsInDomainOfDefinition(x float64) error {
	if p.IsMaxInfinite && p.IsMinInfinite {
		return nil
	} else if p.IsMaxInfinite {
		if p.CanMin {
			if x >= p.Min {
				return nil
			} else {
				return fmt.Errorf("domain of definition( %v ) exceed, %v < %v", p.DomainOfDefinition(), x, p.Min)
			}
		} else {
			if x > p.Min {
				return nil
			} else {
				return fmt.Errorf("domain of definition( %v ) exceed, %v ≤ %v", p.DomainOfDefinition(), x, p.Min)
			}
		}
	} else if p.IsMinInfinite {
		if p.CanMax {
			if x <= p.Max {
				return nil
			} else {
				return fmt.Errorf("domain of definition( %v ) exceed, %v > %v", p.DomainOfDefinition(), x, p.Max)
			}

		} else {
			if x < p.Max {
				return nil
			} else {
				return fmt.Errorf("domain of definition( %v ) exceed, %v ≥ %v", p.DomainOfDefinition(), x, p.Max)
			}
		}
	} else {
		var ls, rs bool
		if p.CanMax {
			rs = (x <= p.Max)
		} else {
			rs = (x < p.Max)
		}

		if p.CanMin {
			ls = (x >= p.Min)
		} else {
			ls = (x > p.Min)
		}

		if rs && ls {
			return nil
		} else {
			return fmt.Errorf("%v is not in domain of definition %v", x, p.DomainOfDefinition())
		}
	}
}
