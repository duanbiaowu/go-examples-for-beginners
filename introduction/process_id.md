# 概述

调用 `os` 包即可。

# 例子

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Process ID = %d\n", os.Getpid())
	fmt.Printf("Parent process ID = %d\n", os.Getppid())
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样
/**
  Process ID = 13962
  Parent process ID =  13860
*/
```