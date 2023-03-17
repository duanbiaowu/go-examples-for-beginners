# 概述

Go 语言刻意没有提供获取 `goroutine ID` 的原因是为了避免滥用。因为大部分用户在轻松拿到 `goroutine ID` 之后，
在之后的编程中会不自觉地编写出强依赖 `goroutine ID` 的代码。

下面介绍两种获取 `goroutine ID` 的方法，一种是通过标准库中的堆栈相关方法获取，一种是通过第三方库 (汇编实现) 获取。

# 通过堆栈调用获取

## 测试代码

```go
package performance

import (
	"bytes"
	"errors"
	"runtime"
	"strconv"
	"testing"
)

func getGoroutineId() (int64, error) {
	// 堆栈结果中需要消除的前缀符
	var goroutineSpace = []byte("goroutine ")

	bs := make([]byte, 128)
	bs = bs[:runtime.Stack(bs, false)]
	bs = bytes.TrimPrefix(bs, goroutineSpace)
	i := bytes.IndexByte(bs, ' ')
	if i < 0 {
		return -1, errors.New("get current goroutine id failed")
	}
	return strconv.ParseInt(string(bs[:i]), 10, 64)
}

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = getGoroutineId()
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000 次，统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=1000x -benchmem > slow.txt
```

# 通过汇编获取

`goid` 是一个开源库，用来获取当前 `goroutine` 的 ID，直接使用汇编实现，性能很高。

## 安装 goid

```shell
$ go get -u github.com/petermattis/goid

# go: downloading github.com/petermattis/goid v0.0.0-20221018141743-354ef7f2fd21
# go: added github.com/petermattis/goid v0.0.0-20221018141743-354ef7f2fd21
```

## 测试代码

```go
package performance

import (
	"testing"

	"github.com/petermattis/goid"
)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		goid.Get()
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

# 输出如下:
name              old time/op    new time/op     delta
BufferWithPool-8    3.04ns ± 0%  5988.00ns ± 0%  +196808.91%  (p=1.000 n=1+1)

name              old alloc/op   new alloc/op    delta
BufferWithPool-8     0.00B         130.00B ± 0%        +Inf%  (p=1.000 n=1+1)

name              old allocs/op  new allocs/op   delta
BufferWithPool-8      0.00            2.00 ± 0%        +Inf%  (p=1.000 n=1+1)
```

**汇编完胜，绝对性降维打击**。

# 小结

> 理解 `goroutine ID` 的获取方法后，请忘记它们。

## 为什么不应该使用 goroutine ID

1. https://groups.google.com/g/golang-nuts/c/Nt0hVV_nqHE
2. https://groups.google.com/g/golang-nuts/c/0HGyCOrhuuI
3. https://stackoverflow.com/questions/19115273/looking-for-a-call-or-thread-id-to-use-for-logging

