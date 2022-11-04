# 概述
建议先阅读 [阻塞通道](channel.md) 和 [非阻塞通道](channel_buffer.md)。
在前面的两个小节中， 为了最小化代码达到演示效果，省略了 `关闭通道` 的步骤，
正确的做法应该是在通道使用完成后关闭。

# 使用规则
通过关键字 `clsoe` 关闭通道。

1. **关闭一个空的通道 (值为 nil) 时，panic**
2. **关闭一个非空 && 已关闭的通道时，panic**
3. **关闭一个非空 && 未关闭的通道时，正常关闭**

这里的规则不必死记硬背，笔者遇到的大多数情况属于第二种，也就是 `重复关闭一个通道`，
读者做到实际开发中遇到 `关闭通道` 的场景时，联系上下文，**确认通道不会出现重复关闭的情况**即可。

# 例子

## 关闭一个空的通道
```go
package main

func main() {
	var ch chan bool
	close(ch)
}
// $ go run main.go
// 输出如下
/**
    panic: close of nil channel

    ...
    ...
    exit status 2
*/
```

## 关闭一个非空 && 已关闭通道
```go
package main

func main() {
	ch := make(chan bool)
	close(ch)
	close(ch) // 重复关闭
}
// $ go run main.go
// 输出如下
/**
    panic: close of nil channel
    
    ...
    ...
    exit status 2
*/
```

## 关闭一个非空 && 未关闭的通道
```go
package main

func main() {
	ch := make(chan bool)
	close(ch)
	println("channel closed")
}
// $ go run main.go
// 输出如下
/**
    channel closed
*/
```