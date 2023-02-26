# main 结束时不考虑 goroutine 执行状态

默认情况下，主程序结束时不会考虑当前是否还有 `goroutine` 正在执行。

## 错误的做法

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			defer func() {
				fmt.Println("goroutine ending")
			}()

			time.Sleep(100 * time.Millisecond) // 模拟耗时操作
		}()
	}

	fmt.Println("main ending")
}

// $ go run main.go
// 输出如下
/**
  main ending
*/
```

从输出结果中看到，只有 `main()` 输出的字符串， 3 个 `goroutine` 没有输出任何字符串。

## 正确的做法

使用 `sync.WaitGroup` **同步原语** 保证主程序结束前所有 `goroutine` 正常退出。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				fmt.Println("goroutine ending")
				wg.Done()
			}()

			time.Sleep(100 * time.Millisecond) // 模拟耗时操作
		}()
	}

	wg.Wait()
	fmt.Println("main ending")
}

// $ go run main.go
// 输出如下 
/**
  goroutine ending
  goroutine ending
  goroutine ending
  main ending
*/
```

# WaitGroup 与 goroutine 竞态

## 错误的做法

```go
package main

import (
	"sync"
)

func main() {
	var wg = &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			println("hello world")
		}()
	}

	wg.Wait()
}

// $ go run main.go
// 没有任何输出
```

**错误的原因在于**: 在循环中启动 `goroutine` 是异步的且需要一定的时间 (虽然这个时间很短)，
接下来循环结束后执行到 `wg.Wait()` 时，循环内部还没有任何 `wg.Add(1)` 执行完成 (未开始或正在执行中)，
`wg.Wait()` 自然也就不会产生任何等待， 到此程序结束。

## 正确的做法

通过更改 `wg.Add(1)` 代码的位置，保证循环结束时 5 个 `wg.Add(1)` 均执行完毕。

```go
package main

import (
	"sync"
)

func main() {
	var wg = &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			println("hello world")
		}()
	}

	wg.Wait()
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样 
/**
  hello world
  hello world
  hello world
  hello world
  hello world
*/
```

# 扩展阅读

- [Go 并发编程 - 数据竞态](../engineering/data_race.md) 