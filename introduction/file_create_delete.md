# 概述

调用 `os` 包即可。

# 例子

## 创建文件

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

	// 记得关闭文件句柄
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

## 删除文件

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("/tmp/test_main.go.bak")
	if err != nil {
		panic(err)
	}

	// 记得关闭文件句柄
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Printf("file name is %s\n", file.Name())

	err = os.Remove("/tmp/test_main.go.bak")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s has been deleted\n", file.Name())
	}
}

// $ go run main.go
// 输出如下 
/**
  file name is /tmp/test_main.go.bak
  /tmp/test_main.go.bak has been deleted
*/
```