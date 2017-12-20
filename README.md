# Math Util

## Description

1. 提供插值法求点（当前支持三次样条插值）

## Usage

```go
package main

import util "github.com/wangxianzhuo/math-util"
import log

func main() {
    pointX := []float64{1, 2, 3, 4, 5}
    pointY := []float64{1, 2, 3, 4, 5}

    x := 9.4

    y, err := util.Interpolation(x, pointX, pointY, util.DefaultInterpolation)
    if err != nil {
        log.Printf("interpolation calc error: %v", err)
        return
    }
    log.Printf("x = %f, y = %f", x, y)
}

```