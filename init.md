# 概述

**`init() 函数` 是一个特殊的函数，一般称为初始化函数，不能被调用。** 在每个文件里面，当程序启动或者文件被作为包引用的时候，
init() 函数就会自动执行，一般用来做一些包的初始化操作。

# 语法规则

`init() 函数` 没有参数，也没有返回值。

```shell
func init() {
    // do something
}
```

# 执行顺序

包的初始化函数按照程序中引入的顺序执行，依赖于具体的顺序优先级，每次初始化一个包。
例如 `包 a` 引入了 `包 b`, 那么确保 `包 b` 的初始化操作 在 `包 a` 的初始化操作之前完成。
初始化过程是自下而上的，`main 包` 最后初始化，也就是说，在 `main 函数` 执行前，
所引用到的包已经全部初始化完成。

```shell
import -> const -> var -> init() -> main()
```

## 例子

### 包变量初始化

```go
package main

import "fmt"

var (
	pageIndex int
	pageSize  int
)

func init() {
	pageIndex = 1
	pageSize = 20
}

func main() {
	fmt.Printf("page index = %d\n", pageIndex)
	fmt.Printf("page size = %d\n", pageSize)
}

// $ go run main.go
// 输出如下 
/**
  page index = 1
  page size = 20
*/
```

### 多个包之间引用初始化顺序

```go
// 定义包 A
package A

func init() {
	println("hello A")
}
```

```go
// 定义包 B
package B

import "A" // 包 B 引用包 A

func init() {
	println("hello B")
}
```

```go
package main

import "B" // 包 main 引用包 B

func init() {
	println("hello main")
}

func main() {
	println("hello world")
}

// $ go run main.go
// 输出如下 
/**
  hello A
  hello B
  hello main
  hello world
*/
```

# reference

1. https://book.douban.com/subject/27044219/