---
title: 单元测试基境
date: 2023-01-01
---

# 概述

编写测试代码时，最繁琐的部分之一是将一些 `公共的状态变量` 设置为某个特定的状态。比如常见的场景:

- 测试开始时，打开一个数据库连接，测试过程中所有测试用例共享这个连接，测试结束时关闭这个连接
- 测试开始时，创建一个临时文件，测试过程中所有测试用例共享这个文件句柄，测试结束时关闭并删除这个文件

这种 **被共享且状态变化次数很少的值** 称为测试的 `基境`，**最佳实践是将测试代码中可以复用的部分** 放入 `基境`。

# TestMain

`TestMain` 函数运行在主 `goroutine` 中 , 可以在调用 `m.Run` 前后设置 `钩子` 函数。
如果测试文件中包含 `TestMain(*testing.M)` 函数，
所有测试方法必须调用参数的 `Run` 方法触发，如果不调用该方法，直接运行 `go test` 没有任何效果。 可以利用这一特性，
在 `Run` 方法执行前后分别挂载 `钩子` 函数实现 `基境` 功能。在 `TestMain` 函数末尾，应该使用 `m.Run`
的返回值作为参数去调用 `os.Exit`。

# 示例

首先写 4 个数据库的基础操作方法: 创建、更新、读取、删除，这里只是作为 `基境` 演示，方法内部并不会去实现具体的功能。

## CURD 方法

将如下代码写入 `main.go` 文件中:

```go
package main

func create() {
	println("create internal")
}

func update() {
	println("update internal")
}

func get() {
	println("get internal")
}

func delete() {
	println("delete internal")
}

func main() {

}
```

## 测试基境

将如下代码写入 `main_test.go` 文件中:

```go
package main

import (
	"os"
	"testing"
)

// 基境初始化操作
func setup() {
	// 模拟数据库建立连接操作
	println("database connected")
}

// 基境释放资源操作
func tearDown() {
	// 模拟数据库断开连接操作
	println("database closed")
}

// 注意参数是 *testing.M, 而非 *testing.T
func TestMain(m *testing.M) {
	setup()
	code := m.Run() // 调用 m.Run() 触发所有 Test_* 测试函数
	tearDown()
	os.Exit(code)
}

func Test_create(t *testing.T) {
	create() // 演示专用，内部不做任何测试
}

func Test_update(t *testing.T) {
	update() // 演示专用，内部不做任何测试
}

func Test_get(t *testing.T) {
	get() // 演示专用，内部不做任何测试
}

func Test_delete(t *testing.T) {
	remove() // 演示专用，内部不做任何测试
}
```

在上面的代码中，4 个数据库操作函数，因为每个操作都必须持有一个数据库连接对象 (这里只是模拟， 并没有真正去实现代码)，
所以这里可以将 `数据库的连接和断开两个操作` 提取为 `基境`。 接着将 `基境` 放入 `setup` 和 `tearDown` 两个 `钩子` 函数，
分别对应数据库的连接和断开操作，最后将两个函数挂载到了 `m.Run()` 执行的前后，这样就可以自动化运行 `基境` 操作了。

## 运行测试

```shell
$go test -v -count=1 . 
# 输出如下 
database connected
=== RUN   Test_create        
create internal
--- PASS: Test_create (0.00s)
=== RUN   Test_update        
update internal
--- PASS: Test_update (0.00s)
=== RUN   Test_get
get internal
--- PASS: Test_get (0.00s)   
=== RUN   Test_delete        
remove internal
--- PASS: Test_delete (0.00s)
PASS
database closed
ok      helloworld      0.002s
```

从输出的结果中可以看到，第一行输出了数据库连接操作，最后一行输出了数据库断开连接操作，中间部分是 4 个测试函数。

# 小结

本小节主要介绍了单元测试中 `基境` 的场景和基本用法，实际开发过程中，`基境` 的代码不会像示例代码中这么简单，会随着测试代码量的增加而增长，
比较好的实践是: **合理切分包、单个文件的功能和大小，将不同的测试函数进行分类，在这个基础上提取 `基境` 函数功能**。 