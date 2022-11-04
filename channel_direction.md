# 概述
通道的方向分为 `发送` 和 `接收`。

# 例子

```go
package main

// 参数是一个写入通道
func ping(pings chan<- string) {
	//<-pings					// 错误: pings 通道只能写入
	pings <- "hello world"
}

func pong(pings <-chan string, pongs chan<- string) {
	//pings <- "hello world"	// 错误: pings 通道只能读取
	//<-pongs 					// 错误: pongs 通道只能写入

	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string)
	pongs := make(chan string)
	done := make(chan bool)

	go ping(pings)
	go pong(pings, pongs)

	go func() {
		msg := <-pongs
		println(msg)
		done <- true
	}()

	<-done

	close(pings)
	close(pongs)
	close(done)
}
// $ go run main.go
// 输出如下
/**
    hello world
*/
```