# 概述

# 语法规则
recover 必须和 defer 配套使用, defer 和 panic 的顺序非常重要。

# 例子

## 内部求和函数
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