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
	file, err := os.Create("/tmp/test_main.go")
	if err != nil {
		panic(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Printf("file name is %s\n", file.Name())
}
// $ go run main.go
// 输出如下 
/**
    file name is /tmp/test_main.go
*/
```