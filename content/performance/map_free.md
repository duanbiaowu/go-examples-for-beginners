---
title: Go 高性能之 map 重置和删除
date: 2023-01-01
modify: 2023-01-01
---

# 概述

**Map 会自动扩容，但是不会自动缩容**。这也意味着，即使调用 `delete()` 将 `Map` 中的数据删除，内存也不会释放 (为以后的数据备用，类似于预分配的功能)，
随着内存占用越来越多，最终导致性能受到影响。

接下来，我们通过基准测试来对比 `不删除 Map 数据`, `及时删除 Map 无用的数据`, `直接重置 Map 数据` 三者之间的性能差异。

# 不删除 Map 数据

申请一定数量的 `Map`, 然后放入一个切片中，初始化数据之后，做一些逻辑操作 (这里省略)，完成之后并不删除 `Map` 数据。

## 测试代码

```go
package performance

import (
	"strconv"
	"testing"
)

type user struct {
	id       int
	name     string
	password string
	email    string
	token    string
}

func Benchmark_Compare(b *testing.B) {
	ms := make([]map[int]*user, b.N)

	for i := 0; i < b.N; i++ {
		ms[i] = make(map[int]*user, 1024)

		for j := 0; j < 1024; j++ {
			name := strconv.Itoa(j + 1)
			ms[i][j] = &user{
				id:   i + 1,
				name: name,
			}
		}

		// do something
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 10000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=10000x -benchmem > no_delete.txt
```

# 及时删除 Map 无用的数据

申请一定数量的 `Map`, 然后放入一个切片中，初始化数据之后，做一些逻辑操作 (这里省略)，完成之后就删除 `Map` 数据。

## 测试代码

```go
package performance

import (
	"strconv"
	"testing"
)

type user struct {
	id       int
	name     string
	password string
	email    string
	token    string
}

func Benchmark_Compare(b *testing.B) {
	ms := make([]map[int]*user, b.N)

	for i := 0; i < b.N; i++ {
		ms[i] = make(map[int]*user, 1024)

		for j := 0; j < 1024; j++ {
			name := strconv.Itoa(j + 1)
			ms[i][j] = &user{
				id:   i + 1,
				name: name,
			}
		}

		// do something

		// 完成之后删除数据
		for j := 0; j < 1024; j++ {
			delete(ms[i], j) // 删除 Map 数据
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 10000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=10000x -benchmem > delete.txt
```

# 直接重置 Map 数据

申请一定数量的 `Map`, 然后放入一个切片中，初始化数据之后，做一些逻辑操作 (这里省略)，完成之后重置 `Map`。

## 测试代码

```go
package performance

import (
	"strconv"
	"testing"
)

type user struct {
	id       int
	name     string
	password string
	email    string
	token    string
}

func Benchmark_Compare(b *testing.B) {
	ms := make([]map[int]*user, b.N)

	for i := 0; i < b.N; i++ {
		ms[i] = make(map[int]*user, 1024)

		for j := 0; j < 1024; j++ {
			name := strconv.Itoa(j + 1)
			ms[i][j] = &user{
				id:   i + 1,
				name: name,
			}
		}

		ms[i] = nil // 重置 Map, 及时释放资源
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 10000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=10000x -benchmem > reset.txt
```

# 使用 benchstat 比较差异

## 不删除数据 VS 及时删除无用数据

```shell
$ benchstat -alpha=100 no_delete.txt delete.txt 

# 输出如下
name        old time/op    new time/op    delta
_Compare-8     198µs ± 0%     173µs ± 0%  -12.42%  (p=1.000 n=1+1)

name        old alloc/op   new alloc/op   delta
_Compare-8     126kB ± 0%     126kB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
_Compare-8     1.95k ± 0%     1.95k ± 0%     ~     (all equal)
```

从输出的结果中可以看到，`及时删除数据` 比 `不删除数据` 性能提升了 `12%`, 内存方面没有差别。

## 及时删除无用数据 VS 重置 Map

```shell
$ benchstat -alpha=100 delete.txt reset.txt

# 输出如下
name        old time/op    new time/op    delta
_Compare-8     173µs ± 0%     106µs ± 0%  -38.63%  (p=1.000 n=1+1)

name        old alloc/op   new alloc/op   delta
_Compare-8     126kB ± 0%     126kB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
_Compare-8     1.95k ± 0%     1.95k ± 0%     ~     (all equal)
```

从输出的结果中可以看到，`重置 Map` 比 `及时删除无用数据` 性能提升了 `38%`, 内存方面没有差别。

## 不删除数据 VS 重置 Map

```shell
$ benchstat -alpha=100 no_delete.txt reset.txt

# 输出如下
name        old time/op    new time/op    delta
_Compare-8     198µs ± 0%     106µs ± 0%  -46.25%  (p=1.000 n=1+1)

name        old alloc/op   new alloc/op   delta
_Compare-8     126kB ± 0%     126kB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
_Compare-8     1.95k ± 0%     1.95k ± 0%     ~     (all equal)
```

从输出的结果中可以看到，`重置 Map` 比 `不删除数据` 性能提升了 `46%`, 内存方面没有差别。

# 小结

`Map 不会自动缩容` 这个特性决定了在必要的情况下需要 `手动优化`，一些比较通用的实践是:

- 栈上分配的容量小的 `Map` 无需优化
- 栈上分配的容量大的 `Map` 使用完成之后
    - 如果后面几乎没有代码，无需优化
    - 如果后续还有大量逻辑代码，应立即重置释放 `Map`
- 堆上分配 `Map`
    - 如果是初始化时，内存就已预分配完成的 `缓存类应用`，定时处理数据稀疏的 `Map`
    - 业务逻辑代码
        - 及时删除过期数据
        - 定时重置数据稀疏的 `Map`

对象到底分配到堆上还是栈上？请参考扩展阅读 - 逃逸分析相关文章。
