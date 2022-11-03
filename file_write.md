# 概述
调用 `os` 包即可。

# 例子

## WriteFile() 直接写入

```go
package main

import "os"

func main() {
	code := `
package main

func main() {
	println("hello world")
}
`

	err := os.WriteFile("/tmp/test_main.go", []byte(code), 0755)
	if err != nil {
		panic(err)
	}
}
// $ go run main.go
// cat /tmp/test_main.go
// 输出如下 
/**
    package main

    func main() {
            println("hello world")
    }
*/
```

## 先获取文件句柄，然后写入
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/tmp/test_main.go", os.O_RDWR, 0755)
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

	code := `
package main

func main() {
	println("hello world")
}
`
	n, err := file.WriteString(code)
	if err != nil {
		panic(err)
	}

	err = file.Sync() // 同步到硬盘
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d characters were successfully written\n", n)
}
// $ go run main.go
// 输出如下 
/**
    55 characters were successfully written
*/
```