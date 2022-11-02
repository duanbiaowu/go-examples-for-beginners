# 概述

# 语法规则
如果 defer 函数只有一行语句，可以省略 `func() { ... }` 代码块，否则就需要用 `func() { ... }` 代码块包起来。 

# 多个 defer 执行顺序
如果一个函数中注册了多个 defer 函数，这些函数会按照 `后进先出` 的顺序执行 (和 `栈` 的出栈顺序一致)。
也就是最后注册的 defer 函数会第一个执行，而第一个注册的 defer 函数会最后执行。

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