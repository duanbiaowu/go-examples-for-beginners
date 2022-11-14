# 概述

当一个变量使用 `var` 进行声明后并未进行初始化 (变量后面没有赋值符 `=`) 操作，会默认分配一个零值 (zero value)。

# 不同类型对应的零值

| 类型        | 零值    |
|-----------|-------|
| bool      | false |
| int       | 0     |
| float     | 0.0   |
| string    | ""    |
| byte      | ''    |
| pointer   | nil   |
| func      | nil   |
| interface | nil   |
| slice     | nil   |
| channel   | nil   |
| map       | nil   |

# 例子

## bool 类型

```go
package main

import "fmt"

func main() {
	var b bool
	fmt.Printf("b = %t\n", b)
}

// $ go run main.go
// 输出如下
/**
  b = false
*/
```

## byte 类型

```go
package main

import "fmt"

func main() {
	var b byte
	fmt.Printf("b = %c\n", b)
}

// $ go run main.go
// 输出如下
/**
  b =
*/
```

## func 类型

```go
package main

import "fmt"

func main() {
	var f func()
	fmt.Printf("f = %v\n", f)
}

// $ go run main.go
// 输出如下
/**
  f = <nil>
*/
```

## channel 类型

```go
package main

import "fmt"

func main() {
	var ch chan bool
	fmt.Printf("ch = %v\n", ch)
}
// $ go run main.go
// 输出如下
/**
    ch = <nil>
*/
```
