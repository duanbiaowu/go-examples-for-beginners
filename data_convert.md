# 概述
**Go 是强类型语言，因此不会进行隐式类型转换 (例子可以直接将一个 `浮点型` 转换为 `整型`)。任何不同类型之间的转换都必须显式说明。**
在类型转换时，要注意两边的值类型大小，可以将一个较小的值类型转换为一个较大的值类型，但是反过来，却有可能报错。
例如：将一个 `int64` 转换为 `int32` 时，可能会产生 [溢出](https://zh.wikipedia.org/wiki/%E7%AE%97%E8%A1%93%E6%BA%A2%E5%87%BA)。

# 直接转换
`直接转换` 适用于两边值类型相同的情况，比如 `int` 和 `float`, 都是数字类型。

## 语法规则
```shell
```shell
变量名称 = 数据类型(变量值)
# 例子
n := 123
floatN := float64(n)  // 将 n 转换为 float64 类型       
```

## 例子
### 整数转换为浮点数
```go
package main

import "fmt"

func main() {
	n := 100
	fmt.Println(float64(n))
}

// $ go run main.go
// 输出如下 
/**
    100
*/
```

### 浮点数转换为整数
```go
package main

import "fmt"

func main() {
	pi := 3.14
	fmt.Println(int64(pi))
}

// $ go run main.go
// 输出如下 
/**
    3
*/
```

# 调用方法转换
对于两边值类型不相同的情况，比如 `int` 和 `string`，无法进行直接转换，比如下面的代码是错误的: 
```go
package main

import "fmt"

func main() {
	pi := 3.14
	fmt.Println(string(pi))
}

// 运行报错: Cannot convert an expression of the type 'float64' to the type 'string'
```

这时候，需要调用一些内置的方法进行转换，这里介绍几个常用的 `数字` 和 `字符串` 之间转换用到的方法。

## 常用方法
* `strconv.Itoa()`         将 `int` 转换为 `string`
* `strconv.Atoi()`         将 `string` 转换为 `int`
* `strconv.FormatFloat()`  将 `float64` 转换为 `string`
* `strconv.ParseFloat()`   将 `string` 转换为 `float64`

### 例子
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Itoa(1024)
	fmt.Printf("%T, %v\n", s, s) // 将整型转换为字符串

	n, _ := strconv.Atoi("1024")
	fmt.Printf("%T, %v\n", n, n) // 将字符串转换为整型

	s2 := strconv.FormatFloat(3.1415, 'f', -1, 64)
	fmt.Printf("%T, %v\n", s2, s2) // 将浮点型转换为字符串

	n2, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("%T, %v\n", n2, n2) // 将字符串转换为浮点型
}
// $ go run main.go
// 输出如下 
/**
    string, 1024
    int, 1024
    string, 3.1415
    float64, 3.1415
*/
```