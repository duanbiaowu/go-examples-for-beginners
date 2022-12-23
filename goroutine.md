# 概述

**goroutine 是 Go 程序并发执行的实体**，对于初学者来讲，可以简单地将 `goroutine` 理解为一个 `超轻量的线程`。

当一个程序启动时，只有一个 goroutine 调用 main 函数，称为 `主 goroutine`, 当 main 函数返回时，
所有 `goroutine` 都会被终止 (不论其是否运行完成)，然后程序退出。

# 语法规则

关键字 `go` 启动一个 `goroutine` (可以理解为在后台运行一个函数), 需要注意的是: 使用 **`go` 启动的函数没有返回值**。

```shell
# 直接调用一个匿名函数
go func() { // 无参数
    // do something ...
}

go func(x int, y bool ...) { // 有参数
    // do something ...
}
```

```shell
# 调用一个已定义的函数
go foo()   // 无参数

go bar(x int, y bool ...)  // 有参数
```

# 例子

## 直接调用一个匿名函数

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 3 个 goroutine 是并发运行的，所以顺序不一定是 1, 2, 3
	// 读者可以多运行几次，看看输出结果

	go func() {
		fmt.Println("goroutine 1")
	}()

	go func() {
		fmt.Println("goroutine 2")
	}()

	go func() {
		fmt.Println("goroutine 3")
	}()

	// 这一行代码不可省略
	// 如果省略掉，意味着主进程不等待 3 个 goroutine 执行完成就退出了，也就不会有 goroutine 的输出信息了
	// 读者可以注释掉这行代码，然后运行看看输出结果
	time.Sleep(1 * time.Second)
}

// $ go run main.go
// 输出如下， 3 个 goroutine 是并发运行的，顺序不一定，所以你的输出可能和这里的不一样
/**
  goroutine 3
  goroutine 1
  goroutine 2
*/
```

调用 `time.Sleep()` 睡眠等待 3 个 goroutine 执行完成，虽然达到了演示效果，但是有很多潜在问题。
更好的解决方案请看 [waitgroup](waitgroup.md)。

## 调用一个已定义的函数

```go
package main

import (
	"fmt"
	"time"
)

func foo() {
	fmt.Println("goroutine foo")
}

func bar() {
	fmt.Println("goroutine bar")
}

func fooBar(s string) {
	fmt.Printf("goroutine %s\n", s)
}

func main() {
	// 3 个 goroutine 是并发运行的，所以顺序不一定是 1, 2, 3
	// 读者可以多运行几次，看看输出结果

	go foo()

	go bar()

	go fooBar("fooBar")

	// 这一行代码不可省略
	// 如果省略掉，意味着主进程不等待 3 个 goroutine 执行完成就退出了，也就不会有 goroutine 的输出信息了
	// 读者可以注释掉这行代码，然后运行看看输出结果
	time.Sleep(1 * time.Second) 
}

// $ go run main.go
// 输出如下， 3 个 goroutine 是并发运行的，顺序不一定，所以你的输出可能和这里的不一样
/**
  goroutine fooBar
  goroutine foo
  goroutine bar
*/
```

## 获取并发线程数量

`GOMAXPROCS()` 获得并发的线程数量，在 CPU 核大于 1 个的情况下，系统会尽可能调度等于核心数的线程并行运行。

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("GOMAXPROCS = %d\n", runtime.GOMAXPROCS(0))
}

// $ go run main.go
// 输出如下，笔者的机器 CPU 是 8 核，你的输出可能和这里的不一样
/**
  GOMAXPROCS = 8
*/
```

# 扩展阅读

1. [协程 - 维基百科](https://zh.wikipedia.org/wiki/%E5%8D%8F%E7%A8%8B)
2. [线程 - 维基百科](https://zh.wikipedia.org/wiki/%E7%BA%BF%E7%A8%8B)
3. [Go 圣经 - 第 8 章](https://book.douban.com/subject/27044219/)