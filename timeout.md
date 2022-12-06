# 概述

利用 `channel (通道)` 和 `time.After()` 方法实现超时控制。

# 例子

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func() {
		defer func() {
			ch <- true
		}()

		time.Sleep(2 * time.Second) // 模拟超时操作
	}()

	select {
	case <-ch:
		fmt.Println("ok")
	case <-time.After(time.Second):
		fmt.Println("timeout!")
	}
}

// $ go run main.go
// 输出如下
/**
  timeout!
*/
```