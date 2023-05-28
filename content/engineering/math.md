---
title: 常用数学方法
date: 2023-01-01
---

## 保留两位小数

```go
package main

import (
	"fmt"
	"math"
)

func RoundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
	fmt.Println(RoundFloat(3.1415926, 2))
	fmt.Println(RoundFloat(1024.2325, 1))
}
```

运行代码输出如下

```shell
$  go run main.go

3.14
1024.2
```