package interpolation

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
	// old method
	m := 1
	Lt := make([]float64, 1)
	dy := make([]float64, 100)
	ddy := make([]float64, 100)
	z := make([]float64, 1)
	dz := make([]float64, 1)
	ddz := make([]float64, 1)
	Lt[0] = x

	dy[0] = (pointsY[1] - pointsY[0]) / (pointsX[1] - pointsX[0])
	dy[len(pointsX)-1] = (pointsY[len(pointsX)-1] - pointsY[len(pointsX)-2]) / (pointsX[len(pointsX)-1] - pointsX[len(pointsX)-2])

	return auxiliary(pointsX, pointsY, len(pointsX), dy, ddy, Lt, m, z, dz, ddz, 1), nil
}

func auxiliary(x, y []float64, n int, dy, ddy, t []float64, m int, z, dz, ddz []float64, ii int) float64 {
	i := 0
	j := 0
	h0 := 0.0
	h1 := 0.0
	alpha := 0.0
	beta := 0.0
	g := 0.0
	s := make([]float64, 100)
	s[0] = dy[0]
	dy[0] = 0.0
	h0 = x[1] - x[0]

	for j = 1; j <= n-2; j++ {
		h1 = x[j+1] - x[j]
		alpha = h0 / (h0 + h1)
		beta = (1.0 - alpha) * (y[j] - y[j-1]) / h0
		beta = 3.0 * (beta + alpha*(y[j+1]-y[j])/h1)
		dy[j] = -alpha / (2.0 + (1.0-alpha)*dy[j-1])
		s[j] = (beta - (1.0-alpha)*s[j-1])
		s[j] = s[j] / (2.0 + (1.0-alpha)*dy[j-1])
		h0 = h1
	}
	for j = n - 2; j >= 0; j-- {
		dy[j] = dy[j]*dy[j+1] + s[j]

	}

	for j = 0; j <= n-2; j++ {
		s[j] = x[j+1] - x[j]
	}
	for j = 0; j <= n-2; j++ {
		h1 = s[j] * s[j]
		ddy[j] = 6.0*(y[j+1]-y[j])/h1 - 2.0*(2.0*dy[j]+dy[j+1])/s[j]
	}
	h1 = s[n-2] * s[n-2]
	ddy[n-1] = 6.0*(y[n-2]-y[n-1])/h1 + 2.0*(2.0*dy[n-1]+dy[n-2])/s[n-2]
	g = 0.0
	for i = 0; i <= n-2; i++ {
		h1 = 0.5 * s[i] * (y[i] + y[i+1])
		h1 = h1 - s[i]*s[i]*s[i]*(ddy[i]+ddy[i+1])/24.0
		g = g + h1
	}
	for j = 0; j <= m-1; j++ {
		if ii == 1 {
			if t[j] >= x[n-1] {
				i = n - 2
				return 0
			} else {
				if t[j] < x[0] {
					return 0
				} else {
					for i = 0; t[j] > x[i+1]; i++ {
					}
				}
			}
		} else {
			if t[j] >= x[n-1] {
				i = n - 2
				return 0
			} else {
				if t[j] < x[0] {
					return 0
				} else {
					for i = 0; t[j] > x[i+1]; i++ {
					}
				}
			}
		}
		h1 = (x[i+1] - t[j]) / s[i]
		h0 = h1 * h1
		z[j] = (3.0*h0 - 2.0*h0*h1) * y[i]
		z[j] = z[j] + s[i]*(h0-h0*h1)*dy[i]
		dz[j] = 6.0 * (h0 - h1) * y[i] / s[i]
		dz[j] = dz[j] + (3.0*h0-2.0*h1)*dy[i]
		ddz[j] = (6.0 - 12.0*h1) * y[i] / (s[i] * s[i])
		ddz[j] = ddz[j] + (2.0-6.0*h1)*dy[i]/s[i]
		h1 = (t[j] - x[i]) / s[i]
		h0 = h1 * h1
		z[j] = z[j] + (3.0*h0-2.0*h0*h1)*y[i+1]
		z[j] = z[j] - s[i]*(h0-h0*h1)*dy[i+1]
		dz[j] = dz[j] - 6.0*(h0-h1)*y[i+1]/s[i]
		dz[j] = dz[j] + (3.0*h0-2.0*h1)*dy[i+1]
		ddz[j] = ddz[j] + (6.0-12.0*h1)*y[i+1]/(s[i]*s[i])
		ddz[j] = ddz[j] - (2.0-6.0*h1)*dy[i+1]/s[i]
	}
	return z[0]
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
