# 概述

**一个 `defer` 语句就是一个普通的函数或方法调用。** `defer` 语句保证了不论是在正常情况下 (return 返回)，
还是非正常情况下 (发生错误, 程序终止)，函数或方法都能够执行。

## 主要特性

- 一个函数可定义多个 `defer` 语句
- `defer` 表达式中的变量值在 `defer` 表达式定义时已经确定
- `defer` 表达式可以修改函数中的命名返回值

## 主要作用

- 简化异常处理 ( 使用 `defer` + `recover`)，避免异常与控制流混合在一起 (`try … catch … finally`)
- 在 `defer` 做资源释放和配置重置等收尾工作

# 语法规则

如果 `defer` 函数只有一行语句，可以省略 `func() { ... }` 代码块，否则就需要用 `func() { ... }` 代码块包起来。

# 多个 defer 执行顺序

如果一个函数中注册了多个 `defer` 函数，这些函数会按照 `后进先出` 的顺序执行 (和 `栈` 的出栈顺序一致)。
也就是最后注册的 defer 函数会第一个执行，而第一个注册的 `defer` 函数会最后执行。

# 例子

## 函数退出前打印字符

```go
package main

func A() {
	defer println("A 函数执行完成")

	println("A 函数开始执行")
}

func B() {
	defer println("B 函数执行完成")

	println("B 函数开始执行")
}

func main() {
	A()
	B()
}

// $ go run main.go
// 输出如下 
/**
  A 函数开始执行
  A 函数执行完成
  B 函数开始执行
  B 函数执行完成
*/
```

## 关闭文件句柄

```go
package main

import (
	"fmt"
	"os"
)

func createFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(file *os.File) {
	n, err := file.WriteString("hello world")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("成功写入 %d 个字符\n", n)
	}
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	file := createFile("/tmp/defer_test.txt")
	defer closeFile(file) // 获取到文件句柄后，第一时间注册 defer 函数

	writeFile(file)
}

// $ go run main.go
// 输出如下 
/**
  成功写入 11 个字符
*/

// $ cat /tmp/defer_test.txt
// 输出如下
/**
  hello world
*/
```

## 多个 defer 函数

```go
package main

func A() {
	defer println("第 1 个 defer 函数")

	defer func() { // 这里为了演示 func() { ... } 的语法
		defer println("第 2 个 defer 函数")
	}()

	defer println("第 3 个 defer 函数")

	println("A 函数开始执行")
}

func main() {
	A()
}

// $ go run main.go
// 输出如下
/**
  A 函数开始执行
  第 3 个 defer 函数
  第 2 个 defer 函数
  第 1 个 defer 函数
*/
```

# reference

1. [Go 圣经](https://book.douban.com/subject/27044219/)