# 概述

相较于主流编程语言，Go 中的 `switch` 语法更加灵活，它接受任意形式的表达式。

# 语法规则

**`switch` 后面的表达式是不需要括号的。**

**`case` 语句块执行完会自动退出整个 `switch` 语句块，也就是不需要使用 `break` 显式声明。**
如果在执行完对应 `case` 语句后，希望继续向下执行，可以使用关键字 `fallthrough`, 这样就和其他编程语言不加 `break` 效果一样了。

```shell
switch expr {   // expr 可以是任意类型
  case v1:
      ...
  case v2:
      ...
  case v3:
      ...
  case v4, v5, v6:    // 可以同时测试多个可能符合条件的值
      ...
  default:  // 默认值
      ...
}
```

# 例子

## 普通表达式

```go
package main

import "fmt"

func main() {
	n := 1024
	switch n {
	case 1023:
		fmt.Println("n = 1023")
	case 1024:
		fmt.Println("n = 1024")
	case 1025:
		fmt.Println("n = 1025")
	}
}

// $ go run main.go
// 输出如下 
/**
  n = 1024
*/
```

## 运算表达式

```go
package main

import "fmt"

func main() {
	n := 1024
	switch n * 2 {
	case 1024:
		fmt.Println("n = 1024")
	case 2048:
		fmt.Println("n = 2048")
	case 0:
		fmt.Println("n = 0")
	}
}

// $ go run main.go
// 输出如下 
/**
  n = 2048
*/
```

## default

```go
package main

import "fmt"

func main() {
	n := 1024
	switch n * 2 {
	case 0:
		fmt.Println("n = 0")
	case 1:
		fmt.Println("n = 1")
	case 2:
		fmt.Println("n = 2")
	default:
		fmt.Println("n = 2048")
	}
}

// $ go run main.go
// 输出如下 
/**
  n = 2048
*/
```

## 省略 expr 表达式

```go
package main

import "fmt"

func main() {
	n := 1024
	switch {
	case n < 1024:
		fmt.Println("n < 1024")
	case n > 1024:
		fmt.Println("n > 1024")
	case n == 1024:
		fmt.Println("n == 1024")
	default:
		fmt.Println("invalid n")
	}
}

// $ go run main.go
// 输出如下 
/**
  n = 1024
*/
```

## 同时测试多个 case

```go
package main

import "fmt"

func main() {
	n := 1024
	switch n {
	case 1023, 1024: // 多个 case, 只要一个匹配就 OK
		fmt.Println("n <= 1024")
	case 1025:
		fmt.Println("n > 1024")
	default:
		fmt.Println("invalid n")
	}
}

// $ go run main.go
// 输出如下 
/**
  n <= 1024
*/
```

## fallthrough

```go
package main

import "fmt"

func main() {
	n := 1024
	switch {
	case n < 1024:
		fmt.Println("n < 1024")
		fallthrough // 继续向下执行
	case n > 1024:
		fmt.Println("n > 1024")
		fallthrough // 继续向下执行
	case n == 1024:
		fmt.Println("n == 1024")
		fallthrough // 继续向下执行
	default:
		fmt.Println("invalid n")
	}
}

// $ go run main.go
// 输出如下 
/**
  n = 1024
  invalid n
*/
```

## 类型断言

`switch` 可以 `.(type)` 来实现类型断言，判断某个值是否为某个数据类型。

```go
package main

func main() {
	var n interface{} // 值必须为 interface 类型
	n = 1024

	switch n.(type) {
	case nil:
		println("n is a nil")
	case int:
		println("n is a integer")
	case float64:
		println("n is a float64")
	case bool:
		println("n is a bool")
	case string:
		println("n is a string")
	default:
		println("n is invalid")
	}
}

// $ go run main.go
// 输出如下 
/**
  n is a integer
*/
```

关于上面代码注释中提到的 `interface`, [后面会讲到](interface.md), 读者这里只需知道关键字 `interface` 代表接口类型即可。