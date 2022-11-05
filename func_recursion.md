# 概述
经典语录: 要想理解递归，首先要理解递归。

递归的概念参考 [递归 - 维基百科](https://zh.wikipedia.org/wiki/%E9%80%92%E5%BD%92)。

# 例子

## 阶乘
```go
package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Printf("1 * 2 * 3 * 4 * 5 = %d\n", factorial(5))
}
// $ go run main.go
// 输出如下 
/**
    1 * 2 * 3 * 4 * 5 = 120
*/
```