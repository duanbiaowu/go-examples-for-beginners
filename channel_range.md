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
1. **遍历一个空的通道 (值为 nil) 时，panic**
2. **遍历一个阻塞 && 未关闭的通道时，panic**
3. **遍历一个阻塞 && 已关闭的通道时，不做任何事情**
4. **遍历一个非阻塞 && 未关闭的通道时，就接收通道内的所有缓存数据，然后 panic**
5. **遍历一个非阻塞 && 已关闭的通道时，就接收通道内的所有缓存数据，然后返回**

# 例子

## 遍历一个空的通道
```go
package main

func main() {
	var ch chan string
	for v := range ch {
		println(v)
	}
}
// $ go run main.go
// 输出如下
/**
    fatal error: all goroutines are asleep - deadlock!
    
    ...
    ...

    exit status 2
*/
```

## 遍历一个阻塞 && 未关闭的通道
```go
package main

func main() {
	ch := make(chan string)
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
  fatal error: all goroutines are asleep - deadlock!

  ...
  ...

  exit status 2
*/
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

### 通道中无缓存数据，直接 panic
```go
package main

func main() {
	ch := make(chan string, 1)
	for v := range ch {
		println(v)
	}
	// 代码执行不到这里
	close(ch)
}
// $ go run main.go
// 输出如下
/**
    fatal error: all goroutines are asleep - deadlock!
    
    ...
    ...
    
    exit status 2
*/
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
  hello world   // 输出 1 个数据
  fatal error: all goroutines are asleep - deadlock!

  ...
  ...

  exit status 2
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

	for v := range ch {
		println(v)
	}

	// 代码执行不到这里
	close(ch)
}
// $ go run main.go
// 输出如下
/**
  hello world 
  hello world 
  hello world 
  fatal error: all goroutines are asleep - deadlock!

  ...
  ...

  exit status 2
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