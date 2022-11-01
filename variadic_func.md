# 概述

# 语法规则

# 例子

## 传递一个参数
```go
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Printf("1 = %d\n", sum(1))
}
// $ go run main.go
// 输出如下 
/**
    1 = 1
*/
```

## 传递多个参数
```go
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Printf("1 + 2 + 3 = %d\n", sum(1, 2, 3))
}
// $ go run main.go
// 输出如下 
/**
    1 + 2 + 3 = 6
*/
```

## 传递切片参数
```go
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	numbers := []int{1, 2, 3}
	fmt.Printf("1 + 2 + 3 = %d\n", sum(numbers...)) // 切片变量后面加 ... 即可
}
// $ go run main.go
// 输出如下 
/**
    1 + 2 + 3 = 6
*/
```

## 不传递任何参数
```go
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Printf("不传递任何参数 = %d\n", sum())
}
// $ go run main.go
// 输出如下 
/**
    不传递任何参数 = 0
*/
```