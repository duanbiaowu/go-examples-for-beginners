# 概述

调用 `sync/atomic` 包即可。

# 错误的并发操作

先来看一个错误的示例。

通过启动 `1000 个 goroutine` 来模拟并发调用，在函数内部对变量 `number` 进行自增操作，
那么可能存在的一个问题是，当多个 goroutine 同时对变量操作时，只有一个成功了，其他的全部失败，造成的后果就是变量
最终的值小于 1000 (正常情况应该是等于 1000)。

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var number uint32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer func() {
				wg.Done()
			}()

			number++
		}()
	}

	wg.Wait()

	fmt.Printf("number = %d\n", number)
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样，多试几次，会发现每次的结果都不一样
/**
  number = 971
*/
```

# 正确的并发操作

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var number uint32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer func() {
				wg.Done()
			}()

			//number++
			atomic.AddUint32(&number, 1) // 使用原子操作
		}()
	}

	wg.Wait()

	fmt.Printf("number = %d\n", number)
}

// $ go run main.go
// 输出如下，多试几次，会发现结果都是一样的
/**
  number = 1000
*/
```

# 附录

## 使用 atomic 包对 uint32 类型实现 -1 功能

```go
var value uint32
newValue := atomic.AddUint32(&value, 1)
t.Logf("new = %d", newValue)

atomic.AddUint32(&value, ^uint32(0))   // ^uint32(0) == -1
newValue = atomic.LoadUint32(&value)
t.Logf("new = %d", newValue)
```
