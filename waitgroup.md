# 概述
建议先阅读 [goroutine](goroutine.md)。

在 [goroutine](goroutine.md) 小节中，为了让并发的 3 个 `goroutine` 正常执行完成，调用 `time.Sleep()` 睡眠等待。
这样的方式存除了实现不优雅之外，最大的问题在于: time.Sleep() 接受的是一个硬编码的时间参数，这就要求我们实现必须知道每个
 goroutine 的执行时间并且要以执行时间最长的 goroutine 为基准，这在大多数场景下是没办法做到的。

如果主进程能够知道每个 `goroutine` 是何时结束的，并且在结束之后发一个通知给主进程， 那么问题就可以完美解决了。
Go 提供的 `sync.WaitGroup` 就是针对这种场景的解决方案。 

# 调用规则
* `Add()` 和 `Done()` 方法必须配对，`Wait()` 方法必须在程序退出前调用。
* `Add()`, `Done()`, `Wait()` 三者必须同属一个作用域。

# 例子

读者可以多运行几次下面的例子，体会 `sync.WaitGroup` 的用法和细节。

## 调用 3 个 goroutine
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // 声明一个 sync.WaitGroup 实例
	wg.Add(3)             // 参数 为 3，正好对应了 3 个 goroutine

	// 3 个 goroutine 是并发运行的，所以顺序不一定是 1, 2, 3
	// 读者可以多运行几次，看看输出结果
	go func() {
		defer func() {
			wg.Done() //  通知主进程，这个 goroutine 已经执行完成
		}()
		fmt.Println("goroutine 1")
	}()

	go func() {
		defer func() {
			wg.Done() //  通知主进程，这个 goroutine 已经执行完成
		}()
		fmt.Println("goroutine 2")
	}()

	go func() {
		defer func() {
			wg.Done() //  通知主进程，这个 goroutine 已经执行完成
		}()
		fmt.Println("goroutine 3")
	}()

	wg.Wait() // 等待所有 goroutine 全部执行完
}
// $ go run main.go
// 输出如下， 3 个 goroutine 是并发运行的，顺序不一定，所以你的输出可能和这里的不一样
/**
    goroutine 3
    goroutine 1
    goroutine 2
*/
```
