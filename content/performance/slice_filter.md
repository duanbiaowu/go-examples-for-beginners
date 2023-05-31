---
title: Go 高性能之切片过滤器
date: 2023-01-01
modify: 2023-01-01
---

# 概述

切片的底层是数组，并且不同的切片之间共享一个底层数组，在实现 `过滤器` 功能时，可以利用这个特点，**将过滤后的结果切片引用为同一个底层数组，实现内存零分配**。

# 不复用底层数组

测试代码如下:

```go
package performance

import "testing"

func filter(x int) bool {
	return x&1 == 1
}

func Benchmark_Filter(b *testing.B) {
	b.StopTimer()

	// 数据初始化操作
	size := 10000
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}

	b.StartTimer()

	for n := 0; n < b.N; n++ {
		res := make([]int, 0, size>>1) // res 重新初始化，不复用 data 的底层数组

		for i := 0; i < size; i++ {
			if filter(data[i]) {
				res = append(res, data[i])
			}
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次，统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x -benchmem . > slow.txt
```

# 复用底层数组

测试代码如下:

```go
package performance

import "testing"

func filter(x int) bool {
	return x&1 == 1
}

func Benchmark_Filter(b *testing.B) {
	b.StopTimer()

	// 数据初始化操作
	size := 10000
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}

	b.StartTimer()

	for n := 0; n < b.N; n++ {
		res := data[:0] // res 复用 data 的底层数组

		for i := 0; i < size; i++ {
			if filter(data[i]) {
				res = append(res, data[i])
			}
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次，统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x -benchmem . > fast.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 fast.txt slow.txt 

# 输出如下
name       old time/op    new time/op     delta
_Filter-8    7.92µs ± 0%    21.11µs ± 0%  +166.67%  (p=1.000 n=1+1)

name       old alloc/op   new alloc/op    delta
_Filter-8     0.00B       40960.00B ± 0%     +Inf%  (p=1.000 n=1+1)

name       old allocs/op  new allocs/op   delta
_Filter-8      0.00            1.00 ± 0%     +Inf%  (p=1.000 n=1+1)
```

输出的结果分为了三行，分别对应基准期间的: 运行时间、内存分配总量、内存分配次数，采用了 `复用底层数组` 方案后:
- 运行时间提升了将近 `1.7 倍`
- 内存分配总量降至 0
- 内存分配次数降至 0

# 小结

通过复用底层数组，极大地提高了性能，最重要的是，将内存分配优化到了 0。当然，**这个优化的前提是: 源切片指向的底层数组的共享数据允许被修改**。
