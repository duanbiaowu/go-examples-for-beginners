# 概述

调用 `path/filepath` 包即可。

# 例子

```go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs("./main.go")
	if err != nil {
		panic(err)
	}

	fmt.Printf("file abs path = %s\n", path)                   // 获取文件的绝对路径
	fmt.Printf("file name = %s\n", filepath.Base("./main.go")) // 获取文件名称
	fmt.Printf("file ext = %s\n", filepath.Ext("./main.go"))   // 获取文件扩展名

	path2 := filepath.Join("/tmp", "code", "test", "main.go")
	fmt.Printf("build file path = %s\n", path2) // 获取构建的文件路径
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样
/**
  file abs path = /home/codes/Go-examples-for-beginners/main.go
  file name = main.go
  file ext = .go
  build file path = /tmp/code/test/main.go
*/
```