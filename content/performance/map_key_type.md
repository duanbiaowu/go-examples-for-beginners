---
title: Go 高性能之 map key 类型
date: 2023-01-01
modify: 2023-01-01
---

# 概述

Map 的 `key` 支持很多数据类型，只要满足 [比较规则](../introduction/type_comparison.md) 即可，
大多数场景下，我们使用到的是 `int` 和 `string` 两种数据类型，那么两者之间，哪个性能更高一些呢？

# key 类型为 string

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
	for i := 0; i < b.N; i++ {
		// string 作为 key
		m := make(map[string]*user, 1024)
		name := strconv.Itoa(i + 1)
		m[name] = &user{
			id:   i + 1,
			name: name,
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x -benchmem > string.txt
```

# key 类型为 int

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
	for i := 0; i < b.N; i++ {
		// int 作为 key
		m := make(map[int]*user, 1024)
		name := strconv.Itoa(i + 1)
		m[i] = &user{
			id:   i + 1,
			name: name,
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x -benchmem > int.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 int.txt string.txt 

# 输出如下
name        old time/op    new time/op    delta
_Compare-8    8.35µs ± 0%   11.73µs ± 0%  +40.41%  (p=1.000 n=1+1)

name        old alloc/op   new alloc/op   delta
_Compare-8    41.1kB ± 0%    57.5kB ± 0%  +39.89%  (p=1.000 n=1+1)

name        old allocs/op  new allocs/op  delta
_Compare-8      4.00 ± 0%      4.00 ± 0%     ~     (all equal)
```

从输出的结果中可以看到，使用 `int` 作为 `key` 相较于使用 `string` 作为 `key`, 性能和内存使用能优化 `40%` 左右。

# 小结

Map 的 `key` 应尽量使用 `int` 类型，但是如果存储的对象集合元素没有唯一的 `int` 类型字段，可以考虑折中的方法: 
将元素 `Hash` 为一个 `int` 数字，当然， `Hash` 之后记得做基准测试。
