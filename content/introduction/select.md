# 概述

`select` 类似 `switch`, 包含一系列逻辑分支和一个可选的默认分支。每一个分支对应通道上的一次操作 (发送或接收)，
**可以将 `select` 理解为专门针对通道操作的 `switch` 语句**。

# 语法规则

```go
select {
case v1 := <- ch1:
// do something ...  
case v2 := <- ch2:
// do something ...
default:
// do something ...
}
```

## 执行顺序

* 当同时存在多个满足条件的通道时，随机选择一个执行
* 如果没有满足条件的通道时，检测是否存在 default 分支
    * 如果存在则执行
    * 否则阻塞等待

通常情况下，把含有 `default 分支` 的 `select` 操作称为 `无阻塞通道操作`。

# 例子

## 随机执行一个

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)

	go func() {
		ch1 <- "hello"
	}()

	go func() {
		ch2 <- "world"
	}()

	go func() {
		done <- true
	}()

	time.Sleep(time.Second) //  休眠 1 秒

	// 此时 3 个通道应该都满足条件，select 会随机选择一个执行
	select {
	case msg := <-ch1:
		fmt.Printf("ch1 msg = %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("ch2 msg = %s\n", msg)
	case <-done:
		fmt.Println("done !")
	}

	close(ch1)
	close(ch2)
	close(done)
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样, 多运行几次看看效果
/**
  ch1 msg = hello
*/
```

## default (无阻塞通道操作)

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "hello"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "world"
	}()

	go func() {
		time.Sleep(time.Second)
		done <- true
	}()

	// 此时 3 个通道都在休眠中, 不满足条件，select 会执行 default 分支
	select {
	case msg := <-ch1:
		fmt.Printf("ch1 msg = %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("ch2 msg = %s\n", msg)
	case <-done:
		fmt.Println("done !")
	default:
		fmt.Println("default !")
	}

	close(ch1)
	close(ch2)
	close(done)
}

// $ go run main.go
// 输出如下
/**
  default !
*/
```

## 和 for 搭配使用

通过在 `select` 外层加一个 `for` 循环，可以达到 `无限轮询` 的效果。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)

	go func() {
		// ch1 goroutine 输出 1 次 
		fmt.Println("[ch1 goroutine]")
		time.Sleep(time.Second)
		ch1 <- "hello"
	}()

	go func() {
		// ch2 goroutine 输出 2 次
		for i := 0; i < 2; i++ {
			fmt.Println("[ch2 goroutine]")
			time.Sleep(time.Second)
		}
		ch2 <- "world"
	}()

	go func() {
		// done goroutine 输出 3 次
		for i := 0; i < 3; i++ {
			fmt.Println("[done goroutine]")
			time.Sleep(time.Second)
		}
		done <- true
	}()

	for exit := true; exit; {
		select {
		case msg := <-ch1:
			fmt.Printf("ch1 msg = %s\n", msg)
		case msg := <-ch2:
			fmt.Printf("ch2 msg = %s\n", msg)
		case <-done:
			fmt.Println("done !")
			exit = false // 通过变量控制外层 for 循环退出
		}
	}

	close(ch1)
	close(ch2)
	close(done)
}

// $ go run main.go
// 输出如下，你的输出顺序可能和这里的不一样
/**
  [done goroutine]
  [ch2 goroutine]
  [ch1 goroutine]
  ch1 msg = hello
  [done goroutine]
  [ch2 goroutine]
  ch2 msg = world
  [done goroutine]
  done !
*/
```

从输出结果看，`[ch1 goroutine]` 输出了 1 次，`[ch2 goroutine]` 输出了 2 次，`[done goroutine]` 输出了 3 次。 

#  附录

## select 和 switch 区别

select 只能应用于 channel 的操作，既可以用于 channel 的数据接收，也可以用于 channel 的数据发送。 如果 select 的多个分支都满足条件，则会随机的选取其中一个满足条件的分支。

switch 可以为各种类型进行分支操作， 设置可以为接口类型进行分支判断 (通过 i.(type))。switch 分支是顺序执行的，这和 select 不同。

## select 设置优先级

当 `ch1` 和 `ch2` 同时达到就绪状态时，优先执行任务1，在没有任务1的时候再去执行任务2。

```go
func worker2(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}
```
