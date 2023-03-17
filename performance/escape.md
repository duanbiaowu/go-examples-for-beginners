# 逃逸分析

**Go 语言的编译器使用 `逃逸分析` 决定哪些变量分配在栈上，哪些变量分配在堆上**。

在栈上分配和回收内存很快，只需要 2 个指令: `PUSH` + `POP`, 也就是仅需要将数据复制到内存的时间，而堆上分配和回收内存，一个相当大的开销是 `GC`。

## 特性

- 指向 `栈` 对象的指针不能分配`堆`上 (避免悬挂指针)
- 指向 `栈` 对象的指针在对象销毁时必须被同时销毁 (避免悬挂指针和内存泄露)

例如对于函数内部的变量来说，不论是否通过 `new` 函数创建，最后会被分配在 `堆` 还是 `栈`，是由编译器使用 `逃逸分析` 之后决定的。
具体来说，当发现变量的作用域没有超出函数范围，分配在 `栈` 上，反之则必须分配在 `堆` 上，也就是说: **如果函数外部没有引用，则优先分配在 `栈` 上，
如果函数外部存在引用，则必须分配在 `堆` 上**。所以，**闭包必然会发生逃逸**。

## 发生场景

- 变量占用内存过大 (如大的结构体) 
- 变量占用内存不确定 (如 链表, slice 导致的扩容)
- 变量类型不确定 (interface{})
- 指针类型
    - 函数返回变量地址 (如一个结构体地址)
- 闭包
- interface

过多的变量逃逸到堆上，会增加 `GC` 成本，我们可以通过控制变量的分配方式，尽可能地降低 `GC` 成本，提高性能。

## 分析命令

- 使用 go 命令

```shell
$ go tool compile -m main.go

# 或者

$ go build -gcflags='-m -l' main.go
```

- 反汇编

```shell
$ go tool compile -S main.go
```

# 示例

## 确定的数据类型和 interface{}

如果变量类型是 `interface`，那么将会 `逃逸` 到堆上。函数返回值应尽量使用确定的数据类型，避免使用 `interface{}`。

```go
package main

func main() {
	data := []interface{}{100, 200}
	data[0] = 100
}
```

```shell
$ go tool compile -m main.go

# 输出如下
main.go:3:6: can inline main
main.go:4:23: []interface {}{...} does not escape
main.go:4:24: 100 does not escape // 未发生逃逸
main.go:4:29: 200 does not escape // 未发生逃逸
main.go:5:2: 100 escapes to heap  // 发生逃逸
```

## 基准测试

这里以使用 `interface{}` 触发逃逸作为例子进行分析。

### 使用 interface{}

```go
package performance

import "testing"

const size = 1024

func genSeqNumbers() interface{} {
	var res [size]int
	for i := 0; i < len(res); i++ {
		if i <= 1 {
			res[i] = 1
			continue
		}
		res[i] = res[i-1] + res[i-2]
	}
	return res
}

func Benchmark_Compare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = genSeqNumbers()
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x -benchmem > slow.txt
```

### 使用确定的数据类型

```go
package performance

import "testing"

const size = 1024

func genSeqNumbers() [1024]int {
	var res [size]int
	for i := 0; i < len(res); i++ {
		if i <= 1 {
			res[i] = 1
			continue
		}
		res[i] = res[i-1] + res[i-2]
	}
	return res
}

func Benchmark_Compare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = genSeqNumbers()
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次, 统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x -benchmem > fast.txt
```

### 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 fast.txt slow.txt 

# 输出如下
name        old time/op    new time/op    delta
_Compare-8    1.44µs ± 0%    3.22µs ± 0%  +123.40%  (p=1.000 n=1+1)

name        old alloc/op   new alloc/op   delta
_Compare-8     0.00B       8192.00B ± 0%     +Inf%  (p=1.000 n=1+1)

name        old allocs/op  new allocs/op  delta
_Compare-8      0.00           1.00 ± 0%     +Inf%  (p=1.000 n=1+1)
```

从输出的结果中可以看到，通过使用确定的数据类型，运行时间提升了 `1 倍多`, 内存分配量和内存分配次数降为 `0`。

> 高性能 Tips: 在 `hot path` 上要尽可能返回具体的数据类型。

# 其他逃逸场景

接下来介绍几种常见的逃逸场景。

## 切片

当切片占用内存超过一定大小，或无法确定切片长度时，对象将分配在 `堆` 上。

### 逃逸场景

```go
package main

func main() {
	weekdays := 7
	data := make(map[string][]int)
	data["weeks"] = make([]int, weekdays)
	data["weeks"] = append(data["weeks"], []int{0, 1, 2, 3, 4, 5, 6}...)
}
```

```shell
$ go tool compile -m main.go

# 输出如下 (发生逃逸)
main.go:3:6: can inline main
main.go:5:14: make(map[string][]int) does not escape
main.go:6:22: make([]int, weekdays) escapes to heap
main.go:7:45: []int{...} does not escape
```

### 避免逃逸方案

如果切片容量较小，可以改为使用数组，避免发生逃逸，详情见 [切片和数组性能差异](slice_with_array.md)。

```go
package main

func main() {
	data := make(map[string][7]int)
	data["weeks"] = [...]int{0, 1, 2, 3, 4, 5, 6}
}
```

```shell
$ go tool compile -m main.go

# 输出如下 (没有发生逃逸)
main.go:3:6: can inline main
main.go:4:14: make(map[string][7]int) does not escape
```

## 指针

### 逃逸场景

```go
package main

func main() {
	n := 10
	data := make([]*int, 1)
	data[0] = &n
}
```

```shell
$ go tool compile -m main.go

# 输出如下 (发生逃逸)
main.go:3:6: can inline main
main.go:4:2: moved to heap: n
main.go:5:14: make([]*int, 1) does not escape
```

## interface{}

### 逃逸场景

```go
package main

func main() {
	data := make(map[interface{}]interface{})
	data[100] = 200
}
```

```shell
$ go tool compile -m main.go

# 输出如下 (发生逃逸)
main.go:3:6: can inline main
main.go:4:14: make(map[interface {}]interface {}) does not escape
main.go:5:2: 100 escapes to heap
main.go:5:2: 200 escapes to heap
```

## 函数返回值

这应该是比较常见的一种情况，在函数中创建了一个对象，返回了这个对象的指针。这种情况下，函数虽然退出了，但是因为指针的存在，
对象的内存不能随着函数结束而回收，因此只能分配在 `堆` 上。

### 逃逸场景

```go
package main

import "math/rand"

func foo(argVal int) *int {
	var fooVal1 = 11
	var fooVal2 = 12
	var fooVal3 = 13
	var fooVal4 = 14
	var fooVal5 = 15

	// 循环是防止编译器将 foo 函数优化为 inline
	// 如果不用随机数指定循环次数，也可能被编译器优化为 inline
	// 如果是内联函数，main 调用 foo 将是原地展开
	//    那么 fooVal1 ... fooVal5 相当于 main 作用域的变量
	//    即使 fooVal3 发生逃逸，地址与其他几个变量也是连续的
	n := rand.Intn(5) + 1
	for i := 0; i < n; i++ {
		println(&argVal, &fooVal1, &fooVal2, &fooVal3, &fooVal4, &fooVal5)
	}

	return &fooVal3
}

func main() {
	mainVal := foo(1)
	println(*mainVal, mainVal)
}
```

---

运行代码

```shell
$ go run main.go

# 输出如下  (发生逃逸)
0xc000114f58 0xc000114f38 0xc000114f30 0xc000120000 0xc000114f28 0xc000114f20
0xc000114f58 0xc000114f38 0xc000114f30 0xc000120000 0xc000114f28 0xc000114f20
13 0xc000120000
```

通过输出的结果可以看到，变量 `fooVal3` 的地址明显与其他变量地址不是连续的。

### 结果分析

```shell
$ go tool compile -m main.go

# 输出如下 (发生逃逸) 
main.go:16:16: inlining call to rand.Intn
main.go:24:6: can inline main
main.go:8:6: **moved to heap: fooVal3**

# 查看逃逸分析详情
$ go build -gcflags='-m -l' main.go

# 输出如下 (发生逃逸) 
./main.go:5:6: cannot inline foo: function too complex: cost 124 exceeds budget 80
...
./main.go:8:6: fooVal3 escapes to heap:
...
./main.go:8:6: **moved to heap: fooVal3**

# 或者
$ go tool compile -S main.go | grep "runtime.newobject"

# 输出如下 
0x0036 00054 (**main.go:8**)        CALL    runtime.newobject(SB)
rel 55+4 t=7 runtime.newobject+0

# main.go 第 8 行正好是 var fooVal3 = 13
```

## 通道

### 逃逸场景

```go
package main

func main() {
	ch := make(chan string)

	s := "hello world"

	go func() {
		ch <- s
	}()

	<-ch
}
```

```shell
$ go tool compile -m main.go

# 输出如下 (发生逃逸)
main.go:19:5: can inline main.func1
main.go:19:5: func literal escapes to heap
```

## 闭包

> 一个函数和对其周围状态（lexical environment，词法环境）的引用捆绑在一起（或者说函数被引用包围），这样的组合就是闭包（closure）。

简单来说，**闭包可以在一个函数内部访问到其外部函数的作用域**。

### 逃逸场景

```go
package main

func inc() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	in := inc()
	println(in()) // 1
	println(in()) // 2
}
```

```shell
$ go tool compile -m main.go

# 输出如下 (发生逃逸) 
main.go:3:6: can inline inc
main.go:5:9: can inline inc.func1
main.go:12:11: inlining call to inc
main.go:5:9: can inline main.func1
main.go:13:12: inlining call to main.func1
main.go:14:12: inlining call to main.func1
main.go:4:2: **moved to heap: n**
main.go:5:9: func literal escapes to heap
main.go:12:11: func literal does not escape
```

`inc()` 函数返回一个闭包函数，闭包函数内部访问了外部变量 `n`, 形成了引用关系，那么外部变量 `n` 将会一直存在，直到 `inc()` 函数被销毁，
所以最终外部变量 `n` 被分配到了 `堆` 上。

## 占用过大空间

**操作系统对内核线程使用的栈空间是有大小限制的，64 位系统上通常是 8 MB**。

```shell
# 查看系统栈内存大小

$ ulimit -s

# 8192
```

因为栈空间通常比较小，因此递归函数实现不当时，容易导致栈溢出。 对于 Go 语言来说，运行时(runtime) 尝试在 goroutine 需要的时候动态地分配栈空间，
goroutine 的初始栈大小为 2 KB。当 goroutine 被调度时，会绑定内核线程执行，栈空间大小也不会超过操作系统的限制。 对 Go 编译器而言，
超过一定大小的局部变量将逃逸到 `堆` 上，不同的 Go 版本的大小限制可能不一样。

### 逃逸场景

```go
package main

import "math/rand"

func generate8192() {
	nums := make([]int, 8192) // = 64KB
	for i := 0; i < 8192; i++ {
		nums[i] = rand.Int()
	}
}

func generate8193() {
	nums := make([]int, 8193) // < 64KB
	for i := 0; i < 8193; i++ {
		nums[i] = rand.Int()
	}
}

func generate(n int) {
	nums := make([]int, n) // 不确定大小
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
}

func main() {
	generate8192()
	generate8193()
	generate(1)
}
```

```shell
$ go tool compile -m main.go

# 输出如下 (发生逃逸) 
main.go:6:14: make([]int, 8192) does not escape
main.go:13:14: make([]int, 8193) escapes to heap
main.go:20:14: make([]int, n) escapes to heap
```

从输出的结果中可以看到，`make([]int, 8192)` 没有发生逃逸，`make([]int, 8193)` 和 `make([]int, n)` 逃逸到 `堆` 上，
说明当切片占用内存超过一定大小，或无法确定当前切片长度时，对象将分配到 `堆` 上。

# 扩展阅读

## 返回值和返回指针

`值传递`会拷贝整个对象，而 `指针传递` 只会拷贝地址，指向的对象是同一个。返回指针可以减少值的拷贝，但是会导致内存分配逃逸到堆中，
增加 `GC` 的负担。在对象频繁创建和删除的场景下，指针传递导致的 `GC` 开销可能会严重影响性能。

**一个通用的实践是: 对于需要修改原对象值，或占用内存比较大的结构体，返回指针，其他情况直接返回值**。

## 避免逃逸

在了解了 `逃逸` 发生的具体规则和场景后，我们可以通过对应的规则来避免逃逸，此外，也可以参考标准库中避免逃逸的方法。

### 标准库源码

```go
// Go 1.19 $GOROOT/src/strings/builder.go:27

// noescape hides a pointer from escape analysis. It is the identity function
// but escape analysis doesn't think the output depends on the input.
// noescape is inlined and currently compiles down to zero instructions.
// USE CAREFULLY!
// This was copied from the runtime; see issues 23382 and 7921.
//
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}
```

`noescape` **通过一个无用的位运算，解除了参数与返回值的关联，切断了 `逃逸分析` 对指针的追踪**，类似代码标准库中有很多，但是正如注释中所写，
**该方法要谨慎使用，除非明确了逃逸是性能瓶颈**。

```shell
# 查看标准库切断逃逸分析相关代码

$ grep -nr "func noescape" "$(dirname $(which go))/../src"
```

# 小结

本文介绍了几种常见的逃逸场景，并且针对其中的常见的逃逸发生场景做了详细的分析，感兴趣的读者可以采用本文提供的分析方法，看看自己的项目中有哪些对象逃逸场景。

在大多数情况下，关于变量应该分配到堆上还是栈上，**Go 编译器已经优化的足够好，无需开发者刻意优化**。为了保证绝对的内存安全，
编译器可能会将一些变量错误地分配到堆上，但是最终 `GC` 会避免内存泄露以及悬挂指针等安全问题，降低开发者心智负担，提高开发效率。

# Reference

- [极客兔兔](https://geektutu.com/post/hpg-escape-analysis.html)
- [escape.go](https://tip.golang.org/src/cmd/compile/internal/escape/escape.go)
- [akutz/lem](https://github.com/akutz/lem)
- [Go: Introduction to the Escape Analysis](https://medium.com/a-journey-with-go/go-introduction-to-the-escape-analysis-f7610174e890)