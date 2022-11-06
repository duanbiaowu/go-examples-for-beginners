# 概述
建议先阅读 [阻塞通道](channel.md), [非阻塞通道](channel_buffer.md), [关闭通道](channel_close.md),
[通道方向](channel_direction.md)。

Go 语言没有提供函数或方法判断一个通道是否关闭。因此只能使用一个变通的办法：接收通道元素，根据返回的布尔值确定通道是否关闭。

# 例子

## 双向通道检测
```go
package main

func main() {
	ch := make(chan string)
	close(ch)

	if _, open := <-ch; !open {
		println("channel closed")
	}
}
// $ go run main.go
// 输出如下
/**
    channel closed
*/
```

## 单向 (只读) 通道检测
```go
package main

import "time"

func main() {
	ch := make(chan string)

	go func(c <-chan string) {
		if _, open := <-c; !open {
			println("channel closed")
		}
	}(ch)

	close(ch)
	time.Sleep(time.Second)
}
// $ go run main.go
// 输出如下
/**
    channel closed
*/
```

## 单向 (只写) 通道检测
对于只写通道，需要采用一个折中的办法: 
* 尝试向通道写入数据
  * 如果写入成功，说明通道还未关闭
  * 写入失败，发生 `panic`, 这时可以利用 `defer` 在 `recover` 中输出原因

```go
package main

import "time"

func main() {
	ch := make(chan string)

	go func(c chan<- string) {
		defer func() {
			if err := recover(); err != nil { // 捕获到 panic
				println("channel closed")
			}
		}()

		c <- "hello world"
	}(ch)

	close(ch)
	time.Sleep(time.Second)
}
// $ go run main.go
// 输出如下，你的输出可能和这里的不一样
/**
    channel closed
*/
```