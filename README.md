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

## Http Server API

### Start server

- 在`conf/app.conf`中配置服务器相关参数

### 计算行列式值

- URL `/det/value`
- Method `GET`
- Param `elements=[[1,2,3],[4,5,6],[7,8,9]]`
- Retrun `{"value":0,"des":""}`

### 行列式展开

- URL `/det/expanse`
- Method `GET`
- Param `elements=[[1,2,3],[4,5,6],[7,8,9]]&order=0&cowRowType=0`
- Retrun `{"value":[{"Elements":[[5,6],[8,9]],"Rank":2,"Factor":1},{"Elements":[[4,6],[7,9]],"Rank":2,"Factor":-2},{"Elements":[[4,5],[7,8]],"Rank":2,"Factor":3}],"des":""}`

### 解方程组

- URL: `/equations/solve`
- Method: `GET`
- Param: `factors=[[2,1,-5,1],[1,-3,0,-6],[0,2,-1,2],[1,4,-7,6]]&values=[8,9,-5,0]&unknownNum=4`
- Return: `{"value":[3,-4,-1,1],"des":"Onlyonesolution"}`

### 计算逆序数

- URL: `/util/inversionnumber`
- Method: `GET`
- Param: `list=[4,3,5,6,23,7,2,54,9]`
- Return: `{"value": 10,"des": ""}`