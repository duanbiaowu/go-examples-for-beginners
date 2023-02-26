# 概述

**如何理解 interface ?**

1. `interface` 是方法声明的集合
2. 某个类型实现了 `interface` 中声明的全部方法，表明该类型实现了该接口
3. `interface` 可以作为一种数据类型，实现了该接口的任何类型都可以赋值给该接口类型变量

Go 语言是松耦合的类型、方法对接口的实现，虽然没有面向对象编程语言中 `类` 的概念，但是面向对象的三大特性 `封装`、`继承`、`多态`, 在 Go 中同样可以实现。

# 封装

**Go 将访问层次简化为两层**:

- 不可导出: 通过标识符首字母小写，`对象` 仅包内可见
- 可导出: 通过标识符首字母大写，`对象` 对所有包可见

## 例子

### 不可导出

`Person` 类型及方法只能在 `foo` 包中使用。

```go
package foo

type person interface {
	name() string
	age() int
}
```

### 可导出

`Person` 类型及方法可以在所有包中使用。

```go
package foo

type Person interface {
	Name() string
	Age() int
}
```

# 继承

通过组合实现，当一个类型 `A` 内嵌另一个类型 `B` 的值或指针时，类型 `A` 可以使用类型 `B` 的所有方法，通过内嵌多个类型，可以达到 `多重继承` 的效果。

## 例子

### 单个继承

```go
package main

import (
	"log"
	"os"
)

type Person struct {
	Name        string
	Age         int
	*log.Logger // 等于继承了 log.Logger 的所有方法 
}

func main() {
	tom := Person{
		Name:   "Tom",
		Age:    6,
		Logger: log.New(os.Stdout, "", 0),
	}

	tom.Printf("My name is %s, age is %d", tom.Name, tom.Age)
}
```

```shell
$ go run main.go
# 输出如下 
My name is Tom, age is 6
```

在上面的示例代码中，`Person` 通过内嵌 `log.Logger` 类型 "继承" 了 `Printf` `Fatal` 等方法。

### 多重继承

```go
package main

import (
	"log"
	"os"
	"sync"
)

type Person struct {
	Name        string
	Age         int
	*log.Logger // 等于继承了 log.Logger 的所有方法
	*sync.Mutex // 等于继承了 sync.Mutex 的所有方法
}

func main() {
	tom := Person{
		Name:   "Tom",
		Age:    6,
		Logger: log.New(os.Stdout, "", 0),
		Mutex:  &sync.Mutex{},
	}

	tom.Lock()
	tom.Printf("My name is %s, age is %d", tom.Name, tom.Age)
	tom.Unlock()
}
```

```shell
$ go run main.go
# 输出如下 
My name is Tom, age is 6
```

在上面的示例代码中，`Person` 通过内嵌 `log.Logger` 类型和 `sync.Mutex` 类型， "继承" 了两者所有的方法，
如 `Printf` `Fatal` `Lock` `Unlock` 等。

### 重载

**Go 不支持重载操作。**

### 重写

通过 `定义同名方法` 实现，当一个类型 `A` 内嵌另一个类型 `B` 的值或指针时，类型 `A` 可以重新声明并定义类型 `B` 的所有方法，可以达到 `重写` 的效果。

```go
package main

import (
	"log"
	"os"
)

type Person struct {
	Name        string
	Age         int
	*log.Logger // 等于继承了 log.Logger 的所有方法
}

// 重写 Logger 的 Printf 方法
func (p *Person) Printf(format string, v ...any) {
	println("hello world")
}

func main() {
	tom := Person{
		Name:   "Tom",
		Age:    6,
		Logger: log.New(os.Stdout, "", 0),
	}

	// 输出 hello world
	tom.Printf("My name is %s", tom.Name)

	// 输出 hello world
	tom.Printf("My age is %d", tom.Age)
}
```

```shell
$ go run main.go
# 输出如下 
hello world
hello world
```

在上面的示例代码中，`Person` 通过内嵌 `log.Logger` "继承" 了 `log.Logger` 的所有方法，同时又重写了 `Printf` 方法，重写之后，
调用 `Printf` 方法时，不论参数是什么，都只会输出 `hello world`。

# 多态

通过接口实现，**某个类型的实例可以赋值给它所实现的任意接口类型的变量**。每个模块不需要了解其他模块的的细节，
`A` 模块定义接口，`B` 模块实现接口就可以，如果接口中没有使用到 `A` 模块中定义的数据类型，那么 `B` 模块中都不需要 `import A`,
也就是 **接口定义方和接口实现方不需要建立 `import 联系`**，非常优秀地实现了接口设计的 `正交性`。

## 例子

### 标准库的 io.Writer 接口

```go
package io

type Writer interface {
	Write(p []byte) (n int, err error)
}
```

上面是标准库中的 `io.Writer` 接口，我们可以在自己的模块中实现它:

```go
package main

type MyBuffer struct{} // 定义时并不需要 import "io"

func (m MyBuffer) Write(p []byte) (n int, err error) {
	return 0, nil
}
```

然后，在使用 `io.Writer` 类型作为参数的函数或方法中，可以直接传入 `MyBuffer` 类型，比如 `log` 包的 `SetOutput` 方法:

```go
package log

func SetOutput(w io.Writer) {
	output = w
}
```

上述代码中，通过接口规则，非常优雅地实现了 `接口的具体实现与解耦`。
在 Go 的标准库中，更是将上述准则应用到几乎所有 Go 包，**多态用得越多，代码就相对越少**。这被认为是 `Go 编程最佳实践之一`。

## 正交性的小问题

一个代码量庞大且结构复杂的项目中，会有很多 `接口`，例如 `ORM` 框架针对不同数据库的连接接口，我们希望知道该接口被哪些类型实现了,
反过来，针对一个具体的类型，我们希望知道它实现了哪些接口。 这个常用的功能，貌似只有 `Goland` 实现了。

# 小结

本小节概括了 Go 语言的 `面向对象` 特性以及如何使用。`封装` 通过首字母大小写区分，不需要 `public` `private` `protected`
等关键字，`继承` 通过类型嵌套实现，直接从语法层面践行了 `组合优于继承` 的理念，`多态` 通过接口实现，从语法层面将接口定义方和接口实现方解耦，
实现了优秀的 `正交性` 设计。

在开发过程中，可以直接添加新接口，**已有的接口无需改动，已有的类型只需实现新接口的方法即可**，
对于已有的函数，可以将函数参数类型扩展成使用 `接口类型` 的约束性参数。整个系统设计可以持续演进，而不用推翻之前的方案。
在代码层面，不需要像基于 `类` 的面向对象语言一样，维护和适应 `整个类层次结构` 的变化。

# 扩展阅读

## Go 和 Java 接口差异

### 侵入性

`Java` 语言中的接口是强制性的，`类` 必须声明实现了某个接口，并实现该接口的所有方法，这种机制被称为 `侵入式` 接口。

`Go` 语言中接口核心理念是 `组合`，而且接口定义方和接口实现方以原生方式解耦，扩展性更强，这种机制被成为 `非侵入式` 接口。
一个类型只需要实现接口定义的所有方法，那么该类型就自动实现了接口，这种方式属于 `隐式实现` 而非 `显式声明`。

> 类型能做什么比类型是什么更重要。

### Go 接口设计思想

Go 结合了 `接口值`，`编译器静态类型检查`（该类型是否实现了某个接口），`运行时动态转换`，并且不需要 `显式地声明` 类型要实现某个接口。
这个特性可以在不修改已有代码的情况下定义和实现新接口，从而提供了 `动态语言` 的优点，但却没有 `动态语言` 在运行时可能发生错误的问题。