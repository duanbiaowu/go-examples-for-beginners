# 概述

虽然切片的底层是数组，但是当切片和数组作为函数的参数时，规则是不一样的。

## 数组传值不会改变原数组

### 错误的做法

```go
package main

import "fmt"

func double(arr [5]int) {
	for i := range arr {
		arr[i] *= 2
	}
}

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}
	for i := range numbers {
		fmt.Printf("%v ", numbers[i])
	}

	double(numbers)
	fmt.Println()

	for i := range numbers {
		fmt.Printf("%v ", numbers[i])
	}
}

// $ go run main.go
// 输出如下
/**
  1 2 3 4 5
  1 2 3 4 5
*/
```

从输出结果中看到，数值元素的值并没有被修改。

### 正确的做法

通过传递数组指针来改变元素的值。

```go
package main

import "fmt"

func double(arr *[5]int) {
	for i := range arr {
		arr[i] *= 2
	}
}

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}
	for i := range numbers {
		fmt.Printf("%v ", numbers[i])
	}

	double(&numbers)
	fmt.Println()

	for i := range numbers {
		fmt.Printf("%v ", numbers[i])
	}
}

// $ go run main.go
// 输出如下 
/**
  1 2 3 4 5
  2 4 6 8 10
*/
```

从输出结果中看到，数值元素的值已经被修改。

## 切片传值可以改变原数组

```go
package main

import "fmt"

func double(arr []int) {
	for i := range arr {
		arr[i] *= 2
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for i := range numbers {
		fmt.Printf("%v ", numbers[i])
	}

	double(numbers)
	fmt.Println()

	for i := range numbers {
		fmt.Printf("%v ", numbers[i])
	}
}

// $ go run main.go
// 输出如下 
/**
  1 2 3 4 5
  2 4 6 8 10
*/
```
