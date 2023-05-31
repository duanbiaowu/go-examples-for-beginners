---
date: 2023-01-01
title: Go 陷阱之 goroutine 泄漏
modify: 2023-01-01
---

# 通道为 nil 造成 goroutine 泄漏

在 `nil 通道` 上发送和接收操作将永久阻塞，造成 `goroutine 泄漏`。

> 最佳实践: 1. 永远不要对 `nil 通道` 进行任何操作，2. 直接使用 `make()` 初始化通道。

## 接收造成的泄漏

示例代码只是为了演示，没有任何实际意义。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan bool

	go func() {
		defer func() { // defer 不会执行
			fmt.Println("goroutine ending") // 不会输出
		}()

		for v := range ch {
			fmt.Println(v)
		}

		fmt.Println("range broken") // 执行不到这里
	}()
	
	time.Sleep(time.Second) // 假设主程序 1 秒后退出
}

// $ go run main.go
// 没有任何输出，goroutine 泄漏
```

## 发送造成的泄漏

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan bool

	go func() {
		defer func() { // defer 不会执行
			fmt.Println("goroutine ending") // 不会输出
		}()

		ch <- true

		fmt.Println("range broken") // 执行不到这里
	}()

	time.Sleep(time.Second) // 假设主程序 1 秒后退出
}

// $ go run main.go
// 没有任何输出，goroutine 泄漏
```

# 遍历未关闭通道时造成 goroutine 泄漏

遍历 `无缓冲 (阻塞) 并且未关闭` 的通道时，如果通道一直未关闭， 将会永久阻塞，造成 `goroutine 泄漏`。

遍历 `缓冲 (非阻塞) 并且未关闭` 的通道时，将通道内的所有缓存数据接收完毕后， 如果通道一直未关闭，将会永久阻塞，造成 `goroutine 泄漏`。

> 最佳实践: 1. 确保 `通道` 可以正常关闭，2. 确保 `goroutine` 可以正常退出。

## 遍历无缓冲并且未关闭的通道

示例代码只是为了演示，没有任何实际意义。

### 错误的做法

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func() {
		defer func() { // defer 不会执行
			fmt.Println("goroutine ending") // 不会输出
		}()

		for v := range ch {
			fmt.Println(v)
			break
		}

		fmt.Println("range broken") // 执行不到这里
	}()

	time.Sleep(time.Second) // 假设主程序 1 秒后退出
}

// $ go run main.go
// 没有任何输出，goroutine 泄漏
```

### 正确的做法

参照最佳实践，对代码进行以下调整: 在 `goroutine` 外部关闭通道，防止 `goroutine` 内部遍历陷入无限阻塞。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func() {
		defer func() { // defer 正常执行
			fmt.Println("goroutine ending") // 正常输出
		}()

		for v := range ch { // 外部关闭通道后，for 循环结束
			fmt.Println(v) // 不会输出
		}

		fmt.Println("range broken") // 可以执行到这里
	}()

	close(ch) // 关闭通道，内存遍历循环立即结束

	time.Sleep(time.Second) // 假设主程序 1 秒后退出
}

// $ go run main.go
// 输出如下
/**
  range broken
  goroutine ending
*/
```

## 遍历缓冲并且未关闭的通道

示例代码只是为了演示，没有任何实际意义。

### 错误的做法

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool, 3)

	go func() {
		defer func() { // defer 不会执行
			fmt.Println("goroutine ending") // 不会输出
		}()

		for v := range ch {
			fmt.Println(v)
		}

		fmt.Println("range broken") // 执行不到这里
	}()

	ch <- true
	ch <- false
	ch <- true

	time.Sleep(time.Second) // 假设主程序 1 秒后退出
}

// $ go run main.go
// 输出如下
/**
  true
  false
  true
  // 接收完缓冲区的 3 个值后, 后面不再有任何输出，goroutine 泄漏
*/
```

### 正确的做法

参照最佳实践，对代码进行以下调整: 在 `goroutine` 外部关闭通道，防止 `goroutine` 内部遍历陷入无限阻塞。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func() {
		defer func() { // defer 正常执行
			fmt.Println("goroutine ending") // 正常输出
		}()

		for v := range ch { // 外部关闭通道后，for 循环结束
			fmt.Println(v) // 不会输出
		}

		fmt.Println("range broken") // 可以执行到这里
	}()

	close(ch) // 关闭通道，内存遍历循环立即结束

	time.Sleep(time.Second) // 假设主程序 1 秒后退出
}

// $ go run main.go
// 输出如下
/**
  true
  false
  true
  range broken
  goroutine ending
*/
```

