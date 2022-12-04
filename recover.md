# 概述

建议先阅读 [defer](defer.md) 小节。

`recover` 会终止 `panic` 状态并且返回 `panic` 的值，函数会从 `panic` 之前执行到的地方直接返回，不会继续往下执行。

# 语法规则

**`recover` 和 `defer` 必须配套使用, 如果 `recover` 在其他地方执行会返回 `nil`，不会产生任何效果。
`defer` 必须在 `panic` 之前声明，否则 `panic` 会直接终止程序。**

# 例子

## 错误捕获

```go
package main

import "fmt"

func main() {
	if r := recover(); r != nil {
		fmt.Printf("捕获到 1 个错误: %v\n", r)
	}

	panic("测试")

	println("程序执行不到这里")
}

// $ go run main.go
// 输出如下 
/**
  panic: 测试
  ...
  ...
  exit status 2
*/
```

## 正确捕获

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到 1 个错误: %v\n", r)
		}
	}()

	panic("测试")

	println("程序执行不到这里")
}

// $ go run main.go
// 输出如下 
/**
  捕获到 1 个错误: 测试
*/
```