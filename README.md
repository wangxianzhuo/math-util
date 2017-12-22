# Math Util

## Description

1. 提供插值法求点（当前支持三次样条插值）

## Usage
```shell
$ go get -u github.com/wangxianzhuo/math-util
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