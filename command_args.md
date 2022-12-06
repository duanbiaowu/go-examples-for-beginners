# 概述

调用 `os` 包即可。

# 例子

## 获取参数个数, 遍历参数

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Number of args is %d\n\n", len(os.Args))

	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}

// $ go build main.go
// $ ./main -a -b --c -d
// 输出如下 
/**
  Number of args is 5

  ./main
  -a
  -b
  --c
  -d
*/
```