# 概述
**Go 箴言: 不要通过共享内存来通信，而要通过通信来共享内存。**

# 语法规则

## 阻塞通道与非阻塞通道
通过关键字 `chan` + `数据类型` 来表明通道数据类型，调用 `make()` 函数来初始化一个通道。
`make()` 函数的第二个参数为通道长度，如果未指定或指定为 0，则该通道为非缓存通道 (阻塞通道), 
否则该通道为缓存通道 ([非阻塞通道](channel_buffer.md))。

## 向通道发送数据
`通道变量 <- 数据`

## 从通道接收数据
`<- 通道变量`

## 发送和接收必须配对

## 死锁问题

# 例子

## 搭配 goroutine
```go
package main

func main() {
	ch := make(chan string) // 没有设置通道的长度

	go func() {
		ch <- "hello world"
	}()

	msg := <-ch // 一直阻塞，直到接收到通道消息
	println(msg)
}
// $ go run main.go
// 输出如下
/**
    hello world
*/
```

## 死锁
```go
package main

func main() {
	ch := make(chan string) // 没有设置通道的长度

	ch <- "hello world" // 向通道发送数据，但是没有接收者

	msg := <-ch // 代码执行不到这里, 因为上面阻塞发送数据时，就已经死锁了
	println(msg)
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
