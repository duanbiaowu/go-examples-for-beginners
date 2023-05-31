---
title: Go 高性能之切片预分配
date: 2023-01-01
modify: 2023-01-01
---

# 概述

`切片` 追加元素时，直接调用 `append` 函数即可，开发者不需要考虑 `切片` 容量不足问题，因为 `append` 函数内部已经实现了 `自动扩容机制`，
从开发者的角度看，这大大提高了生产力并降低了心智负担。

但是, **软件工程没有银弹**，开发便利性的背后必然是以函数内部实现的复杂性为代价的。如果我们使用 `预分配机制`，在 `切片` 初始化的时候就定义好容量，
那么就可以规避 `append` 函数内部触发 `自动扩容`，从而提高程序的性能。

接下来，我们通过基准测试来比较一下 `append` 函数的 `自动扩容机制` 和 `预分配机制` 的性能差异。

# append 自动扩容机制

测试代码如下:

```go
package performance

import (
	"testing"
)

func Benchmark_Slice(b *testing.B) {
	size := 10000

	for n := 0; n < b.N; n++ {
		data := make([]int, 0) // 没有预先分配容量
		for k := 0; k < size; k++ {
			// 容量不足时，append 函数内部会自动扩容
			data = append(data, k)
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

import (
	"testing"
)

func Benchmark_Slice(b *testing.B) {
	size := 10000

	for n := 0; n < b.N; n++ {
		data := make([]int, 0, size) // 预先分配了容量
		for k := 0; k < size; k++ {
			// append 函数内部不会自动扩容
			data = append(data, k)
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x > fast.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 fast.txt slow.txt

# 输出如下 
name      old time/op  new time/op  delta
_Slice-8  17.3µs ± 0%  74.0µs ± 0%  +328.43%  (p=1.000 n=1+1)
```

从输出的结果中可以看到，`预分配机制` 比 `append` 函数的 `自动扩容机制` 性能高出 `3 倍+`。

感兴趣的读者可以将 `切片` 容量调大一些，观察性能提升的巨大差异。

# 小结

- 调用 `append` 函数前初始化 `切片` 容量 
- 初始化 `切片` 容量时，尽可能设置到足够使用，避免扩容 
- 当 `切片` 的容量越大，`预分配机制` 带来的性能提升越明显

