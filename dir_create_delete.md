# 概述
调用 `os` 包即可。

# 例子

## 目录创建

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("/tmp/test_go_main_dir", 0755) // 创建 1 级目录
	if err != nil {
		panic(err)
	} else {
		fmt.Println("/tmp/test_go_main_dir has been created")
	}

	err = os.MkdirAll("/tmp/test_go_main_dir/1/2/3", 0755) // 创建多级目录
	if err != nil {
		panic(err)
	} else {
		fmt.Println("/tmp/test_go_main_dir/1/2/3 has been created")
	}
}
// $ go run main.go
// 输出如下
/**
    /tmp/test_go_main_dir has been created
    /tmp/test_go_main_dir/1/2/3 has been created
*/
```

## 目录删除
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.RemoveAll("/tmp/test_go_main_dir")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("/tmp/test_go_main_dir has been deleted")
	}
}
// $ go run main.go
// 输出如下
/**
    /tmp/test_go_main_dir/1/2/3 has been deleted
*/
```