# 概述

`nil` 可以作为函数参数传入，这意味着函数内部逻辑处理时，不能依赖于传入的实参 (有可能是 nil)， 一定要做必要的 `零值` 判断。

# 例子

示例代码只是为了演示，没有任何实际意义。

## 参数类型为切片

当 `切片` 为 `nil` 时，直接读取和赋值都会 `panic` 。

### 错误的做法

```go
package main

// 计算前 N 个数总和
func sumTopN(numbers []int, n int) int {
	total := 0
	for _, v := range numbers[:n] {
		total += v
	}

	return total
}

func main() {
	println(sumTopN(nil, 3))
}

// $ go run main.go
// 输出如下
/**
  panic: runtime error: slice bounds out of range [:3] with capacity 0

  goroutine 1 [running]:
  ...
  ...

  exit status 2

*/
```

### 正确的做法

```go
package main

// 计算前 N 个数总和
func sumTopN(numbers []int, n int) int {
	// 首先检测参数是否为 nil
	if n > len(numbers) {
		n = len(numbers)
	}

	total := 0
	for _, v := range numbers[:n] {
		total += v
	}

	return total
}

func main() {
	println(sumTopN(nil, 3))
}

// $ go run main.go
// 输出如下 
/**
  0
*/
```

## 参数类型为 map

当 `map` 为 `nil` 时，直接赋值会 `panic` 。

### 错误的做法

```go
package main

// 将目标字符串计数重置为 0
func reset(counter map[string]int, target []string) {
	for _, s := range target {
		counter[s] = 0
	}
}

func main() {
	reset(nil, []string{"hello", "world"})
}

// $ go run main.go
// 输出如下
/**
  panic: assignment to entry in nil map

  goroutine 1 [running]:
  ...
  ...

  exit status 2

*/
```

### 正确的做法

```go
package main

// 将目标字符串计数重置为 0
func reset(counter map[string]int, target []string) {
	// 首先检测参数是否为 nil
	if counter == nil {
		return
	}
	for _, s := range target {
		counter[s] = 0
	}
}

func main() {
	reset(nil, []string{"hello", "world"})
}
```

## 参数类型为指针

当 `指针` 为 `nil` 时，直接读取和赋值都会 `panic` 。

### 错误的做法

```go
package main

import "fmt"

type person struct {
	name string
}

func setHi(p *person) {
	fmt.Printf("Hi, I'm %s", p.name)
}

func main() {
	setHi(nil)
}

// $ go run main.go
// 输出如下
/**
  panic: runtime error: invalid memory address or nil pointer dereference

  goroutine 1 [running]:
  ...
  ...

  exit status 2

*/
```

### 正确的做法

```go
package main

import "fmt"

type person struct {
	name string
}

func setHi(p *person) {
	// 首先检测参数是否为 nil
	if p != nil {
		fmt.Printf("Hi, I'm %s", p.name)
	}
}

func main() {
	setHi(nil)
}
```

> 代码的首要目标是正确性。