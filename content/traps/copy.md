---
date: 2023-01-01
title: Go 陷阱之 copy 函数复制失败
modify: 2023-01-01
---

# 概述

`copy` 函数可以将一个切片里面的元素拷贝至另外一个切片，函数的原型如下:

```go
func copy(dst []Type, src []Type) int
```

将切片 `src` 里面的元素拷贝至切片 `dst`, 返回拷贝成功的元素数量。需要注意的一点是，`copy` 函数默认切片 `dst` 有足够的容量存放拷贝的元素，
如果容量不足的话，那么切片 `src` 中超过 `dst` 容量长度的元素将不再拷贝。

# 错误的做法

```go
package main

import "fmt"

func main() {
	var src, dst []int
	src = []int{1, 2, 3}
	n := copy(dst, src)

	fmt.Printf("the number of copied elements is %d\n", n)
	fmt.Printf("dst = %v\n", dst)
}
```

```shell
$ go run main.go
# 输出如下

the number of copied elements is 0
dst = []
```

从输出结果中看到，返回拷贝成功的元素数量为 0, 变量 `dst` 依然是一个空切片，**错误的原因在于: 变量 `dst` 没有容量来存放变量 `src` 的元素**。
接下来，我们修正这个错误。

# 正确的做法

```go
package main

import "fmt"

func main() {
	var src, dst []int
	src = []int{1, 2, 3}
	dst = make([]int, len(src)) // 提前初始化 dst 容量
	n := copy(dst, src)

	fmt.Printf("the number of copied elements is %d\n", n)
	fmt.Printf("dst = %v\n", dst)
}
```

```shell
$ go run main.go
# 输出如下

the number of copied elements is 3
dst = [1 2 3]
```

从输出结果中看到，返回拷贝成功的元素数量为 3, 变量 `src` 的元素已经全部拷贝到 `dst` 里面。和刚才 `错误的例子` 不同，
我们在变量 `dst` 初始化完成后指定了容量为 3，这样正好可以存放变量 `src` 所有的元素。  