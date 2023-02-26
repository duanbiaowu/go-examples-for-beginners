# 概述

`函数` 是将一个或者一类问题包装为一个代码块，可以被多次调用，提高代码重用性。

Go 函数中声明、定义、参数、返回值这些基础概念，和其他编程语言中的一致，这里不再赘述。

# 语法规则

**Go 函数支持单个返回值和多个返回值。**

```shell
# 单个返回值
# 参数可省略
func 函数名称(参数 1 值 参数 1 类型, 参数 2 值 参数 2 类型 ...) 返回值类型 {
    // do something
}

# 多个返回值，不指定名称
# 参数可省略
func 函数名称(参数 1 值 参数 1 类型, 参数 2 值 参数 2 类型 ...) (返回值 1 类型, 返回值 2 类型) {
    // do something
}

# 多个返回值，指定名称
# 参数可省略
func 函数名称(参数 1 值 参数 1 类型, 参数 2 值 参数 2 类型 ...) (返回值 1 名称 返回值 1 类型, 返回值 2 名称 返回值 2 类型) {
    // do something
}
```

# 例子

## 单个返回值

```go
package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Printf("max = %d\n", max(1, 2))
}

// $ go run main.go
// 输出如下 
/**
  max = 2
*/
```

## 多个返回值，不指定名称

```go
package main

import "fmt"

func getNumbers() (int, int, int) {
	return 1, 2, 3
}

func main() {
	x, y, z := getNumbers()
	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
}

// $ go run main.go
// 输出如下 
/**
  x = 1, y = 2, z = 3
*/
```

## 多个返回值，指定名称

```go
package main

import "fmt"

func getNumbers() (x int, y float64, z string) {
	x = 1024
	y = 3.14
	z = "hello world"
	return
}

func main() {
	x, y, z := getNumbers()
	fmt.Printf("x = %d, y = %.2f, z = %s\n", x, y, z)
}

// $ go run main.go
// 输出如下 
/**
  x = 1024, y = 3.14, z = hello world
*/
```

## 参数/返回值 类型相同简化

* 当参数类型相同时，可以将类型放在最后一个参数变量后面
* 当返回值类型相同时，可以将类型放在最后一个返回值变量后面

```go
package main

import "fmt"

func getMaxAndMin(x, y, z int) (max, min int) {
	if x > y {
		max = x
		min = y
		if x < z {
			max = z
		} else if z < y {
			min = z
		}
	} else {
		max = y
		min = x
		if y < z {
			max = z
		} else if z < x {
			min = x
		}
	}
	return
}

func main() {
	max, min := getMaxAndMin(100, 200, 300)
	fmt.Printf("max = %d, min = %d\n", max, min)
}

// $ go run main.go
// 输出如下 
/**
  max = 300, min = 100
*/
```

# 扩展阅读

1. https://zh.wikipedia.org/wiki/%E5%87%BD%E6%95%B0%E5%8E%9F%E5%9E%8B