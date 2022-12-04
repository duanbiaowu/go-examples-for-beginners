# 概述

使用的场景：**在函数内部有很多重复性代码并且严重依赖上下文变量**。此时可以在函数内部声明一个函数，专门用来处理重复性的代码。

# 例子

## 内部求和函数

```go
package main

import "fmt"

func main() {
	var sum func(...int) int // 声明 sum 函数

	sum = func(numbers ...int) int { // 定义 sum 函数
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