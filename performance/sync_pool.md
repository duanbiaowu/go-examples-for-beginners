# 概述

`sync.Pool` 用来复用对象，减少内存分配，降低 GC 压力。

# 特性

> `sync.Pool` 的大小可伸缩，高负载时会动态扩容，池中的对象在不活跃时会被自动清理。

# 如何使用

只需实现 `sync.Pool` 对象的 `New` 方法即可，当对象池中没有对象时，将会调用自定义的 `New` 方法创建。

```go
package main

import (
	"sync"
)

type person struct {
	name string
	age  int
}

var (
	// 实现 New 方法
	personPool = sync.Pool{
		New: func() interface{} {
			return new(person)
		},
	}
)

func main() {
	// Get 方法从池中申请一个对象
	// 因为返回值是 interface{}, 这里再加一个类型转换
	tom := personPool.Get().(*person)

	tom.name = "Tom"
	tom.age = 6
	println(tom.name, tom.age)

	// Put 方法将对象归还到池中
	personPool.Put(tom)
}
```

# 普通方法

测试代码如下:

```go
package performance

import (
	"bytes"
	"testing"
)

func BenchmarkBufferWithPool(b *testing.B) {
	data := make([]byte, 1024)

	for n := 0; n < b.N; n++ {
		// 每次初始化一个新对象
		var buf bytes.Buffer
		buf.Write(data)
		buf.Reset()
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000 次，统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=1000x -benchmem > slow.txt
```

# 对象池

测试代码如下:

```go
package performance

import (
	"bytes"
	"sync"
	"testing"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func BenchmarkBufferWithPool(b *testing.B) {
	data := make([]byte, 1024)

	for n := 0; n < b.N; n++ {
		// 对象复用，从对象池中获取对象
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
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
name              old time/op    new time/op    delta
BufferWithPool-8    75.9ns ± 0%   500.0ns ± 0%    +558.94%  (p=1.000 n=1+1)

name              old alloc/op   new alloc/op   delta
BufferWithPool-8     2.00B ± 0%  1024.00B ± 0%  +51100.00%  (p=1.000 n=1+1)

name              old allocs/op  new allocs/op  delta
BufferWithPool-8      0.00           1.00 ± 0%       +Inf%  (p=1.000 n=1+1)
```

输出的结果分为了三行，分别对应基准期间的: 运行时间、内存分配总量、内存分配次数，采用了 `复用池` 方案后:

- 运行时间提升了 `5 倍+`
- 内存分配总量降低了 `500 倍+`
- 内存分配次数降至 0

因为时间关系，基准测试只运行了 1000 次，运行次数越大，优化的效果越明显。感兴趣的读者可以将 `-benchtime` 调大后看看优化效果。

## 性能分析

`普通方法` 每次重新申请一个新的 `bytes.Buffer` 对象，用完之后再释放掉，大部分时间浪费在了申请资源和释放资源上面，
而 `对象池` 通过资源池复用了 `bytes.Buffer` 对象，避免了申请资源和释放资源的时间损耗，所以性能远高于 `普通方法`。

## 使用建议

- `sync.Pool` 有自动清除对象机制，因此重要 (不可改变) 的对象不要使用 `sync.Pool`
- 从 `sync.Pool` 获取到的对象状态可能不同，例如长度会发生变化的数据类型 (如切片)
- 从 `sync.Pool` 获取到的对象使用完成后，要及时调用 `Put` 归还

# 小结

本小节对 `普通方法` 申请对象和 `对象池` 申请对象两种方法进行基准测试，并比较了两者在性能和内存方面的差异。
从结果中可以看到，`对象池` 所带来的性能提升是非常大的，这也提示我们，在申请和释放大的资源对象 (如数据库连接, 复杂的结构体)时， 可以使用 `sync.Pool` 来优化提升性能。

# 扩展阅读

- [一篇不错的实战文章](https://blog.thinkeridea.com/201901/go/you_ya_de_du_qu_http_qing_qiu_huo_xiang_ying_de_shu_ju.html)