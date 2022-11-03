# 概述
调用 `os` 包，自定义 `status code`。

# 例子

```go
package main

import "os"

func main() {
	defer func() {
		println("exiting ...")  // 不会执行到这里
	}()

	os.Exit(3)
}
// $ go run main.go
// 输出如下 
/**
    exit status 3
*/
```