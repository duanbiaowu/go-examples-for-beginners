# 概述

# 语法规则
先声明，后定义

# 例子

## 阶乘
```go
package main

import "fmt"

func main() {
	var sum func(...int) int // 声明 sum 函数

	sum = func(numbers ...int) int {	// 定义 sun 函数
		total := 0
		for _, num := range numbers {
			total += num
		}
		return total
	}

	fmt.Printf("1 + 2 + 3 = %d\n", sum(1, 2, 3))
}
// $ go run main.go
// 输出如下 
/**
    1 + 2 + 3 = 6
*/
```