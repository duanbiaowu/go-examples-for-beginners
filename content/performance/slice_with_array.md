---
title: Go 高性能之切片和数组
date: 2023-01-01
modify: 2023-01-01
---

# 概述

> Array or Slice, that's the question!

Go 的数组采用 `值传递` 的方式，直观上看，比采用 `引用传递` 方式的指针要慢，但事实真的是这样吗？

# 使用数组

## 测试代码

```go
package performance

import "testing"

const (
	// 数组容量为 1024
	size = 1024
)

func generate() [size]int {
	res := [size]int{}
	for i := 0; i < size; i++ {
		res[i] = i + 1
	}

	return res
}

func Benchmark_Compare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate()
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x -benchmem > array.txt
```

# 使用切片

## 测试代码

```go
package performance

import "testing"

const (
	size = 1024
)

func generate() []int {
	// 切片容量为 1024
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = i + 1
	}

	return res
}

func Benchmark_Compare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate()
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x -benchmem > slice.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 array.txt slice.txt

# 输出如下
name        old time/op    new time/op    delta
_Compare-8     384ns ± 0%     485ns ± 0%  +26.30%  (p=1.000 n=1+1)

name        old alloc/op   new alloc/op   delta
_Compare-8     0.00B          0.00B          ~     (all equal)

name        old allocs/op  new allocs/op  delta
_Compare-8      0.00           0.00          ~     (all equal)
```

从输出的结果中可以看到，数组的性能略优于切片。感兴趣的读者可以调整常量 `size` 的大小，来比较数组和切片在不同容量下的性能差异。

# 小结

对于容量较小的 `列表型数据` (比如 1 年的 12 个月，1 周的 7 天, 订单状态列表 这种常见的业务数据)，返回数组优于返回切片，
因为**数组的复制成本远远小于切片分配到堆上的成本加上后续 `GC` 的成本**。

对于容量较大的 `内存持久性数据` 或需要 `共享底层数组的切片数据`， 使用切片可以比使用数组获得更好的性能。

