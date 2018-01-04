# Math Util

## Description

1. 提供插值法求点（当前支持三次样条插值）

## Plan

- 行列式（`package` determinant）
  - 余子式
  - 代数余子式
  - 行列式展开
  - 行列式求值
  - 逆序数
  - 行列式乘系数
  - 取行列式元素
  - 克拉默法则解方程组（TODO）
- 插值法（`package` interpolation）
  - 三次样条插值（未完善）
  - 拉格朗日插值（TODO）
  - 牛顿插值（TODO）
  - Hermite插值（TODO）
- 三次函数（`package` polynomial)
  - 函数求值
  - 函数定义域判断
- 方程（`package` equation）

## Usage

```shell
go get -u github.com/wangxianzhuo/math-util
```

```go
package main

import interpolation "github.com/wangxianzhuo/math-util/interpolation"
import log

func main() {
    pointX := []float64{1, 2, 3, 4, 5}
    pointY := []float64{1, 2, 3, 4, 5}

    x := 9.4

    y, err := interpolation.Interpolation(x, pointX, pointY, interpolation.DefaultInterpolation)
    if err != nil {
        log.Printf("interpolation calc error: %v", err)
        return
    }
    log.Printf("x = %f, y = %f", x, y)
}

```