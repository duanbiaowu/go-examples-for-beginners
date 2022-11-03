# 概述
调用 `os` 包即可。

# 例子

## 设置参数
```go
package main

import (
	"flag"
	"fmt"
)

var (
	name     = flag.String("name", "Tom", "Please input your name:")
	age      = flag.Int("age", 6, "Please input your age:")
	hasMoney = flag.Bool("hasMoney", true, "Do you have any money?")
)

func main() {
	flag.PrintDefaults()	// 打印参数提示信息
	flag.Parse()
    
	fmt.Printf("name is %s\n", *name)
	fmt.Printf("name is %d\n", *age)
	fmt.Printf("name is %t\n", *hasMoney)
}
// 默认参数
// $ go run main.go
// 输出如下 
/**
      -age int
            Please input your age: (default 6)
      -hasMoney
            Do you have any money? (default true)
      -name string
            Please input your name: (default "Tom")
    name is Tom
    name is 6
    name is true
*/


// 设置参数
// $ go run main.go -name=Jerry -age=8 -hasMoney=false
// 输出如下 
/**
      -age int
            Please input your age: (default 6)
      -hasMoney
            Do you have any money? (default true)
      -name string
            Please input your name: (default "Tom")
    name is Jerry
    name is 8
    name is false
*/
```