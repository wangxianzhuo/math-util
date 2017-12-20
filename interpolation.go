package util

import (
	"fmt"
)

const (
	// CubicSplineInterpolation 三次样条插值
	CubicSplineInterpolation = iota
	// LagrangeInterpolation 拉格朗日插值
	LagrangeInterpolation
	// NewtonInterpolation 牛顿插值
	NewtonInterpolation
	// HermiteInterpolation Hermite插值
	HermiteInterpolation
)

const lastIntertpolation = HermiteInterpolation

// DefaultInterpolation 默认插值算法，默认使用三次样条插值
const DefaultInterpolation = CubicSplineInterpolation

// Interpolation 提供插值法求点
// x 是所求点的横坐标值
// pointsX, pointsY 是一组离散的点集
// algorithm 是算法选择可选0（CubicSplineInterpolation）、1（LagrangeInterpolation）、2（NewtonInterpolation）、3（HermiteInterpolation），默认使用CubicSplineInterpolation算法
func Interpolation(x float64, pointsX, pointsY []float64, algorithm int) (y float64, err error) {
	if ok := pointsCheck(pointsX, pointsY); !ok {
		return 0, fmt.Errorf("Interpolation error: point set{x: %v, y: %v} illegal", pointsX, pointsY)
	}

	if algorithm < 0 || algorithm > lastIntertpolation {
		algorithm = DefaultInterpolation
	}

	switch algorithm {
	case CubicSplineInterpolation:
		return CubicSpline(x, pointsX, pointsY)
	default:
		return CubicSpline(x, pointsX, pointsY)
	}
}

// CubicSpline 三次样条插值
func CubicSpline(x float64, pointsX, pointsY []float64) (y float64, err error) {
	return
}

// Lagrange 拉格朗日插值
func Lagrange(x float64, pointsX, pointsY []float64) (y float64, err error) {
	return
}

// Newton 牛顿插值
func Newton(x float64, pointsX, pointsY []float64) (y float64, err error) {
	return
}

// Hermite Hermite插值
func Hermite(x float64, pointsX, pointsY []float64) (y float64, err error) {
	return
}

func pointsCheck(pointsX, pointsY []float64) bool {
	if len(pointsX) != len(pointsY) {
		return false
	}
	return true
}
