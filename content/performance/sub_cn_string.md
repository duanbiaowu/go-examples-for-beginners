---
title: Go 高性能之截取中文字符串
date: 2023-01-01
modify: 2023-01-01
---

# 概述

如果字符串中全部都是 `ASCII` 字节，直接使用切片的方式截取，是最简单和最高效的方式，如:

```go
package main

func main() {
	s := "hello world"
	s2 := s[2:5]
	println(s2) // llo
}
```

但是，如果字符串中有中文，这种方式会出现乱码:

```go
package main

func main() {
	s := "Go 语言的优势是什么？"
	s2 := s[2:5]
	println(s2) //  � 
}
```

**如何从一个中英文 + 数字混合的字符串中，截取一部分中文字符串呢？**

# rune

首先能想到的是将字符串转换为 `[]rune` 类型，这样就不会出现乱码问题了。

```go
package main

import "fmt"

func main() {
	s := "Go 语言的优势是什么？"
	rs := []rune(s)
	s2 := rs[2:5]
	fmt.Printf("%s\n", string(s2)) // 语言
}
```

虽然可以得到正确答案，但是代码中出现了两次类型转换过程，我们来做一下基准测试，代码如下:

```go
package performance

import (
	"testing"
)

func Benchmark_SubCNString(b *testing.B) {
	s := "Go 语言的优势是什么？"

	for i := 0; i < b.N; i++ {
		rs := []rune(s)
		s2 := rs[0:5]
		_ = string(s2)
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次，统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x > slow.txt
```

# range

`range` 遍历字符串时，默认使用 `字符` 迭代，也就是 `ASCII` 和 `中文` 都算作一个 `字符`, 可以利用该特性来实现截取功能:

```go
package main

import "fmt"

// 这里作为测试，只提供一个参数，就是截取的长度
// 感兴趣的读者可以在此基础上扩展，加一个参数，截取的起始位置
func substr(s string, length int) string {
	var cnt, start int
	for start = range s {
		if cnt == length {
			break
		}
		cnt++
	}

	return s[:start]
}

func main() {
	s := "Go 语言的优势是什么？"

	fmt.Printf("%s\n", substr(s, 5)) // Go 语言
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 100000 次，统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=100000x > fast.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 fast.txt slow.txt 

# 输出如下:
name            old time/op    new time/op    delta
_SubCNString-8    16.1ns ± 0%   212.1ns ± 0%  +1214.13%  (p=1.000 n=1+1)
```

从输出的结果中可以看到，采用了 `range` 方案后，性能提升了 `12 倍` 之多。

# 扩展阅读

- [一个更高效和通用的类库](https://blog.thinkeridea.com/201910/go/efficient_string_truncation.html)