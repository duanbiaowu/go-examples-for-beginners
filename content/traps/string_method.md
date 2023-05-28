---
date: 2023-01-01
---

# 概述

通常我们会对某个对象自定义一些方法，大多数情况下，这没有任何问题。但是有一种情况需要注意下，那就是自定义的 `String` 方法。

Go 标准库中有一个 `Stringer` 接口，原型如下:

```go
type Stringer interface {
	String() string
}
```

文件路径为 `$GOROOT/src/fmt/print.go`，笔者的 Go 版本为 `go1.19 linux/amd64`。

如果某个对象实现了自定义 `String` 方法，那么等于实现了 `Stringer` 接口。 如果在方法内部实现中调用了 `fmt.Prinf*` 系列方法，会导致错误。

# 内存溢出

## 错误的做法

```go
package main

import (
	"fmt"
)

type number int

func (n number) String() string {
	return fmt.Sprintf("%v", n)
}

func main() {
	var n number = 100
	println(n.String())
}

// $ go run main.go
// 没有任何输出，阻塞住，内存耗尽...
```

**错误原因**: 类型 `number` 自定义 `String` 方法实现了 `Stringer` 接口，方法内部调用了 `fmt.Sprintf` 方法, 
然而 `fmt.Sprintf` 方法内部会检测参数是否实现了 `Stringer` 接口，如果实现了，会直接调用参数的 `String` 方法, 
等于又反过来调用 `number.String` 方法, 这样就进入了无限递归，最终内存溢出。

![Stringer](/images/Stringer.png)

## 正确的做法

实现 `String` 方法时，不调用内部调用了 `String` 的方法，避免无限递归。

```go
package main

import (
	"strconv"
)

type number int

func (n number) String() string {
	return strconv.Itoa(int(n))
}

func main() {
	var n number = 100
	println(n.String())
}

// $ go run main.go
// 输出如下 
/**
  100
*/
```

# 扩展阅读

- [Go 的面向对象编程](https://dbwu.tech/posts/golang_oop/)
- [如何实现 implements](https://dbwu.tech/posts/golang_implements/)