---
title: Go 高性能之获取调用堆栈优化
date: 2023-01-01
modify: 2023-01-01
---

# 概述

在工程代码中需要在异常场景打印相应的日志，记录重要的上下文信息。如果遇到 `panic` 或 `error` 的情况，
这时候就需要详细的 `堆栈信息` 作为辅助来排查问题，本小节就来介绍两种常见的获取 `堆栈信息` 方法，
然后对两种方法进行基准测试，最后使用测试的结果进行性能对比并分析差异。

# runtime.Stack

通过标准库提供的 `runtime.Stack` 相关 API 来获取。

## 示例

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, true)

	fmt.Printf("%s\n", buf[:n])
}
```

```shell
$ go run main.go

# 输出如下 (你的输出代码路径应该和这里的不一样)
goroutine 1 [running]:
main.main()
    /home/codes/go-high-performance/main.go:10 +0x45
...

```

## 测试代码如下

```go
package performance

import (
	"runtime"
	"testing"
)

func Benchmark_StackDump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 1024)
		n := runtime.Stack(buf, true)

		_ = buf[:n]
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000 次，统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=1000x -benchmem > slow.txt
```

# runtime.Caller

通过标准库提供的 `runtime.Caller` 相关 API 来获取。

## 示例

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; ; i++ {
		if _, file, line, ok := runtime.Caller(i); ok {
			fmt.Printf("file: %s, line: %d\n", file, line)
		} else {
			break
		}
	}
}
```

```shell
$ go run main.go

# 输出如下 (你的输出代码路径应该和这里的不一样)
file: /home/codes/go-high-performance/main.go, line: 10
file: /usr/local/go/src/runtime/proc.go, line: 250
file: /usr/local/go/src/runtime/asm_amd64.s, line: 1594
...
```

从输出的结果中可以看到，`runtime.Caller` 的返回值包含了 `文件名称` 和 `行号`，但是相比 `runtime.Stack` 的输出而言，
缺少了 `goroutine` 和 `调用方法` 字段，我们可以通过 `runtime.Callers` 配合 `runtime.CallersFrames` 输出和 `runtime.Stack` 一样的结果。

```go
package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	pcs := make([]uintptr, 16)
	n := runtime.Callers(0, pcs)

	frames := runtime.CallersFrames(pcs[:n])

	var sb strings.Builder
	for {
		frame, more := frames.Next()

		sb.WriteString(frame.Function)
		sb.WriteByte('\n')
		sb.WriteByte('\t')
		sb.WriteString(frame.File)
		sb.WriteByte(':')
		sb.WriteString(strconv.Itoa(frame.Line))
		sb.WriteByte('\n')

		if !more {
			break
		}
	}

	fmt.Println(sb.String())
}
```

```shell
$ go run main.go

# 输出如下 (你的输出代码路径应该和这里的不一样)
runtime.Callers
        /usr/local/go/src/runtime/extern.go:247
main.main
        /home/codes/go-high-performance/main.go:12
runtime.main
        /usr/local/go/src/runtime/proc.go:250
runtime.goexit
        /usr/local/go/src/runtime/asm_amd64.s:1594
...
```

## 测试代码

```go
package performance

import (
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func stackDump() string {
	pcs := make([]uintptr, 16)
	n := runtime.Callers(0, pcs)

	frames := runtime.CallersFrames(pcs[:n])

	var buffer strings.Builder
	for {
		frame, more := frames.Next()

		buffer.WriteString(frame.Function)
		buffer.WriteByte('\n')
		buffer.WriteByte('\t')
		buffer.WriteString(frame.File)
		buffer.WriteByte(':')
		buffer.WriteString(strconv.Itoa(frame.Line))
		buffer.WriteByte('\n')

		if !more {
			break
		}
	}

	return buffer.String()
}

func Benchmark_StackDump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stackDump()
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000 次，统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=1000x -benchmem > fast.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 fast.txt slow.txt 

# 输出如下
name          old time/op    new time/op    delta
_StackDump-8    2.28µs ± 0%   68.89µs ± 0%  +2926.85%  (p=1.000 n=1+1)

name          old alloc/op   new alloc/op   delta
_StackDump-8    1.36kB ± 0%    1.02kB ± 0%    -24.71%  (p=1.000 n=1+1)

name          old allocs/op  new allocs/op  delta
_StackDump-8      12.0 ± 0%       1.0 ± 0%    -91.67%  (p=1.000 n=1+1)
```

输出的结果分为了三行，分别对应基准测试期间的: 运行时间、内存分配总量、内存分配次数，可以看到:

- 运行时间: `runtime.Callers` 比 `runtime.Stack` 提升了将近 `30 倍`
- 内存分配总量: 两者差不多
- 内存分配次数: `runtime.Callers` 比 `runtime.Stack` 降低了将近 `10 倍`，当然笔者的测试代码也需要再优化下

## 性能分析

> 最根本的差异点在于 `runtime.Stack` 会触发 `STW` 操作。

# 小结

本小节介绍了两种获取堆栈信息的方法，并通过基准测试来分析两种方法的性能差异，读者可以在此基础上封装自己的高性能组件类库。
