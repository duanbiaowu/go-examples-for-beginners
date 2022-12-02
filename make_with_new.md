# 两者区别

- new(T) 为数据类型 T 分配一块内存，初始化为类型 T 的零值，返回类型为指向数据的指针，可以用于所有数据类型
- make(T) 除了为数据类型 T 分配内存外，还可以指定长度和容量，返回类型为数据的初始化结构，只限于 `切片`, `Map`, `通道`

<p align="center">
<img width="600" src="./images/make_with_new.png">
</p>

# make

## 什么时候用?

声明并初始化 [切片](slice.md), [Map](map.md), 通道(后面会讲到) 

## 为什么定义

为什么专门针对切片, Map 和 通道类型定义一个 `make` 函数呢？
因为这 3 种数据类型要求使用时必须完成初始化，未初始化就使用可能会引发错误，具体规则如下:

- 未初始化的切片值为 `nil`, 如果直接获取或设置元素数据会报错
- 未初始化的 `Map` 值为 `nil`, 如果直接设置元素数据会报错
- 未初始化的 `通道` 值为 `nil`, 发送数据和接收数据会阻塞 (详情在后面通道章节介绍)

### 未初始化的切片

```go
package main

func main() {
	var s []int

	// 直接获取值: 报错
	_ = s[0]

	// 直接设置值: 同样报错
	//s[0] = 100
}
// $ go run main.go
// 输出如下 
/**
panic: runtime error: index out of range [0] with length 0
...
exit status 2
*/
```

### 未初始化的 Map

```go
package main

func main() {
	var m map[int]string

	// 直接设置值: 报错
	m[100] = "hello world"
}
// $ go run main.go
// 输出如下 
/**
panic: panic: assignment to entry in nil map
...
exit status 2
*/
```

### append()

为什么切片即使是 `nil`, 却可以调用 `append()` 函数呢? 因为 `append()` 函数内部实现中做了兼容，如果切片为 `nil`，
那么会先申请好需要的内存空间，然后在复制给切片，等于 `覆盖掉原来的切片`，这样就不会报错了。

## 使用 new 初始化切片和 Map

如果我们不使用 `make()` 函数创建切片和 `Map` 可以吗？ 当然是可以的， `new()` 函数可以创建任何数据类型，当然也包括切片和 `Map`，
但是 `new()` 函数返回的指针指向的是类型的 `零值`，对于切片和 `Map` 来说， 零值依然是 `nil`, 这又回到了上面的问题。

### new() 创建的切片

```go
package main

import "fmt"

func main() {
	s := new([]int)

	fmt.Printf("s type = %T, val = %#v\n", *s, *s)

	// 直接获取值: 报错
	_ = (*s)[0]

	// 直接设置值: 同样报错
	//s[0] = 100
}
// $ go run main.go
// 输出如下 
/**
s type = []int, val = []int(nil)
panic: runtime error: index out of range [0] with length 0
...
exit status 2
*/
```

### new() 创建的 Map

```go
package main

import "fmt"

func main() {
	m := new(map[int]string)

	fmt.Printf("s type = %T, val = %#v\n", *m, *m)

	// 直接设置值: 报错
	(*m)[100] = "hello world"
}
// $ go run main.go
// 输出如下 
/**
m type = map[int]string, val = map[int]string(nil)
panic: assignment to entry in nil map
...
exit status 2
*/
```

# new

## 什么时候用?

除了 [切片](slice.md), [Map](map.md), 通道(后面会讲到) 以外的其他数据类型。
