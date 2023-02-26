# 概述

`数据竞态` 是并发系统编程中最常见和最难调试的错误类型之一。当两个 `goroutine` 同时访问同一个变量并且至少有一个是写入操作时，
就会发生 `数据竞态`。详细的信息，请参阅 [Go 内存模型](https://go.dev/ref/mem/)。

# 使用方法

为了帮助诊断 `数据竞态` 导致的错误，Go 内置了一个 `数据竞态检测器`，只需要在具体的命令中加上 `-race` 参数即可。

```shell
$ go test -race mypkg    // 测试单个包
$ go run -race mysrc.go  // 运行单个文件
$ go build -race mycmd   // 构建命令
$ go install -race mypkg // 安装包
```

# 报告格式

当 `数据竞态检测器` 在程序中发现 `数据竞态` 时，会输出一份报告。该报告包含访问冲突的调用堆栈，以及相关 `goroutine` 的创建堆栈。

```shell
# 示例输出
WARNING: DATA RACE
Read by goroutine 185:
  net.(*pollServer).AddFD()
      src/net/fd_unix.go:89 +0x398
  net.(*pollServer).WaitWrite()
      src/net/fd_unix.go:247 +0x45
  net.(*netFD).Write()
      src/net/fd_unix.go:540 +0x4d4
  net.(*conn).Write()
      src/net/net.go:129 +0x101
  net.func·060()
      src/net/timeout_test.go:603 +0xaf
...
...
```

# 可选项

Go 的环境变量 `GORACE` 对应 `数据竞态检测器` 的参数设置，格式如下:

```shell
GORACE="option1=val1 option2=val2"
```

其中，`options` 参数的可选项为:

- log_path (default stderr): `数据竞态检测器` 将报告写入名为 `log_path.pid` 的文件中。`stdout` 和 `stderr` 则写入标准输出和标准错误
- exitcode (default 66): `数据竞态检测器` 检测到竞态后退出时的错误码
- strip_path_prefix (default ""): 从报告中删除指定的前缀符
- history_size (default 1):  每个 `goroutine` 的历史内存访问容量是 32K * 2 **history_size 个元素。增加这个值可以避免报告中出现 "无法恢复堆栈" 错误，代价是增加内存使用量**
- halt_on_error (default 0): 控制程序是否在第一次 `数据竞态` 报告后退出
- atexit_sleep_ms (default 1000): 程序退出前在主 `goroutine` 休眠的毫秒数

```shell
# 示例
GORACE="log_path=/tmp/race/report strip_path_prefix=/my/go/sources/" go test -race
```

# 排除测试

当在命令中加入 `-race` 参数时，Go 定义了额外的构建标签 `race`，这样就可以使用标签在运行 `数据竞态检测器` 时排除一些代码和测试。
从本质上来说，其实也就是 [条件编译](conditional_compilation.md)。

## 示例

有了 `条件编译` 之后，运行 `go test -race` 命令时，这个文件里面的单元测试就会被排除，不会运行。

```go
// +build !race

package foo

// The test contains a data race. See issue 123.
func TestFoo(t *testing.T) {
	// ...
}

// The test fails under the race detector due to timeouts.
func TestBar(t *testing.T) {
	// ...
}

// The test takes too long under the race detector.
func TestBaz(t *testing.T) {
	// ...
}
```

# 典型的数据竞争

下面是一些典型的数据竞争的小例子，所有的例子都可以使用 `数据竞态检测器` 检测到。

## 循环里面的数据竞争

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // Not the 'i' you are looking for.
			wg.Done()
		}()
	}
	wg.Wait()
}
```

```shell
$ go run -race main.go
# 输出如下 
5
5
5
==================
WARNING: DATA RACE
Read at 0x00c0000180f8 by goroutine 10:
  main.main.func1()
...
...
...
==================
5
5
Found 1 data race(s)
exit status 66
```

**错误原因在于:** 函数字面量中的 `变量 i` 与循环使用的变量相同，因此 `goroutine` 中的 `读取与循环变量自增竞争`，可以通过复制变量来修复该程序。

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Println(j) // Good. Read local copy of the loop counter.
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

```shell
$ go run -race main.go
# 输出如下，你的输出可能和这里的不一样
4
1
0
2
3
```

## 不受保护的全局变量

如果从多个 `goroutine` 调用以下代码，则会导致 `service map` 的竞态。并发读写一个 map 是不安全的，解决方案可以使用 `互斥锁`。

```go
package main

import (
	"sync"
)

var service map[string]int

func RegisterService(name string, addr int) {
	// 并发写 map
	service[name] = addr
}

func LookupService(name string) int {
	return service[name]
}

func main() {
	service = make(map[string]int)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			RegisterService("zero", n)
			LookupService("zero")
		}(i)
	}

	wg.Done()
}
```

```shell
$ go run -race main.go
# 输出如下
==================
WARNING: DATA RACE
Write at 0x00c00007e000 by goroutine 6:
  runtime.mapassign_faststr()
...
...
...
Found 2 data race(s)
exit status 66
```

**解决方案:** 通过增加 `互斥锁` 来消除 `数据竞态`，修正后的代码如下：

```go
package main

import (
	"sync"
)

var (
	service   map[string]int
	serviceMu sync.Mutex
)

func RegisterService(name string, addr int) {
	// 并发写 map 之前，加锁
	serviceMu.Lock()
	defer serviceMu.Unlock()
	service[name] = addr
}

func LookupService(name string) int {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	return service[name]
}

func main() {
	service = make(map[string]int)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			RegisterService("zero", n)
			LookupService("zero")
		}(i)
	}

	wg.Done()
}
```

```shell
$ go run -race main.go
# 没有任何输出
```

从输出中可以看到，`数据竞态` 错误已经消除。

## 通道的发送和关闭操作未同步

同一个通道上面，发送和关闭操作如果不同步，也可能产生 `数据竞态`。

```go
package main

func main() {
	c := make(chan struct{})
	go func() { c <- struct{}{} }()
	close(c)
}
```

```shell
$ go run -race main.go
# 输出如下
==================
WARNING: DATA RACE
Read at 0x00c000110070 by goroutine 6:
  runtime.chansend()
...
...
...
Found 1 data race(s)
exit status 66
```

**错误原因在于: 根据 Go 内存模型，通道上面的发送操作发生在该通道相应的接收完成之前**。

要同步发送和关闭操作，请使用 **接收操作，可以保证发送操作在关闭之前完成**。

```go
package main

func main() {
	c := make(chan struct{})
	go func() { c <- struct{}{} }()
	<-c
	close(c)
}
```

```shell
$ go run -race main.go
# 没有任何输出
```

从输出中可以看到，`数据竞态` 错误已经消除。

# 必要条件

`数据竞态检测器` 需要启用 cgo 并且支持 linux/amd64, linux/ppc64le, linux/arm64, freebsd/amd64, netbsd/amd64,
darwin/amd64, darwin/arm64, and windows/amd64。

# 运行时开销

**`数据竞态` 检测的成本因程序而异，一般来说，内存增加 5-10 倍，执行时间增加 2-20 倍**。

`数据竞态检测器` 为每个 `defer` 和 `recover` 语句额外分配 8 个字节，在 `goroutine` 退出之前，这些额外的分配不会被回收。
这意味着如果你有一个长期运行的 `goroutine` 定期调用 `defer` 和 `recover`，程序内存会无限增长。
这些内存分配不会显示在 `runtime.ReadMemStats` 或 `runtime/pprof` 的输出中。

通过上述官网描述的运行时开销，可以对我们的实践给予一定指导。比如: 生产环境中慎用 `-race`, 开发过程中单元测试使用 `-race` 多多益善。

# 扩展阅读

1. [Go 官方原文](https://go.dev/doc/articles/race_detector)
2. [Go 内存模型](https://go.dev/ref/mem/)
3. [条件编译](conditional_compilation.md)
4. [一个经典案例](https://dave.cheney.net/2014/06/27/ice-cream-makers-and-data-races)