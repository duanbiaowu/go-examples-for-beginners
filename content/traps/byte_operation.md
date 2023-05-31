---
date: 2023-01-01
title: Go 陷阱之 byte 加减
modify: 2023-01-01
---

# 概述

标准库中的 `byte` 类型是 `uint8` 类型的别名，在所有方面都相当于 `uint8`，主要作用是用来区分字节类型和无符号整数类型。

**两个 `byte` 值使用 `+` 相加，并不会产生 `字符拼接` 的效果形成字符串，相反，会先将两个 `byte` 值转换成对应的无符号整数类型，
然后进行相加，最后的结果是一个整数。**

# 例子

## 错误的做法

```go
package main

import "fmt"

func main() {
	a, b := 'a', 'b'
	c := a + b
	fmt.Printf("a type = %T, val = %v\n", a, a)
	fmt.Printf("b type = %T, val = %v\n", b, a)
	fmt.Printf("c type = %T, val = %v\n", c, c)
}
```

```shell
$ go run main.go

# 输出如下
a type = int32, val = 97
b type = int32, val = 97 
c type = int32, val = 195
```

通过输出结果可以看到，字符 `a + b` 没有得到预料之中的结果 `ab`, 而是先分别将 `a` 和 `b` 转换为数字 `97`, `98`，然后相加得到结果 195。 

## 正确的做法

有两种方法解决上述问题: 
1. 先将 `byte` 类型转换为 `string` 类型，然后再拼接
2. 使用 fmt.Sprintf 方法的格式化功能，将 `byte` 类型转换为 `string` 类型

### 先转字符串再拼接

```go
package main

import "fmt"

func main() {
	a, b := 'a', 'b'
	c := string(a) + string(b)
	fmt.Printf("c type = %T, val = %v\n", c, c)
}
```

```shell
$ go run main.go

# 输出如下 
c type = string, val = ab
```

### 使用 fmt.Sprintf 方法

```go
package main

import "fmt"

func main() {
	a, b := 'a', 'b'
	c := fmt.Sprintf("%c%c", a, b)
	fmt.Printf("c type = %T, val = %v\n", c, c)
}
```

```shell
$ go run main.go

# 输出如下 
c type = string, val = ab
```