package polynomial

import (
	"fmt"
	"testing"
)

func Test_CubicString(t *testing.T) {
	p2 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        false,
		IsMaxInfinite: true,
		IsMinInfinite: true,
	}
	fmt.Printf("%v\n", p2)

	p3 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        true,
		CanMin:        true,
		IsMaxInfinite: true,
		IsMinInfinite: true,
	}
	fmt.Printf("%v\n", p3)

	p1 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        false,
		IsMaxInfinite: false,
		IsMinInfinite: false,
	}
	fmt.Printf("%v\n", p1)

	p4 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        true,
		CanMin:        true,
		IsMaxInfinite: false,
		IsMinInfinite: false,
	}
	fmt.Printf("%v\n", p4)

	p5 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        true,
		CanMin:        false,
		IsMaxInfinite: false,
		IsMinInfinite: false,
	}
	fmt.Printf("%v\n", p5)

	p6 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        true,
		IsMaxInfinite: false,
		IsMinInfinite: false,
	}
	fmt.Printf("%v\n", p6)

	p7 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        true,
		IsMaxInfinite: true,
		IsMinInfinite: false,
	}
	fmt.Printf("%v\n", p7)

	p8 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        false,
		IsMaxInfinite: true,
		IsMinInfinite: false,
	}
	fmt.Printf("%v\n", p8)

	p9 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        true,
		CanMin:        true,
		IsMaxInfinite: false,
		IsMinInfinite: true,
	}
	fmt.Printf("%v\n", p9)

	p10 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        false,
		IsMaxInfinite: false,
		IsMinInfinite: true,
	}
	fmt.Printf("%v\n", p10)

	p11 := CubicPolynomial{
		A:             0,
		B:             0,
		C:             0,
		D:             0,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        false,
		IsMaxInfinite: false,
		IsMinInfinite: true,
	}
	fmt.Printf("%v\n", p11)
}

func Test_CheckDOD(t *testing.T) {
	var err error
	x := 0.0
	p2 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        false,
		IsMaxInfinite: true,
		IsMinInfinite: true,
	}
	err = p2.IsInDomainOfDefinition(x)
	if err != nil {
		fmt.Printf("p2  %v, error :%v\n", p2, err)
	}

	p3 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        true,
		CanMin:        true,
		IsMaxInfinite: true,
		IsMinInfinite: true,
	}
	err = p3.IsInDomainOfDefinition(x)
	if err != nil {
		fmt.Printf("p3  %v, error :%v\n", p3, err)
	}

	p1 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        false,
		IsMaxInfinite: false,
		IsMinInfinite: false,
	}
	err = p1.IsInDomainOfDefinition(x)
	if err != nil {
		fmt.Printf("p1  %v, error :%v\n", p1, err)
	}

	p4 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        true,
		CanMin:        true,
		IsMaxInfinite: false,
		IsMinInfinite: false,
	}
	err = p4.IsInDomainOfDefinition(x)
	if err != nil {
		fmt.Printf("p4  %v, error :%v\n", p4, err)
	}

	p5 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        true,
		CanMin:        false,
		IsMaxInfinite: false,
		IsMinInfinite: false,
	}
	err = p5.IsInDomainOfDefinition(x)
	if err != nil {
		fmt.Printf("p5  %v, error :%v\n", p5, err)
	}

	p6 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        true,
		IsMaxInfinite: false,
		IsMinInfinite: false,
	}
	err = p6.IsInDomainOfDefinition(x)
	if err != nil {
		fmt.Printf("p6  %v, error :%v\n", p6, err)
	}

	p7 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        true,
		IsMaxInfinite: true,
		IsMinInfinite: false,
	}
	err = p7.IsInDomainOfDefinition(x)
	if err != nil {
		fmt.Printf("p7  %v, error :%v\n", p7, err)
	}

	p8 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        false,
		IsMaxInfinite: true,
		IsMinInfinite: false,
	}
	err = p8.IsInDomainOfDefinition(x)
	if err != nil {
		fmt.Printf("p8  %v, error :%v\n", p8, err)
	}

	p9 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        true,
		CanMin:        true,
		IsMaxInfinite: false,
		IsMinInfinite: true,
	}
	err = p9.IsInDomainOfDefinition(x)
	if err != nil {
		fmt.Printf("p9  %v, error :%v\n", p9, err)
	}

	p10 := CubicPolynomial{
		A:             2,
		B:             3,
		C:             4,
		D:             5,
		Max:           1,
		Min:           0,
		CanMax:        false,
		CanMin:        false,
		IsMaxInfinite: false,
		IsMinInfinite: true,
	}
	err = p10.IsInDomainOfDefinition(x)
	if err != nil {
		fmt.Printf("p10 %v, error :%v\n", p10, err)
	}
}

func Test_F(t *testing.T) {
	_, err := NewCubicPolynomial(3.0, 1.5, 6.0, 9.0, 0, 0, false, false, false, false)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	_, err = NewCubicPolynomial(3.0, 1.5, 6.0, 9.0, 2, 1, true, true, false, false)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	p, err := NewCubicPolynomial(3.0, 1.5, 6.0, 9.0, 0, 0, false, false, true, true)
	if err != nil {
		t.Error(err)
	}

	x := 7.0
	y, err := p.F(x)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v, F(%v) = %v\n", p, x, y)

	p1, err := NewCubicPolynomial(3.0, 1.5, 6.0, 9.0, 0, 2, false, false, false, false)
	if err != nil {
		t.Error(err)
	}

	x1 := 7.0
	_, err = p1.F(x1)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
