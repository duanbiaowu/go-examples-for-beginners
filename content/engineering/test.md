---
title: 单元测试基础必备
date: 2023-01-01
---

# 概述

测试旨在发现 bug，而不是证明其不存在。一个工程质量良好的项目，一定会有充分的单元测试和合理的测试覆盖率，**单元测试就是业务逻辑**。

> `go test` 命令用来对程序进行测试。

# 规则

**在一个目录中，以 `_test.go` 结尾的文件是 `go test` 编译的目标，`go build` 将会自动忽略。**
`go test` 工具扫描以 `_test.go` 结尾的文件来寻找特殊函数，并生成一个临时的 `main` 包来编译和运行，最后清除过程中产生的临时文件。

## 常用规则:

- 运行当前目录对应的包下面某个测试用例: `go test run='^Pattern$'`，其中单引号中为正则表达式
- 运行当前目录下的测试用例: `go test .`
- 运行子目录下的测试用例: `go test ./package_name`
- 运行当前目录以及所有子目录下的测试用例: `go test ./...`

# 四种函数

在以 `_test.go` 结尾的文件中，一共有 4 种类型的函数:
- 功能测试函数: `Test` 前缀命名，用来测试程序逻辑的正确性
- 基准测试函数: `Benchmark` 前缀命名，用来测试程序的性能
- 示例函数: `Example` 前缀命名，用来提供文档
- 模糊测试函数: `Fuzz` 前缀命名，用来提供自动化测试技术

# 功能测试

为了简化演示代码的复杂性，这里直接将测试函数写在 `main.go` 文件。

## 普通测试方法

### 测试未通过

首先写一个空方法，不实现具体的功能，来演示 `测试未通过`。

创建 `main.go` 文件，将如下代码写入:

```go
package main

func sum(numbers ...int) int {
    return 0
}

func main() {

}
```

创建 `main_test.go` 文件，将如下代码写入:
    
```go
package main

import "testing" // 引入 testing 包

func Test_sum(t *testing.T) {	// 功能测试以 `Test` 前缀命名
	if v := sum(); v != 0 {
		// t.Errorf 类似fmt.Printf()
		t.Errorf("sum() = %v, want %v", v, 0)
	}

	if v := sum(1); v != 1 {
		t.Errorf("sum() = %v, want %v", v, 1)
	}

	if v := sum(1, 2, 3); v != 6 {
		t.Errorf("sum() = %v, want %v", v, 6)
	}
}
```

使用 `go test` 命令运行测试

```shell
$ go test .
 
# 输出如下
--- FAIL: Test_sum (0.00s)
    main_test.go:12: sum() = 0, want 1
    main_test.go:16: sum() = 0, want 6
FAIL
FAIL    helloworld      0.001s
FAIL
```

通过输出可以看到测试失败了，主要原因在于 `sum` 函数的实现，接下来我们修正 `sum` 函数。

### 修正失败用例

将如下代码写入 `main.go` 文件:

```go
package main

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {

}
```

### 测试通过

使用 `go test` 命令运行测试

```shell
$ go test .
 
# 输出如下
ok      helloworld      0.001s
```

## 常用参数

### -count

**运行测试的次数，默认为 1**。

多次运行 `go test` 命令，可以看到输出的结果中有了一个 `(cache)` 标识

```shell
$ go test .
ok      helloworld      (cached)

$ go test .
ok      helloworld      (cached)

$ go test .
ok      helloworld      (cached)
```

这是因为源文件 `main.go` 和测试文件 `main_test.go` 都未发生变化，所以这里直接读取了测试的缓存结果。
通过使用参数 `-count=1` 可以达到 "禁用缓存" 的效果。

```shell
$ go test -count=1 .
# 输出如下
ok      helloworld      0.001s
```

### -v

**输出测试运行的详细信息**。

```shell
$ go test -v .
# 输出如下
=== RUN   Test_sum
--- PASS: Test_sum (0.00s)
PASS
ok      helloworld      0.001s


# 配合 -count 使用
$ go test -v -count=3  . 
# 输出如下
=== RUN   Test_sum
--- PASS: Test_sum (0.00s)
=== RUN   Test_sum
--- PASS: Test_sum (0.00s)
=== RUN   Test_sum
--- PASS: Test_sum (0.00s)
PASS
ok      helloworld      0.002s
```

通过使用参数 `-v`, 可以从输出结果中清晰地看到每个测试用例的运行情况。

### -timeout

**测试运行超时时间，默认为 10 分钟**。

### -run

**运行特定的测试函数，比如 `-run sum` 只测试函数名称中包含 `sum` 的函数**。

## 基于表的测试用例

在刚才的测试方法中，只写了 3 个测试用例，却写了 3 个不同的 `if` 语句，如果测试用例有几十上百个，那这种方法显然太不灵活了。
仔细观察 3 个 `if` 语句会发现除了参数有变化外， 其他部分都是一样的的，这时候就可以将相同的部分剥离出来，进行合并。

将如下代码写入 `main_test.go` 文件:

```go
package main

import "testing" // 引入 testing 包

func Test_sum(t *testing.T) { // 功能测试以 `Test` 前缀命名
	tests := []struct {
		numbers []int // 将可变参数转换为一个切片
		want    int   // 正确的返回值，用于和结果进行比较
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 2, 3},
			6,
		},
	}

	for _, tt := range tests {
		if got := sum(tt.numbers...); got != tt.want {
			// t.Errorf 类似fmt.Printf()
			t.Errorf("sum() = %v, want %v", got, tt.want)
		}
	}
}
```

```shell
$ go test -v .
# 输出如下
=== RUN   Test_sum
--- PASS: Test_sum (0.00s)
PASS
ok      helloworld      0.001s
````

在上述代码中，将 3 个测试用例合并到了一起，这样断言部分就只剩下 1 个 `if` 语句了，最重要的是，如果以后需要对测试用例增加/修改/删除，
仅需修改 `tests` 结构体切片就可以，其余部分无需改动。

比如可以增加一个负数测试用例，只需要追加一个结构体即可: 

```go
{
    []int{},
    0,
},
...
...
{
    []int{-1, -2, -3},
    -6,
},
```

### 测试失败时终止

默认情况下，所有测试用例都是独立的，如果其中一个用例测试失败，其他用例会继续运行测试，这样可以捕获到所有的失败测试用例。
如果希望测试失败时终止测试，可以将 `t.Errorf()` 函数更换为 `t.Fatalf()`。

# 基准测试

默认情况下，不会运行任何基准测试，参数 `-bench` 指定要运行的基准测试。

## 常用参数:

- `benchtime`           表示时间或运行次数，比如 `-benchtime=10s` 表示基准测试运行 10 秒，`-benchtime=100x` 表示基准测试运行 100 次
- `benchmem`            统计内存分配情况
- `cpuprofile`          CPU 性能剖析 `-cpuprofile=cpu.out`
- `memprofile=$FILE`    内存 性能剖析 `-memprofile=mem.out`
- `blockprofile=$FILE`  阻塞 性能剖析 `blockprofile=block.out`

## 基准测试用例

将如下代码追加到 `main_test.go` 文件中:

```go
func Benchmark_sum(b *testing.B) {
	for i := 0; i < b.N; i++ {  // b.N 表示测试用例运行的次数
		sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}
```

```shell
# 运行基准测试
$ go test -v -bench=.
# 输出如下
=== RUN   Test_sum
--- PASS: Test_sum (0.00s)
goos: linux
goarch: amd64
pkg: helloworld
cpu: Intel(R) Core(TM) i5-8300H CPU @ 2.30GHz
Benchmark_sum
Benchmark_sum-8         305606659                4.020 ns/op
PASS
ok      helloworld      1.631s
```

输出结果中的 `Benchmark_sum-8` 表示 `GOMAXPROCS` 的值为 8，也就是默认的 CPU 核数，这个值会影响到并发测试结果。
可以通过 `-cpu` 参数改变 `GOMAXPROCS`，-cpu 支持传入一个列表作为参数，如 `-cpu=2,4`。
除此之外，结果中还给出了相关的时间数据: 测试总用时 1.631 秒，`sum` 函数调用耗费 4.020 纳秒 (305606659 次调用的平均值),

因为基准测试运行器起初并不了解函数调用的具体耗时，所以 `b.N` 的值从一个小的数字慢慢增长到足够大的数字，直到能检测到稳定的调用耗时。
一般来说，`b.N` 的值从 1 开始，如果该用例能够在 `1s` 内完成，`b.N` 的值便会增加继续执行，越往后面，每次增加的值越大。

## ResetTimer

如果在 `Benchmark` 运行开始前，需要一些初始化准备工作 (例如初始化一些配置信息)，可以调用 `ResetTimer` 方法忽略掉这部分工作，不计入基准测试耗时中。

```go
func Benchmark_sum(b *testing.B) {
    time.Sleep(time.Second * 3) // 模拟耗时准备任务, 这个时间不会被计入基准测试耗时
    b.ResetTimer()              // 重置定时器
	
    for i := 0; i < b.N; i++ {  // b.N 表示测试用例运行的次数
        sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    }
}
```

## StartTimer, StopTimer

除了函数的初始化操作外，还有一种场景是: 函数每次调用前需要一些工作 (例如创建一些资源)，调用后需要一些工作 (例如关闭和释放这些资源)，
这两者的耗时同样不应该计入基准测试耗时中。

```go
func Benchmark_sum(b *testing.B) {
    b.StopTimer()
    time.Sleep(time.Second) // 模拟创建资源耗时 
    b.StartTimer()
	
    for i := 0; i < b.N; i++ {  // b.N 表示测试用例运行的次数
        sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    }

    b.StopTimer()
    time.Sleep(time.Second) // 模拟释放资源耗时 
    b.StartTimer()
}
```

# 示例函数

示例函数没有参数和返回值，可以给 `sum` 函数新增一个 `示例函数` 作为文档，例如像下面这个例子一样。

```go
func Example_sum() {
    fmt.Println(sum(0))
    fmt.Println(sum(1, 2, 3))
    // Output:
    // 0
    // 6
}
```

## 三个作用

- 文档 (主要目的)
- 可以通过 `go test` 运行的可执行测试
  - 如果示例函数内含有类似上面例子中的 `// Output:` 格式的注释，测试工具会执行示例函数，然后检查标准输出与示例函数的注释是否匹配。

```shell
  # 例如将 `// Output:` 改为如下代码:
  Output:
  0
  66
  # 运行测试
  $  go test -v -count=1 .
  # 输出如下，报错了
  === RUN   Test_sum
  --- PASS: Test_sum (0.00s)
  === RUN   Example_sum
  --- FAIL: Example_sum (0.00s)
  got:
  0
  6
  want:
  0
  66
  FAIL
  FAIL    helloworld      0.002s
  FAIL
```

- 提供手动实验代码
  - http://golang.org 是由 `godoc` 提供的文档服务，使用了 [Go Playground](https://go.dev/play/) 让用户可以在浏览器中在线编辑和运行每个示例函数。

# 先编译出 .test 文件

**使用场景**

1. 这台机器上编译，另一个地方跑单测；
2. debug 单测程序；

```shell
go test -c -o example.test
# 运行
 ./example.test
# 指定运行某一个文件
-test.timeout=10m0s -test.v=true -test.run=TestPutAndGetKeyValue
```


# 最佳实践

1. 测试用例失败时，应该输出有用的内容: 错误信息，输出参数，返回值以及正确的返回值
2. 测试用例修复后，应该运行所有测试，确保没有引入新的问题 - 回归测试
3. 测试代码中不要调用 `log.Fatal`, `os.Exit`, 因为这两个调用会阻止跟踪过程(一般来说，这两个函数只在 `main` 函数中调用)
4. 测试不应该在失败时终止，而是要在一次运行中尝试报告多个错误，因为错误发生的方式本身会揭露错误的原因
5. 测试真正需要关心的数据，保持测试代码的简洁和内部结构的稳定性

# 小结

本小节介绍了 Go 测试的 3 中相关函数: `功能测试函数`, `基准测试函数`, `示例函数` (模糊测试后面会单独讲)，还有几个常用的测试参数。
功能测试函数结果分为两种: `ok, pass` (通过), `FAIL` (测试未通过)，基准测试函数会输出程序耗时相关数据，示例函数可以作为文档，
同时可以通过添加 `// Output:` 注释来提供精确的函数返回值或输出结果，并在 `go test` 命令运行时自动对注释进行测试校验，
保证了调用方看到的文档示例中的输出结果一定是正确的。

题外话: 笔者在第一次了解到 `// Output:` 这个功能时，不得不惊叹于 Go 的工程化设计，细节决定成败。

# Reference

- [Go 圣经](https://book.douban.com/subject/27044219/)