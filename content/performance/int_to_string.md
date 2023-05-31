---
title: Go 高性能之整数转字符串
date: 2023-01-01
modify: 2023-01-01
---

# 概述

基础数据类型之间相互转化是开发中常见的功能代码，以 `int` 类型转换为 `string` 类型举例来说，最常用的方法是标准库提供的 `fmt.Sprintf` 和 `strconv.Itoa` 方法，
那么两者之间的性能差异有多大呢？

我们通过基准测试来比较一下。

# 调用 fmt.Sprintf 方法转换 

测试代码如下:

```go
package performance

import (
	"fmt"
	"testing"
)

func Benchmark_IntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", i)
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 10000000 次
$ go test -run='^$' -bench=. -count=1 -benchtime=10000000x . > fmt.txt
```

# 调用 strconv.Itoa 方法转换

测试代码如下:

```go
package performance

import (
	"strconv"
	"testing"
)

func Benchmark_IntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(i)
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000000 次
$ go test -run='^$' -bench=. -count=1 -benchtime=10000000x . > strconv.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 strconv.txt fmt.txt

# 输出如下
name           old time/op  new time/op  delta
_ForeachSet-8  26.3ns ± 0%  85.5ns ± 0%  +224.68%  (p=1.000 n=1+1)
```

从输出的结果中可以看到，`strconv.Itoa` 方法比 `fmt.Sprintf` 方法性能提升了 `2 倍+`。

# 小结

`int` 类型转换为 `string` 类型，直接使用 `strconv.Itoa` 方法。

