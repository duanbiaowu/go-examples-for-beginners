# 概述
调用 `os` 包即可。

# 例子

## 文件是否存在
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("/tmp/not_found_main.go"); os.IsNotExist(err) {
		fmt.Printf("%s\n", err)
	}
}
// $ go run main.go
/**
    stat /tmp/not_found_main.go: no such file or directory
*/
```

## 文件是否拥有权限
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.ReadFile("/root/passwd")
	if err != nil && os.IsPermission(err) {
		fmt.Printf("%s\n", err)
	}
}
// $ go run main.go
// 输出如下 
/**
    open /root/passwd: permission denied
*/
```