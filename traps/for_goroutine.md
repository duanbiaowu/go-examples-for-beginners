# 循环中 goroutine 执行顺序不一致

## 错误的做法

```go
package main

import "sync"

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			println(n)
		}(i)
	}

	wg.Wait()
}
// $ go run main.go
// 输出如下，顺序是乱序的，你的输出可能和这里的不一样，可以多试几次，看看效果
/**
    5 
    1 
    4 
    2 
    3
*/
```

**错误的原因在于**: 虽然 `goroutine` 是在循环中顺序启动的，但是其执行是并发的 (开始和结束时间不一定)，所以最终输出的结果中，也是乱序的。

## 正确的做法

知道错误的原因后，一个简单的解决方案是: 使用通道保证 `goroutine` 是顺序执行的，这样最终的输出结果一定是顺序的。

```go
package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	done := make([]chan bool, 5)

	// 初始化所有通道
	for i := range done {
		done[i] = make(chan bool)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer func() {
				wg.Done()
			}()
			
			x := <-done[n]  // 接收当前 `goroutine` 里面的通道的数据
			println(n)
			if n < 4 { // 向下一个 `goroutine` 里面的通道发送数据
				done[n+1] <- x
			}
		}(i)
	}

	done[0] <- true // 向第一个 `goroutine` 里面的通道发送数据

	wg.Wait()

	// 关闭所有通道
	for i := range done {
		done[i] = make(chan bool)
	}
}
// $ go run main.go
// 输出如下 
/**
    0
    1
    2
    3
    4
*/
```

**代码注解:** 在循环结束后，所有 `goroutine` 内部的通道都在阻塞接收数据，在接下来的操作中: 

1. 向第一个 `goroutine` 里面的通道发送数据
2. 第 1 个 `goroutine` 里面的通道接收到数据后，打印值，然后向第 2 个 `goroutine` 里面的通道发送数据
3. 第 2 个 `goroutine` 里面的通道接收到数据后，打印值，然后给第 3 个 `goroutine` 里面的通道发送数据 
4. ... 以此类推 
  
按照上面的描述依次类推，每一个 `goroutine` 里面通道都会依赖于前一个 `goroutine` 里面的通道，
最终所有的 `goroutine` 由 `乱序并发执行` 变为 `顺序串行执行`，所以输出结果当然也是顺序的。

# 扩展阅读

1. [Go 并发编程 - 数据竞态]()
2. [Go 陷阱 - goroutine 竞态](goroutine_race.md)
3. [Go 原子操作]()
4. [内存重排 - 维基百科](https://zh.wikipedia.org/zh-tw/%E5%86%85%E5%AD%98%E6%8E%92%E5%BA%8F)
5. [内存屏障 - 维基百科](https://zh.wikipedia.org/wiki/%E5%86%85%E5%AD%98%E5%B1%8F%E9%9A%9C)