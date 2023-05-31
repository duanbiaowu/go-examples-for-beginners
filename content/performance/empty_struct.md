---
title: Go 高性能之空结构体
date: 2023-01-01
modify: 2023-01-01
---

# 概述

Go 的标准库没有内置的 `Set` 类型，在不引用第三方包的情况下，一般是结合内置的 `map` 类型实现 `Set` 相关功能。

# map 实现 set

这里假设 `Set` 元素类型为 `int`, 那么我们就以 `int` 作为 `map` 的键类型，以 `bool` 作为 `map` 的值类型
(之所以选择 `bool` 类型，是因为其大小为 1 个字节，相对其他数据类型可以节省内存，当然，也可以使用 `byte` 类型，其大小同样是 1 个字节)。

```go
package main

import "fmt"

// Set 类型定义
type set map[int]bool

// 初始化一个新的 Set
func newSet() set {
	return make(set)
}

// 元素是否存在于与集合中
func (s set) contains(ele int) bool {
	_, ok := s[ele]
	return ok
}

// 添加元素到集合
func (s set) add(ele int) {
	s[ele] = true
}

// 从集合中删除某个元素
func (s set) remove(ele int) {
	delete(s, ele)
}

func main() {
	s := newSet()

	fmt.Println(s.contains(100))
	s.add(100)
	fmt.Println(s.contains(100))
	s.remove(100)
	fmt.Println(s.contains(100))
}
```

```shell
$ go run main.go

# 输出如下
false
true
false
```

上述示例代码通过内置类型 `map` 实现了 `set` 类型相关功能，美中不足的一点在于: 每个元素都要浪费一个 `bool` 类型作为标识占位符，
能否进一步的优化呢？ 当然是可以的，这就是接下来要讲到的 `空结构体` 了。

# 空结构体

Go 中的空结构体 `struct{}{}` 是一个底层的内置变量，不占用任何内存:

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("size = %d\n", unsafe.Sizeof(struct{}{}))
}
```

```shell
$ go run main.go

# 输出如下
size = 0
```

结合刚才的例子，可以将 `struct{}{}` 作为 `Set` 的元素，这样不论 `Set` 有多少个元素，`标志位` 内存占用始终为 0 。

# 使用 bool 实现 Set 

## 测试代码

```go
package performance

import (
	"testing"
)

// Set 类型定义, 使用 bool 类型作为占位符
type set map[int]bool

// 初始化一个新的 Set
func newSet() set {
	return make(set)
}

// 元素是否存在于与集合中
func (s set) contains(ele int) bool {
	_, ok := s[ele]
	return ok
}

// 添加元素到集合
func (s set) add(ele int) {
	s[ele] = true
}

// 从集合中删除某个元素
func (s set) remove(ele int) {
	delete(s, ele)
}

func Benchmark_Compare(b *testing.B) {
	s := newSet()

	for i := 0; i < b.N; i++ {
		s.add(i)
	}
	for i := 0; i < b.N; i++ {
		_ = s.contains(i)
		s.remove(i)
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
$ go test -run='^$' -bench=. -count=1 -benchtime=10000000x -benchmem > bool.txt
```

# 使用空结构体实现 Set

## 测试代码

```go
package performance

import (
	"testing"
)

// Set 类型定义, 使用 bool 类型作为占位符
type set map[int]struct{}

// 初始化一个新的 Set
func newSet() set {
	return make(set)
}

// 元素是否存在于与集合中
func (s set) contains(ele int) bool {
	_, ok := s[ele]
	return ok
}

// 添加元素到集合
func (s set) add(ele int) {
	s[ele] = struct{}{}
}

// 从集合中删除某个元素
func (s set) remove(ele int) {
	delete(s, ele)
}

func Benchmark_Compare(b *testing.B) {
	s := newSet()

	for i := 0; i < b.N; i++ {
		s.add(i)
	}
	for i := 0; i < b.N; i++ {
		_ = s.contains(i)
		s.remove(i)
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
$ go test -run='^$' -bench=. -count=1 -benchtime=10000000x -benchmem > struct.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 bool.txt struct.txt

# 输出如下
name        old time/op    new time/op    delta
_Compare-8     371ns ± 0%     332ns ± 0%  -10.47%  (p=1.000 n=1+1)

name        old alloc/op   new alloc/op   delta
_Compare-8     44.0B ± 0%     40.0B ± 0%   -9.09%  (p=1.000 n=1+1)

name        old allocs/op  new allocs/op  delta
_Compare-8      0.00           0.00          ~     (all equal)
```

从输出的结果中可以看到，相比于使用 `bool` 作为 `Set` 元素占位符，使用 `空结构体` 在性能和内存占用方面，都有了小幅度的优化提升。 
因为时间关系，这里的基准测试只运行了 10000000 次，运行次数越大，优化的效果越明显。感兴趣的读者可以将 `-benchtime` 调大后看看优化效果。

# 小结

Go 中的空结构体 `struct{}{}` 不占用任何内存，而且有很清晰的语义性质 (作为占位符使用)。除了刚才示例中实现 `Set` 功能外，
还可以使用空结构体作为 `通道信号标识`、`空对象` 等，各种使用场景请读者自行探究。

# 彩蛋

除了空结构体 `struct{}{}` 之外，还有一个鲜为人知的大小为 0 的数据类型是: `空数组`。

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("size = %d\n", unsafe.Sizeof([0]int{}))
}
```

```shell
$ go run main.go

# 输出如下
size = 0
```