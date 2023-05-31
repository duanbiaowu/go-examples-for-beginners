---
title: Go 高性能之互斥锁和读写锁
date: 2023-01-01
modify: 2023-01-01
---

# 概述

标准库 `sync` 提供了 2 种锁，`sync.Mutex (互斥锁)` 和 `sync.RWMutex (读写锁)`。

### 互斥锁

简单来说，`互斥锁` 可以保证同一临界区的代码，在同一时刻只有一个线程可以执行 (更多理论知识可以参考附录 1)，`sync.Mutex` 提供了
2 个方法:

- Lock: 获取锁
- Unlock: 释放锁

`Lock` 方法是一个阻塞操作，并发线程中一旦有一个线程获得锁，那么其他线程陷入阻塞等待，直至该线程调用 `Unlock` 方法释放锁。

### 读写锁

简单来说，`读写锁` 也称 `共享 - 互斥锁`，读操作是并发可重入的，也就是说多个线程可以并发执行临界区代码，写操作是互斥的，
规则同 `互斥锁` 一致，`sync.RWMutex` 提供了 4 个方法:

- Lock: 获取写锁
- Unlock: 释放写锁
- RLock: 获取读锁
- RUnlock: 释放读锁

# 测试场景

有了基本了解后，接下来通过基准测试，看看在不同场景下，两者之间的性能差异是多少，这里模拟 `常见的 3 种场景`:

- 读多写少 (读占 90%, 写占 10%)
- 写多读少 (写占 10%, 写占 90%)
- 读写一致 (读写各占 50%)

# 测试代码

```go
package performance

import (
	"sync"
	"testing"
	"time"
)

const cost = time.Microsecond // 模拟操作耗时

// 读写接口
type RW interface {
	Write()
	Read()
}

// 互斥锁实现读写接口
type Lock struct {
	count int
	mu    sync.Mutex
}

// 互斥锁实现 Write 方法
func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost) // 模拟操作耗时
	l.mu.Unlock()
}

// 互斥锁实现 Read 方法
func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost) // 模拟操作耗时
	_ = l.count
	l.mu.TryLock()
	l.mu.Unlock()
}

// 读写锁实现读写接口
type RWLock struct {
	count int
	mu    sync.RWMutex
}

// 读写锁实现 Write 方法
func (l *RWLock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost) // 模拟操作耗时
	l.mu.Unlock()
}

// 读写锁实现 Read 方法
func (l *RWLock) Read() {
	l.mu.RLock()
	_ = l.count
	time.Sleep(cost) // 模拟操作耗时
	l.mu.RUnlock()
}

// 基准测试
func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// 读写比例 9:1
func BenchmarkReadMore(b *testing.B)   { benchmark(b, &Lock{}, 9, 1) }
func BenchmarkReadMoreRW(b *testing.B) { benchmark(b, &RWLock{}, 9, 1) }

// 读写比例 1:9
func BenchmarkWriteMore(b *testing.B)   { benchmark(b, &Lock{}, 1, 9) }
func BenchmarkWriteMoreRW(b *testing.B) { benchmark(b, &RWLock{}, 1, 9) }

// 读写比例 5:5
func BenchmarkEqual(b *testing.B)   { benchmark(b, &Lock{}, 5, 5) }
func BenchmarkEqualRW(b *testing.B) { benchmark(b, &RWLock{}, 5, 5) }
```

运行测试:

```shell
$ go test -run='^$' -bench=. -count=1 -benchmem

# 输出结果如下
BenchmarkReadMore-8                   19          63654389 ns/op          124577 B/op       2064 allocs/op
BenchmarkReadMoreRW-8                157           7996424 ns/op          112528 B/op       2006 allocs/op
BenchmarkWriteMore-8                  18          69739556 ns/op          116934 B/op       2052 allocs/op
BenchmarkWriteMoreRW-8                18          66364517 ns/op          115617 B/op       2038 allocs/op
BenchmarkEqual-8                      16          67880962 ns/op          117561 B/op       2058 allocs/op
BenchmarkEqualRW-8                    33          36549494 ns/op          113765 B/op       2019 allocs/op
```

从输出的结果中可以看到：

- 读写比为 9 : 1 时，读写锁是互斥锁的 8 倍+
- 读写比为 1 : 9 时，读写锁和互斥锁基本持平
- 读写比为 5 : 5 时，读写锁是互斥锁的 2 倍

当然，上述测试代码过于简单，并不能充分地说明 `互斥锁` 和 `读写锁` 真正的差异，实际开发中的场景更加复杂、影响的因素也更多，
需要在严格的测试基础上选定适合的方案。

# 小结

- 写比例远大于读比例时，使用 `sync.Mutex`
- 其他情况，使用 `sync.RWMutex`
- 根据具体场景以基准测试结果为准

# 扩展阅读

- [互斥锁 - 维基百科](https://zh.wikipedia.org/wiki/%E4%BA%92%E6%96%A5%E9%94%81)
- [读写锁 - 维基百科](https://zh.wikipedia.org/wiki/%E8%AF%BB%E5%86%99%E9%94%81)