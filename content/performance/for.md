---
title: Go 高性能之 for 循环
date: 2023-01-01
modify: 2023-01-01
---

# 概述

`for` 循环遍历时，第一个参数为遍历对象列表 (假设列表变量名为 `items`) 的当前索引，第二个参数为遍历对象列表的当前对象，一般来说，我们有两种方法获取到当前遍历到的元素:

- 使用列表变量名 + 索引，如 `items[1]`
- 直接使用第二个参数 

那么两者之间的性能差异有多大呢？我们通过基准测试来比较一下。

# 通过索引读取元素

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

func Benchmark_ForeachPersons(b *testing.B) {
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
			// 通过索引读取元素
			_ = persons[j].name
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000000 次
$ go test -run='^$' -bench=. -count=1 -benchtime=1000000x . > index.txt
```

# 通过第二个参数读取元素

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

func Benchmark_ForeachPersons(b *testing.B) {
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
		for _, p := range persons {
			// 通过第二个参数读取元素
			_ = p.name
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000000 次
$ go test -run='^$' -bench=. -count=1 -benchtime=1000000x . > value.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 index.txt value.txt

# 输出如下 
name               old time/op  new time/op  delta
_ForeachPersons-8  51.2ns ± 0%  71.5ns ± 0%  +39.44%  (p=1.000 n=1+1)
```

从输出的结果中可以看到，`for-range` 循环时使用第二个参数的方式要比使用索引的方式 `慢 40%`。

# 小结

本小节通过基准测试比较了在循环中使用索引和使用值的区别，**使用索引除了性能优势之外，还有一个额外的好处是可以修改元素值**。
而通过第二个参数获取到的元素，则无法直接修改 (因为获取到的元素是一个拷贝后的元素值，而非指针)。 
如果对性能有要求，建议使用索引获取元素，如果对性能没有要求，可以使用第二个参数获取元素 (这种方式代码可读性很好)。
