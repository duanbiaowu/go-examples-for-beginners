# 概述
建议先阅读 [函数](func.md) 和 [接口](interface.md)。

# 例子

## errors.New() 创建错误

```go
package main

import (
	"errors"
	"fmt"
)

// 自定义除法函数
func myDivide(dividend, divisor float64) (float64, error) {
	if divisor == 0 { // 除数不能为 0
		return 0, errors.New("divide by zero") // 返回一个错误
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
// 输出如下，你的输出可能和这里的不一样
/**
    Error: divide by zero
*/
```

## fmt.Errorf() 创建错误

不同于 `errors.New()`, `fmt.Errorf()` 在构建错误时，可以进行格式化。

```go
package main

import (
	"fmt"
)

// 自定义除法函数
func myDivide(dividend, divisor float64) (float64, error) {
	if divisor == 0 { // 除数不能为 0
		return 0, fmt.Errorf("%.2f divide by zero", dividend) // 返回一个错误
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
// 输出如下，你的输出可能和这里的不一样
/**
    Error: 100.00 divide by zero
*/
```

# 最佳实践
**永远不要忽略错误，否则可能会导致程序崩溃！**