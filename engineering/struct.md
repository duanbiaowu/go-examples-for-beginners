# 概述

`Go` 语言中，**结构体和它所包含的数据在内存中是以连续块的形式存在的**，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。
不像 `Java` 中的引用类型，一个对象和它里面包含的对象可能会在不同的内存空间中，和 `Go` 语言中的指针很像。
下面的例子清晰地说明了这些情况：

```go
type Point struct {X, Y int}

type Rect1 struct {Min, Max Point }
type Rect2 struct {Min, Max *Point }
```

![结构体内存布局](./images/struct_mem_layout.png)

# 强制字面量方式创建结构体

在一个结构体中定义一个非导出的零大小字段，编译器将会禁止使用非字面量 (不指明字段名称) 来创建结构体。
备注: **该方法仅针对包外调用，包内调用不受影响**。

## 编译失败

新建 `foo/person.go` 文件, 将如下代码写入: 

```go
package foo

type Person struct {
	_    [0]int
	Name string
	Age  int
}
```

新建 `main.go` 文件, 将如下代码写入:

```go
package main

import "helloworld/foo"

func main() {
	_ = foo.Person{[0]int{}, "bar", 123}
}
```

```shell
$ go run main.go
# 输出如下 
./main.go:6:17: implicit assignment to unexported field _ in foo.Person literal
```

## 编译通过

```go
package main

import "helloworld/foo"

func main() {
	p := foo.Person{Name: "bar", Age: 123}
	println(p.Name, p.Age)
}
```

```shell
$ go run main.go
# 输出如下 
bar 123
```
