---
title: Go 高性能之 map 预分配
date: 2023-01-01
modify: 2023-01-01
---

# 概述

`map` 可以直接设置元素，如果对应的 `key` 不存在，内部运行时会生成一个新的 `key`，开发者不需要考虑 `map` 容量不足问题，因为内部运行时已经实现了 `自动扩容机制`，
从开发者的角度看，这大大提高了生产力并降低了心智负担。

但是, **软件工程没有银弹**，开发便利性的背后必然是以函数内部实现的复杂性为代价的。如果我们使用 `预分配机制`，在 `map` 初始化的时候就定义好容量，
那么就可以规避内部运行时触发 `自动扩容`，从而提高程序的性能。

接下来，我们通过基准测试来比较一下内部运行时的 `自动扩容机制` 和 `预分配机制` 的性能差异。

# 自动扩容机制

测试代码如下:

```go
package performance

import "testing"

func Benchmark_Map(b *testing.B) {
	size := 10000

	for n := 0; n < b.N; n++ {
		data := make(map[int]int) // 没有预先分配容量
		for k := 0; k < size; k++ {
			// 容量不足时会发生自动扩容
			data[k] = k
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x > slow.txt
```

# 预分配机制

测试代码如下:

```go
package performance

import "testing"

func Benchmark_Map(b *testing.B) {
	size := 10000

	for n := 0; n < b.N; n++ {
		data := make(map[int]int, size) // 预先分配了容量
		for k := 0; k < size; k++ {
			// 不会发生自动扩容
			data[k] = k
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
$ go test -run='^$' -bench=. -count=1 -benchtime=10000x > fast.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 fast.txt slow.txt 

# 输出如下
name    old time/op  new time/op  delta
_Map-8   306µs ± 0%   687µs ± 0%  +124.15%  (p=1.000 n=1+1)
```

从输出的结果中可以看到，`预分配机制` 比 `自动扩容机制` 性能高出 `1 倍+`。

感兴趣的读者可以将 `map` 容量调大一些，观察性能提升的巨大差异。

# 小结

- 设置 `map` 的 `key` 之前初始化 `map` 容量
- 初始化 `map` 容量时，尽可能设置到足够使用，避免扩容 
- 当 `map` 的容量越大，`预分配机制` 带来的性能提升越明显
