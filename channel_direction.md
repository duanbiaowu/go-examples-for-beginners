# 概述
通道的方向分为 `发送` 和 `接收`。默认情况下，通道是双向的 (同时发送和接收)，但是可以通过标识符指明通道为单向 (只读或只写)。

# 语法规则
## 可读写通道 (同时支持发送和接收)
```shell
变量 := make(chan 数据类型)
# 例子
ch := make(chan string)
```

## 只读通道 (只支持接收)
```shell
变量 := make(<-chan 数据类型)
# 例子
ch := make(<-chan string)
```

## 只写通道 (只支持发送)
```shell
变量 := make(chan<- 数据类型)
# 例子
ch := make(chan<- string)
```

## 类型转换
双向通道可以转换为单向通道，但是单向通道无法转换为双向通道。

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