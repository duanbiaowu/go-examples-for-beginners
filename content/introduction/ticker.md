# 概述

调用 `time.NewTicker()` 方法即可。

# 例子

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second) // 模拟耗时操作
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-ticker.C:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样
/**
  2021-01-03 15:40:21
  2021-01-03 15:40:22
  2021-01-03 15:40:23
  2021-01-03 15:40:24
  2021-01-03 15:40:25
  Done!
*/
```