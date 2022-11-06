# 概述
`error` 本质上就是一个接口，原型如下:

```go
package builtin

type error interface {
	Error() string
}
```

# 例子

## 实现 error 接口
```go
package main

import (
	"fmt"
)

// 自定义错误结构体
type divideError struct {
	msg string
}

// 实现 error 接口
func (d *divideError) Error() string {
	return d.msg
}

func newDivideError() *divideError {
	return &divideError{
		msg: "divide by zero",
	}
}

// 自定义除法函数
func myDivide(dividend, divisor float64) (float64, error) {
	if divisor == 0 { // 除数不能为 0
		return 0, newDivideError() // 返回一个自定义错误
	}
	return dividend / divisor, nil
}

func main() {
	divide, err := myDivide(100, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err) // 输出错误信息
	} else {
		fmt.Printf("100 / 0 = %.2f\n", divide) // 代码执行不到这里
	}
}
// $ go run main.go
// 输出如下
/**
    Error: divide by zero
*/
```