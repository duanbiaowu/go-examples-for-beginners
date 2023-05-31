---
title: Go 高性能之 channel 缓冲和非缓冲
date: 2023-01-01
modify: 2023-01-01
---

# 概述

缓冲通道还是无缓冲通道，在高性能场景下，如何选择？

# 无缓冲通道测试代码

实现功能如下: 初始化一个 `无缓冲通道`，启动 N 个 `goroutine` 向通道写入数据，然后在 `主 goroutine` 读取通道数据，数据全部读取完成后关闭通道。

```go
package performance

import (
	"sync"
	"testing"
)

func Benchmark_Compare(b *testing.B) {
	var wg sync.WaitGroup
	ch := make(chan struct{})

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-ch
		}()
	}

	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}

	wg.Wait()
	close(ch)
}
```

运行测试，并将基准测试结果写入文件:

```shell
$ go test -run='^$' -bench=. -count=1 -benchmem > nobuffer.txt 
```

# 缓冲通道测试代码

实现功能如下: 初始化一个 `缓冲通道` (容量为 1024)，启动 N 个 `goroutine` 向通道写入数据，然后在 `主 goroutine` 读取通道数据，数据全部读取完成后关闭通道。

```go
package performance

import (
	"sync"
	"testing"
)

func Benchmark_Compare(b *testing.B) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1024)

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-ch
		}()
	}
    
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}

	wg.Wait()
	close(ch)
}
```

运行测试，并将基准测试结果写入文件:

```shell
$ go test -run='^$' -bench=. -count=1 -benchmem > buffer.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 nobuffer.txt buffer.txt

# 输出如下
name        old time/op    new time/op    delta
_Compare-8    2.29µs ± 0%    1.94µs ± 0%  -15.14%  (p=1.000 n=1+1)

name        old alloc/op   new alloc/op   delta
_Compare-8      573B ± 0%      577B ± 0%   +0.70%  (p=1.000 n=1+1)

name        old allocs/op  new allocs/op  delta
_Compare-8      2.00 ± 0%      2.00 ± 0%     ~     (all equal)
```

从输出的结果中可以看到，通过 `给通道设置缓冲容量`, 性能可以优化 `15%` 左右。

# 小结

使用通道时，除了考虑 `性能` 因素外，还需要考虑场景语义，**一个适用于大部分场景的万金油规则是: 设置通道缓冲容量为 1**。
