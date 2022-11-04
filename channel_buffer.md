# 概述
**Go 箴言: 不要通过共享内存来通信，而要通过通信来共享内存。**

建议先阅读 [阻塞通道](channel.md)。

# 语法规则

## 阻塞通道与非阻塞通道
通过关键字 `chan` + `数据类型` 来表明通道数据类型，调用 `make()` 函数来初始化一个通道。
`make()` 函数的第二个参数为通道长度，如果未指定或指定为 0，则该通道为非缓存通道 ([阻塞通道](channel.md)),
否则该通道为缓存通道 (非阻塞通道)。

# 例子

## 避免死锁
```go
package main

func main() {
	ch := make(chan string, 2)

	ch <- "hello" // 不会死锁，因为 ch 是缓冲通道
	ch <- "world"

	println(<-ch)
	println(<-ch)
}
// $ go run main.go
// 输出如下
/**
    hello world
*/
```