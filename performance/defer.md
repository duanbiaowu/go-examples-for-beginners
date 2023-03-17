# 概述

`defer` 语句保证了不论是在正常情况下 (return 返回)，还是非正常情况下 (发生错误, 程序终止)，函数或方法都能够执行。
**一个完整的 defer 过程要经过函数注册、参数拷⻉、函数提取、函数调用，这要比直接调用函数慢得多**。

# defer 延时释放锁

## 测试代码

```go
package performance

import (
	"sync"
	"testing"
	"time"
)

var (
	m sync.Mutex
)

func foo() {
	m.Lock()
	url := "https://go.dev" // 模拟从队列中获取一个下载 URL
	defer m.Unlock()        // 延迟释放锁

	//http.Get(url)
	_ = url

	time.Sleep(time.Millisecond) // 模拟 HTTP 请求耗时
}

func Benchmark_Compare(b *testing.B) {
	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			foo()
		}()
	}

	wg.Wait()
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000 次, 统计内存分配
$ go test -run='^$'  -bench=. -count=1 -benchtime=1000x -benchmem > slow.txt
```

# 直接释放锁

## 测试代码

```go
package performance

import (
	"sync"
	"testing"
	"time"
)

var (
	m sync.Mutex
)

func foo() {
	m.Lock()
	url := "https://go.dev" // 模拟从队列中获取一个下载 URL
	m.Unlock()              // 直接释放锁

	//http.Get(url)
	_ = url

	time.Sleep(time.Millisecond) // 模拟 HTTP 请求耗时
}

func Benchmark_Compare(b *testing.B) {
	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			foo()
		}()
	}

	wg.Wait()
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000 次, 统计内存分配
$ go test -run='^$'  -bench=. -count=1 -benchtime=1000x -benchmem > fast.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 fast.txt slow.txt 

# 输出如下
name        old time/op    new time/op     delta
_Compare-8    2.75µs ± 0%  1134.99µs ± 0%  +41217.58%  (p=1.000 n=1+1)

name        old alloc/op   new alloc/op    delta
_Compare-8      561B ± 0%       633B ± 0%     +12.83%  (p=1.000 n=1+1)

name        old allocs/op  new allocs/op   delta
_Compare-8      3.00 ± 0%       4.00 ± 0%     +33.33%  (p=1.000 n=1+1)
```

输出的结果分为了三行，分别对应基准测试期间的: 运行时间、内存分配总量、内存分配次数，可以看到:
- 运行时间: `直接释放锁` 比 `defer 释放锁` 提升了 `400 多倍`
- 内存分配总量: `直接释放锁` 比 `defer 释放锁` 降低了 `10% 左右`
- 内存分配次数: `直接释放锁` 比 `defer 释放锁` 降低了 `25%%`

因为时间关系，基准测试只运行了 1000 次，运行次数越大，优化的效果越明显。感兴趣的读者可以将 `-benchtime` 调大后看看优化效果 (值越大，运行时间越长)。

## 性能分析

**使用 `defer 释放锁` 的方案时，`互斥锁` 需要等待 HTTP 请求访问结束，函数退出前调用才能释放，这就导致了 `并发` 的锁争用彻底降级为 `串行` 方式。
这也是为什么使用 `defer 释放锁` 比 `直接释放锁` 的性能低这么多的主要原因**。

# 小结

对于 `资源类` 变量来说，获取并使用完之后，应该尽早地释放。如果代码本就处于 `hot path` 上，应该在 `临界区` 结束之后，立马释放资源，
而不要等到函数返回时才释放。 另外需要注意的一点是尽量不要在循环语句使用 `defer`, 因为这会产生多个 `defer` 语句，导致 `资源类` 释放延迟，性能恶化，
还有可能出现 BUG (参考扩展阅读文章)。
