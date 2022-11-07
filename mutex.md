# 概述
**对于任一共享资源，同一时间保证只有一个访问者，这种方法称为 `互斥机制`。**

关键字 `Mutex` 表示互斥锁类型，它的 `Lock` 方法用于获取锁，`Unlock` 方法用于释放锁。
在 `Lock` 和 `Unlock` 之间的代码，可以读取和修改共享资源，这部分区域称为 `临界区`。

# 错误的并发操作
先来看一个错误的示例。

在 [Map](map.md) 小节中讲到， **`Map` 不是并发安全的，** 也就是说，如果在多个线程中，同时对一个 Map 进行读写，会报错。
现在来验证一下， 通过启动 `100 个 goroutine` 来模拟并发调用，每个 goroutine 都对 Map 的 key 进行设置。

```go
package main

import "sync"

func main() {
	m := make(map[int]bool)

	var wg sync.WaitGroup

	for j := 0; j < 100; j++ {
		wg.Add(1)

		go func(key int) {
			defer func() {
				wg.Done()
			}()

			m[key] = true	// 对 Map 进行并发写入
		}(j)
	}

	wg.Wait()
}

// $ go run main.go
// 输出如下，报错信息
/**
    fatal error: concurrent map writes
    fatal error: concurrent map writes

    goroutine 104 [running]:
    main.main.func1(0x0?)
            /home/codes/Go-examples-for-beginners/main.go:18 +0x66
    created by main.main
            /home/codes/Go-examples-for-beginners/main.go:13 +0x45
    
    goroutine 1 [semacquire]:
    sync.runtime_Semacquire(0xc0000112c0?)
            /usr/local/go/src/runtime/sema.go:62 +0x25
    sync.(*WaitGroup).Wait(0x60?)
            /usr/local/go/src/sync/waitgroup.go:139 +0x52
    main.main()
            /home/codes/Go-examples-for-beginners/main.go:22 +0x105

    ...
    ...
    ...
*/
```

通过输出信息 `fatal error: concurrent map writes` 可以看到，并发写入 Map 确实会报错。

# 正确的并发操作
Map 并发写入如何正确地实现呢？

一种简单的方案是在并发临界区域 (也就是设置 Map key 的地方) 进行加互斥锁操作， **互斥锁保证了同一时刻
只有一个 goroutine 获得锁**，其他 goroutine 全部处于等待状态，这样就把并发写入变成了串行写入， 
从而消除了报错问题。

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	m := make(map[int]bool)

	var wg sync.WaitGroup

	for j := 0; j < 100; j++ {
		wg.Add(1)

		go func(key int) {
			defer func() {
				wg.Done()
			}()

			mu.Lock()     // 写入前加锁
			m[key] = true // 对 Map 进行并发写入
			mu.Unlock()   // 写入完成解锁
		}(j)
	}

	wg.Wait()

	fmt.Printf("Map size = %d\n", len(m))
}
// $ go run main.go
// 输出如下
/**
    Map size = 100
*/
```

# 扩展阅读
1. [互斥锁 - 维基百科](https://zh.wikipedia.org/wiki/%E4%BA%92%E6%96%A5%E9%94%81)
2. [临界区 - 百度百科](https://baike.baidu.com/item/%E4%B8%B4%E7%95%8C%E5%8C%BA/8942134)