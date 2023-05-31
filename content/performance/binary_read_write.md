---
title: Go 高性能之字节序优化
date: 2023-01-01
modify: 2023-01-01
---

# 概述

> encoding/binary 包用于数字和字节序列之间的简单转换以及 varints 的编码和解码。

`varints` 是一种使用可变字节表示整数的方法，其中数值本身越小，其所占用的字节数越少。

标准库中的 `binary.Read` 方法和 `binary.Write` 方法内部使用 `反射` 实现，会对性能有一定影响。如果相关代码在 `hot path` 上面，
那么应该考虑是否可以手动实现相关功能，避免直接使用这两个函数。

# 直接使用 binary.read

## 测试代码

```go
package performance

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// 将网络字节序解析到 uint32
func convert(bys []byte) uint32 {
	var num uint32
	buf := bytes.NewReader(bys)
	_ = binary.Read(buf, binary.BigEndian, &num)
	return num
}

func Benchmark_Convert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = convert([]byte{0x7f, 0, 0, 0x1})
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x -benchmem > slow.txt
```

# 手动实现函数

## 测试代码

```go
package performance

import (
	"testing"
)

// 将网络字节序解析到 uint32
func convert(bys []byte) uint32 {
	return uint32(bys[3]) | uint32(bys[2])<<8 | uint32(bys[1])<<16 | uint32(bys[0])<<24
}

func Benchmark_Convert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = convert([]byte{0x7f, 0, 0, 0x1})
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x -benchmem > fast.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 fast.txt slow.txt

# 输出如下
name          old time/op    new time/op    delta
_Convert-8    0.26ns ± 0%   83.85ns ± 0%  +31782.13%  (p=1.000 n=1+1)

name          old alloc/op   new alloc/op   delta
_Convert-8     0.00B         60.00B ± 0%       +Inf%  (p=1.000 n=1+1)

name          old allocs/op  new allocs/op  delta
_Convert-8      0.00           4.00 ± 0%       +Inf%  (p=1.000 n=1+1)
```

从输出的结果中可以看到，通过手动实现函数功能的方式，避免了 `反射` 的开销，运行时间提升了 `300 多倍`, 内存分配量和内存分配次数降为 `0`。

# 小结

> 任何出现在 `hot path` 上面的反射代码都应该尽可能地优化。
