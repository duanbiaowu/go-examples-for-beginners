# 概述

# 语法规则

# 例子

## 自增序列号生成器
```go
package main

import "fmt"

// 自增序列号生成器
func incSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := incSeq()

	fmt.Printf("初始值 = %d\n", next())

	for i := 1; i <= 5; i++ {
		fmt.Printf("第 %d 次迭代后, 值 = %d\n", i, next())
	}
}
// $ go run main.go
// 输出如下 
/**
    初始值 = 1
    第 1 次迭代后, 值 = 2
    第 2 次迭代后, 值 = 3
    第 3 次迭代后, 值 = 4
    第 4 次迭代后, 值 = 5
    第 5 次迭代后, 值 = 6
*/
```