# 概述
建议先阅读 [指针](pointer.md) 。

# 例子

## 指针变量参数
```go
package main

import "fmt"

func double(n int) {
	n *= 2
}

func doubleWithPtr(n *int) {
	*n *= 2
}

func main() {
	n := 100

	double(n)
	fmt.Printf("使用普通变量作为函数参数执行完成后, n = %d\n", n)    // 可以看到，变量值并未发生变化

	doubleWithPtr(&n)
	fmt.Printf("使用指针变量作为函数参数执行完成后, n = %d\n", n)    // 变量值已经发生变化
}
// $ go run main.go
// 输出如下 
/**
    使用普通变量作为函数参数执行完成后, n = 100
    使用指针变量作为函数参数执行完成后, n = 200
 */
```

## 指针数组变量参数
```go
package main

import "fmt"

func double(numbers [3]int) {
	for _, v := range numbers {
		v *= 2
	}
}

func doubleWithPtr(numbers *[3]int) {
	for i := range numbers {
		numbers[i] *= 2
	}
}

func main() {
	numbers := [3]int{100, 200, 300}

	double(numbers)
	fmt.Printf("使用数组变量作为函数参数执行完成后, n = %d\n", numbers)  // 可以看到，数组元素并未发生变化

	doubleWithPtr(&numbers)
	fmt.Printf("使用指针数组变量作为函数参数执行完成后, n = %d\n", numbers)    // 数组元素已经发生变化
}
// $ go run main.go
// 输出如下 
/**
    使用数组变量作为函数参数执行完成后, n = [100 200 300]
    使用指针数组变量作为函数参数执行完成后, n = [200 400 600]
*/
```

## 切片参数
在 [切片](slice.md) 小节中讲到，切片的底层引用了一个数组，可以简单地理解为：**切片本身是一个指针，指向底层数组的元素**，
所以常用的方式的是将函数参数定义为切片类型。
```go
package main

import "fmt"

func double(numbers []int) {
	for i := range numbers {
		numbers[i] *= 2
	}
}

func main() {
	numbers := []int{100, 200, 300}

	double(numbers)
	fmt.Printf("使用切片变量作为函数参数执行完成后, n = %d\n", numbers)  // 切片元素已经发生变化
}
// $ go run main.go
// 输出如下 
/**
    使用切片变量作为函数参数执行完成后, n = [200 400 600]
*/
```


