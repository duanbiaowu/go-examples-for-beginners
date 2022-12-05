# 概述

建议先阅读 [range](range.md), [阻塞通道](channel.md), [非阻塞通道](channel_buffer.md)。

`range` 除了可以遍历字符串、切片、数组等数据结构外，还可以遍历通道。

# 语法规则

和遍历其他数据结构不同，遍历通道时没有 `索引` 的概念，只有值，语法如下:

```go
for v := range ch { // v 是从通道接收到的值
// do something ...
}
```

# 使用规则

1. **遍历一个空的通道 (值为 nil) 时，阻塞**
2. **遍历一个阻塞 && 未关闭的通道时，阻塞**
3. **遍历一个阻塞 && 已关闭的通道时，不做任何事情**
4. **遍历一个非阻塞 && 未关闭的通道时，就接收通道内的所有缓存数据，然后阻塞**
5. **遍历一个非阻塞 && 已关闭的通道时，就接收通道内的所有缓存数据，然后返回**

# 例子

## 遍历一个空的通道

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	var done chan bool

	go func() {
		for v := range done {
			fmt.Printf("v = %v\n", v)
			break
		}

		fmt.Println("range broken") // 执行不到这里
	}()

	time.Sleep(time.Second)
}

// $ go run main.go
// 没有任何输出
```

## 遍历一个阻塞 && 未关闭的通道

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for v := range done {
			fmt.Printf("v = %v\n", v)
			break
		}

		fmt.Println("range broken") // 执行不到这里
	}()

	time.Sleep(time.Second)
}

// $ go run main.go
// 没有任何输出
```

## 遍历一个阻塞 && 已关闭的通道

```go
package main

func main() {
	ch := make(chan string)
	close(ch)
	for v := range ch {
		println(v)
	}
}

// $ go run main.go
// 没有任何输出
```

## 遍历一个非阻塞 && 未关闭的通道

### 通道中无缓存数据，阻塞

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		for v := range ch {
			fmt.Printf("v = %v\n", v)
			break
		}

		fmt.Println("range broken") // 执行不到这里
	}()

	time.Sleep(time.Second)
}

// $ go run main.go
// 没有任何输出
```

### 通道中有 1 个数据

```go
package main

func main() {
	ch := make(chan string, 1)
	ch <- "hello world"

	for v := range ch {
		println(v)
	}

	// 代码执行不到这里
	close(ch)
}

// $ go run main.go
// 输出如下
/**
  v = hello world
*/
```

### 通道中有多个数据

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)
	for i := 0; i < 3; i++ {
		ch <- "hello world"
	}

	go func() {
		for v := range ch {
			fmt.Printf("v = %v\n", v)
		}

		fmt.Println("range broken") // 执行不到这里
	}()

	time.Sleep(time.Second)
}

// $ go run main.go
// 输出如下
/**
  v = hello world
  v = hello world
  v = hello world
*/
```

## 遍历一个非阻塞 && 已关闭的通道时

### 通道中无缓存数据，直接返回

```go
package main

func main() {
	ch := make(chan string, 1)
	close(ch)

	for v := range ch {
		println(v)
	}
}

// $ go run main.go
// 没有输出
```

### 通道中有 1 个数据

```go
package main

func main() {
	ch := make(chan string, 1)
	ch <- "hello world"
	close(ch)

	for v := range ch {
		println(v)
	}
}

// $ go run main.go
// 输出如下
/**
  hello world
*/
```

### 通道中有多个数据

```go
package main

func main() {
	ch := make(chan string, 3)
	for i := 0; i < 3; i++ {
		ch <- "hello world"
	}
	close(ch)

	for v := range ch {
		println(v)
	}
}

// $ go run main.go
// 输出如下
/**
  hello world
  hello world
  hello world
*/
```