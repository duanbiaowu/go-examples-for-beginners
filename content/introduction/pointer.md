# 概述

**Go 提供了指针操作，但是没有指针运算。**
也就是说，不像 C 语言中那般强大，毕竟 `指针是 C 语言的灵魂`。
即使如此，指针依然是非常重要的，在一些 `性能敏感` 的场景中，指针的身影随处可见。
如果是系统编程、操作系统或者网络应用，指针更是不可或缺的一部分。

**指针的值是一个变量的地址。当然了，指针也是变量的一种，但是一般称其为 `指针变量`。**

# 取地址

**关键字 `&` 表示取地址符**。

程序运行时，数据通常存储在内存中，每个内存块都有一个地址， 通常使用 `十六进制` 表示，比如 `0xc0000160a0`。

```shell
# 将 & 放到一个变量前，就会获得该变量对应的内存地址, 例如
x := 1024

# p 变量是一个指针变量，值对应着变量 x 的地址
p := &x
```

## 例子

```go
package main

import "fmt"

func main() {
	pi := 3.1415
	fmt.Printf("%p\n", &pi) // 直接取地址, 输出的是变量 pi 的地址

	var p *float64        // 浮点型指针变量 
	p = &pi               // 通过变量取地址
	fmt.Printf("%p\n", p) // 输出的是指针的地址, 输出的是指针 p 的地址
}

// $ go run main.go
// 输出如下 
/**
  0xc0000160a0    // 这个是我电脑的内存地址，你的输入可能和这个不一样
  0xc0000b2000
*/
```

# 改变值

在刚才的例子中，获取到了变量的地址后，直接进行了输出。 那么，应该如何输出指针对应的变量的值呢？

**关键字 `*` 表示指针调用符**。

```shell
# 将 * 放到一个指针变量前，就会获得该指针变量对应的变量的值, 例如
x := 1024

# p 变量是一个指针变量，值对应着变量 x 的地址
p := &x

# *p 表示 p 对应的变量的值，也就是 x 的值，也就是 1024，
# *p = 1025, 表示将 x 的值修改为 1025
*p = 1025   
```

## 例子

```go
package main

import "fmt"

func main() {
	ok := true
	var p *bool            // 布尔型指针变量
	p = &ok                // 获取 ok 的地址
	fmt.Printf("%t\n", *p) // 输出指针变量 p 对应的变量 ok 的值

	*p = false             // 改变了变量 ok 的值
	fmt.Printf("%t\n", *p) // 输出指针变量 p 对应的变量 ok 的值
}

// $ go run main.go
// 输出如下 
/**
    true
    false
*/
```

# 小结

本小节假定读者已经了解指针的概念，所以只是简单地介绍一下 Go 中指针的相关语法。
如果读者对指针的概念没有基础了解，推荐阅读下面的文章。

# 附录

## uintptr 是什么

![uintptr with unsafe.Pointer](https://github.com/duanbiaowu/go-examples-for-beginners/assets/10596220/fa27cf2f-57de-4b51-b8ea-ca104403bc28)


uintptr 是一个可以表示任何指针地址的整数。

1. 任何类型的指针和 unsafe.Pointer 可以相互转换。
2. uintptr 类型和 unsafe.Pointer 可以相互转换。

```go
package main

import (
	"math"
	"unsafe"
)

func main() {
	var p uintptr
	x := math.MaxFloat64
	p = uintptr(unsafe.Pointer(&x))
	println(unsafe.Sizeof(p)) // 8
}
```

# 扩展阅读

1. [指针 - 维基百科](https://zh.m.wikipedia.org/zh-cn/%E6%8C%87%E6%A8%99_(%E9%9B%BB%E8%85%A6%E7%A7%91%E5%AD%B8))
