# 概述

> Benchstat 命令用来计算和比较基准测试的统计数据。

# 安装 benchstat

```shell
$ go install golang.org/x/perf/cmd/benchstat@latest
# 输出如下
go: downloading golang.org/x/perf v0.0.0-20220920022801-e8d778a60d07
go: downloading github.com/google/safehtml v0.0.2

# 安装完成
# 查看使用帮助
$ benchstat -h

usage: benchstat [options] old.txt [new.txt] [more.txt ...]
options:
  -alpha α
        consider change significant if p < α (default 0.05)
  ...
  ...
  -split labels
        split benchmarks by labels (default "pkg,goos,goarch")

```

# 参数规则

- 当参数是单个文件时，打印该文件中的 `Benchmark` 统计结果
- 当参数是两个文件时，打印两个文件的 `Benchmark` 统计结果以及比较信息
- 当参数是两个以上文件时，分别打印所有文件的 `Benchmark` 统计结果

# 例子

这里以一个 `求和函数` 的小例子作为演示，首先实现一个较慢的版本，接着在这个基础上进行优化，最后使用 `benchstat` 比较两者差异。

## 较慢的函数

将如下代码写入 `main.go` 文件中:

```go
package main

// 求和函数 (较慢的版本)
func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}
```

将如下代码写入 `main_test.go` 文件中:

```go
package main

import "testing"

func Benchmark_sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(100)
	}
}
```

运行基准测试，并将测试结果写入 `slow.txt` 文件

```shell
# 因为函数实现过于简单，设置运行次数为 1000 万
$ go test -run='^$' -bench=. -count=1 -benchtime=10000000x  > slow.txt

# 查看基准测试结果
$ cat slow.txt 

goos: linux
goarch: amd64
pkg: helloworld
cpu: Intel(R) Core(TM) i5-8300H CPU @ 2.30GHz
Benchmark_sum-8         10000000                34.10 ns/op
PASS
ok      helloworld      0.345s
```

从输出文件结果中可以看到，函数每次执行耗时 34.10 纳秒。

## 较快的函数

在较慢的函数基础上，使用 `高斯公式` 进行优化，将如下代码写入 `main.go` 文件:

```go
package main

// 求和函数 (优化后版本)
func sum(n int) int {
	return (1 + n) * n / 2 // 高斯公式
}
```

运行基准测试，并将测试结果写入 `fast.txt` 文件

```shell
# 因为函数实现过于简单，设置运行次数为 1000 万
$ go test -run='^$' -bench=. -count=1 -benchtime=10000000x  > fast.txt

# 查看基准测试结果
$ cat fast.txt 

goos: linux
goarch: amd64
pkg: helloworld
cpu: Intel(R) Core(TM) i5-8300H CPU @ 2.30GHz
Benchmark_sum-8         10000000                 0.2653 ns/op
PASS
ok      helloworld      0.006s
```

从输出文件结果中可以看到，函数每次执行耗时 0.2653 纳秒，**优化后的性能非常高！**

## benchstat 比较差异

最后，使用 `benchstat` 命令来比较较慢版本和高斯公式版本的差别:

```shell
$ benchstat slow.txt fast.txt
# 输出结果如下
name    old time/op  new time/op  delta
_sum-8  34.1ns ± 0%   0.3ns ± 0%   ~     (p=1.000 n=1+1)

# 第 1 列: 基准测试函数名称 
# 第 2 列: 较慢函数基准测试统计
# 第 3 列: 高斯公式基准测试统计
# 第 4 列: 优化比例
```

### -alpha 参数

从刚才 `benchstat` 命令的输出结果中可以看到，高斯公式优化的比例非常大，但是为什么 `delta` 列却显示的是 `~` 呢？
这是因为 `benchstat` 有一个 `-alpha` 参数用来表示 `重大变化的考虑因素`，默认值为 0.05 (
虽然笔者也没查到这个值的具体计算方式)，
但是这里不妨先把这个值设置的大一些，看看是否可以输出 `delta` 列的数据。

```shell
# 将 -alpha 参数设置为 100
$ benchstat -alpha=100 slow.txt fast.txt
# 输出如下
name    old time/op  new time/op  delta
_sum-8  34.1ns ± 0%   0.3ns ± 0%  -99.22%  (p=1.000 n=1+1)
```

可以对结果简单地进行验证: `( 34.10 - 0.2653 ) / 34.10 = 0.9922199....`，这下和 `delta` 列数据对应上了。

> Tips: 对于优化空间非常大的基准测试结果，比较时应该加上 -alpha 参数。

# 小结

本小节主要介绍了 `benchstat` 命令的基本应用，使用一个求和函数作为示例，并将两个实现方案 (普通累计和高斯公式) 进行对比，
因为优化的空间非常大，最后使用 -alpha 参数才输出了 `delta` 列的数据。

> 题外话：示例直观地表现出了算法优化的威力，而算法的基础是数学。据说高斯公式是高斯在 8 岁的时候提出来的，数学王子的美誉名不虚传。
