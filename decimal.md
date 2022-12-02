# 概述

在大多数处理浮点数的场景中，为了提高可读性，往往只需要精确到 2 位或 3 位，一般来说，常用的方法有两种。

## fmt.Sprintf()

```go
package main

import "fmt"

func main() {
	pi := 3.1415926
	s1 := fmt.Sprintf("%.2f", pi) // 保留 2 位小数
	fmt.Printf("%T %v\n", s1, s1)
    
	s2 := fmt.Sprintf("%.1f", pi) // 保留 1 位小数
	fmt.Printf("%T %v\n", s2, s2)
}

// $ go run main.go
// 输出如下 
/**
    string 3.14
    string 3.1
*/
```

通过调用 `fmt.Sprintf()` 方法转换非常简单，但是**不足之处在于返回值是一个字符串**，
如果需要保留精度的值依然要求为 `浮点型`，可能需要使用二次 [类型转换](data_convert.md)，不太友好。

## math 包
本质上是通过两个 `浮点型` 数字进行计算，最后根据需要的精度，进行四舍五入。 例如 `12345 / 100 = 123.45`, 保留 1 位小数等于 `123.5`。 

```go

package main

import (
	"fmt"
	"math"
)

func main() {
	pi := 3.1415926     
	var ratio float64   // 使用一个变量作为精度范围, 比如 2 位小数时，精度范围应该为 100 

	ratio = math.Pow(10, 2)             // 计算精度范围，2 位小数 = 100
	s1 := math.Round(pi*ratio) / ratio // 保留 2 位小数
	fmt.Printf("%T %v\n", s1, s1)

	ratio = math.Pow(10, 1)             // 计算精度范围，1 位小数 = 10
	s2 := math.Round(pi*ratio) / ratio // 保留 1 位小数
	fmt.Printf("%T %v\n", s2, s2)
}
```

通过调用 `math 包` 方法转换，除了可以保留精度外，还可以保证转换后的值依然是一个 `浮点型`。
但是该方法在一些 `边缘场景` 中，可能会报错，例如需要转化的数值已经是一个 `很大的浮点型数字`，这时候再乘以精度范围值，
可能会产生 [溢出](https://zh.wikipedia.org/wiki/%E7%AE%97%E8%A1%93%E6%BA%A2%E5%87%BA)。

## 实践

如果对场景产生的数据没办法做到 `100%` 预测，建议:
1. 使用 `fmt.Sprintf()` 将 `数字` 转为 `字符串`
2. 使用 [类型转换](data_convert.md)，将 `字符串` 转化为 `浮点型`