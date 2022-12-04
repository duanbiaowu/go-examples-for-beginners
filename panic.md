# 概述

`panic` 会终止程序并退出，因此只有在发生严重的错误时才会使用 `panic`。

# 例子

## 主动触发

```go
package main

func main() {
	panic("some error...")
}

// $ go run main.go
// 输出如下 
/**
    panic: some error...

goroutine 1 [running]:
main.main()
        /home/codes/Go-examples-for-beginners/main.go:4 +0x27
exit status 2
*/
```

## 除 0

```go
package main

import "fmt"

func main() {
	fmt.Println("除数不能为 0")

	n := 0
	fmt.Printf("5 / 0 = %d", 5/n)
}

// $ go run main.go
// 输出如下 
/**
  除数不能为 0
  panic: runtime error: integer divide by zero

  goroutine 1 [running]:
  main.main()
          /home/codes/Go-examples-for-beginners/main.go:15 +0x57
  exit status 2
*/
```