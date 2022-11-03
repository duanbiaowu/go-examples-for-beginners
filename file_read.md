# 概述
调用 `os` 包即可。建议先阅读 [创建文件](file_create.md) 和 [写文件](file_write.md)。

# 例子

## 直接读取

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	code, err := os.ReadFile("/tmp/test_main.go")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", code)
}
// $ go run main.go
// 输出如下 
/**
    package main

    func main() {
        println("hello world")
    }
*/
```

## 先获取文件句柄，然后读取
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/tmp/test_main.go", os.O_RDONLY, 0755)
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

	code := make([]byte, 1024) // 注意: 切片的长度决定了读取内容的长度
	n, err := file.Read(code)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d characters were successfully read\n", n)
	fmt.Printf("%s\n", code)
}
// $ go run main.go
// 输出如下 
/**
    55 characters were successfully read

    package main
    
    func main() {
        println("hello world")
    }
*/
```