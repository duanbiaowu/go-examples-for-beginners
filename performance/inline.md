# 概述

`内联 (inline)` 就是 **将函数的调用代码替换为函数的具体实现代码** (编译器实现)，程序运行过程中直接执行内联后展开的代码，
节省了函数调用的开销(创建栈帧、读写寄存器、栈溢出检测等)，可以提升性能，但是带来的一个问题是编译后的二进制文件体积增大。

接下来，我们先通过一个示例来了解下 `内联`。

# 示例

```go
// 原代码:
package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	z := max(1, 2)
	println(z)
}
```

## 内联后代码猜测

上面的代码，内联之后展开成类似下面的代码:

```go
package main

func main() {
	var z int
	if 1 > 2 {
		z = 1
	} else {
		z = 2
	}
	println(z)
}
```

当然，因为这个程序实在过于简单，编译器可以直接优化为:

```go
// 最终优化后代码
package main

func main() {
	println(2)
}
```

直接优化为一行代码， 编译器真的有这么强大吗？ 接下来我们通过构建和反汇编代码一起来验证一下。

## 内联后反汇编代码

对 `原代码` 进行内联构建:

```shell
# 构建时开启内联优化
$ go build -gcflags -m main.go

# 输出如下
# command-line-arguments
./main.go:3:6: can inline max
./main.go:10:6: can inline main
./main.go:12:10: inlining call to max
```

查看内联构建后的反汇编代码:

```shell
$ go tool objdump -s "main.main" main | grep CALL

# 输出如下
main.go:4             0x457c14                e84776fdff              CALL runtime.printlock(SB)
main.go:4             0x457c20                e83b7dfdff              CALL runtime.printint(SB)
main.go:4             0x457c25                e89678fdff              CALL runtime.printnl(SB)
main.go:4             0x457c2a                e8b176fdff              CALL runtime.printunlock(SB)
main.go:3             0x457c39                e802cdffff              CALL runtime.morestack_noctxt.abi0(SB)
```

## 验证猜测

对 `最终优化后代码` 进行构建:

```shell
# 构建时开启内联优化
$ go build main.go
```

查看内联构建后的反汇编代码:

```shell
$ go tool objdump -s "main.main" main | grep CALL

# 输出如下
main.go:4             0x457c14                e84776fdff              CALL runtime.printlock(SB)
main.go:4             0x457c20                e83b7dfdff              CALL runtime.printint(SB)
main.go:4             0x457c25                e89678fdff              CALL runtime.printnl(SB)
main.go:4             0x457c2a                e8b176fdff              CALL runtime.printunlock(SB)
main.go:3             0x457c39                e802cdffff              CALL runtime.morestack_noctxt.abi0(SB)
```

**最后，经过两次生成的反汇编代码对比，结果是一样的，这验证了我们的猜想，编译器的优化确实非常强大**。

# 内联条件

并不是所有条件下编译器都会内联，以下场景不会内联 (可能随着 Go 版本的变化而变化):

- for
- select
- defer
- recover
- go
- 闭包
- 不能以 go:noinline 或 go:unitptrescapes 作为编译指令

除此之外，还有其他的限制，当解析 `AST` 时，Go 申请了 `80 个节点` 作为内联的数量上限，每个节点都会消耗一个预算。
比如，`a = a + 1` 包含了5个节点：`AS, NAME, ADD, NAME, LITERAL`。

# 内联与非内联性能差异

## 测试代码

```go
package performance

import (
	"strconv"
	"testing"
)

type person struct {
	name string
	age  int
}

func Benchmark_Compare(b *testing.B) {
	b.StopTimer()

	// 初始化数据
	persons := make([]*person, 100)
	for i := range persons {
		persons[i] = &person{
			name: strconv.Itoa(i),
			age:  i,
		}
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for j := range persons {
			_ = persons[j].name
		}
	}
}
```

## 内联测试

运行测试，并将基准测试结果写入文件:

```shell
# 运行 10000000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=10000000x -benchmem > inline.txt
```

## 非内联测试

运行测试，并将基准测试结果写入文件:

```shell
# 运行 10000000 次, 统计内存分配
$ go test -gcflags "-N -l" -run='^$' -bench=. -count=1 -benchtime=10000000x -benchmem > noinline.txt
```

## 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 inline.txt noinline.txt

# 输出如下
name               old time/op    new time/op    delta
_Compare-8    63.3ns ± 0%   180.2ns ± 0%  +184.81%  (p=1.000 n=1+1)

name               old alloc/op   new alloc/op   delta
_Compare-8     0.00B          0.00B           ~     (all equal)

name               old allocs/op  new allocs/op  delta
_Compare-8      0.00           0.00           ~     (all equal)
```

从输出的结果中可以看到，默认的 `内联优化` 性能比 `非内联优化` 的性能提升了将近 `2 倍`。

# 更激进的内联

`gcflags` 参数可以设置多个 `-l` 选项，每多加 1 个，表示编译器将采用更加激进的内联方式，同时也可能生成更大的二进制文件。

- `-gcflags='-l -l'`    2 级内联
- `-gcflags='-l -l -l'` 3 级内联
- `gcflags=-l=4`        4 级别内联

# 禁用内联

大多数情况下，不需要禁用内联设置，这里提到的禁止方法暂且作为备忘。

## 单个函数禁用

加上 `//go:noinline` 编译指令，如下代码所示:

```go
//go:noinline
func max(x, y int) int {
if x > y {
return x
}
return y
}
```

## 全局禁用

```shell
# 编译时禁用内联优化
$ go build -gcflags="-l" main.go

# 同时禁止编译器优化和内联优化
$ go build -gcflags="-N -l" main.go
```

# 小结

默认的内联优化在不断优化和完善，这意味我们无需额外配置，然后定期升级 Go 版本，就可以享受到内联带来的性能提升红利。

# 扩展阅读

- [Go语言inline内联的策略与限制](https://www.pengrl.com/p/20028/)
- [聊聊Go内存优化和相关底层机制](https://wudaijun.com/2019/09/go-performance-optimization/)