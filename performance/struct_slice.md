# 概述

业务开发中，一个常见的场景是将多个相同类型的 `结构体` 变量存入一个数据容器中，通常我们会使用 `切片` 作为数据容器。
那么对于结构体来说，存储其值和存储其指针，性能差异有多大呢？

# 切片元素为结构体

测试代码如下:

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

func Benchmark_PointerWithValue(b *testing.B) {
	b.StopTimer()

	// 初始化数据
	persons := make([]person, 10000)
	for i := range persons {
		persons[i] = person{
			name: strconv.Itoa(i),
			age:  i,
		}
	}

	b.StartTimer()

	for n := 0; n < b.N; n++ {
		// 切片存储结构体的值
		clonedPersons := make([]person, 10000)
		for i := range persons {
			clonedPersons[i] = persons[i]
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 10000 次, 统计内存分配
$ go test -run='^$' -bench='Benchmark_PointerWithValue' -count=1 -benchtime=10000x -benchmem > value.txt 
```

# 切片元素为结构体指针

测试代码如下:

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

func Benchmark_PointerWithValue(b *testing.B) {
	b.StopTimer()

	// 初始化数据
	persons := make([]person, 10000)
	for i := range persons {
		persons[i] = person{
			name: strconv.Itoa(i),
			age:  i,
		}
	}

	b.StartTimer()

	for n := 0; n < b.N; n++ {
		// 切片存储结构体的指针
		clonedPersons := make([]*person, 10000)
		for i := range persons {
			clonedPersons[i] = &(persons[i])
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 10000 次, 统计内存分配
$ go test -run='^$' -bench='Benchmark_PointerWithValue' -count=1 -benchtime=10000x -benchmem > pointer.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 pointer.txt value.txt 

# 输出如下
name                 old time/op    new time/op    delta
_PointerWithValue-8    34.6µs ± 0%    65.2µs ± 0%   +88.38%  (p=1.000 n=1+1)    

name                 old alloc/op   new alloc/op   delta
_PointerWithValue-8    81.9kB ± 0%   245.8kB ± 0%  +200.00%  (p=1.000 n=1+1)    

name                 old allocs/op  new allocs/op  delta
_PointerWithValue-8      1.00 ± 0%      1.00 ± 0%      ~     (all equal)
```

输出的结果分为了三行，分别对应基准期间的: 运行时间、内存分配总量、内存分配次数，相对于 `结构体值` 类型，采用了 `结构体指针` 类型后:

- 运行时间提升了将近 `1 倍`
- 内存分配总量减少了 `2 倍`

感兴趣的读者可以调整 `benchtime` 参数大小，观察一下性能的变化趋势。

# 小结

当需要将多个相同类型的 `结构体` 变量存入一个 `切片` 时，请存储 `结构体` 变量的 `指针`。

