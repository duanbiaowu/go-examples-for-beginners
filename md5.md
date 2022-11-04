# 概述
调用 `crypto/md5` 包即可。

# 例子

```go
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	s := "hello world"
	h := md5.New()
	h.Write([]byte(s))
	res := h.Sum(nil)
	fmt.Printf("md5(`hello world`) = %x\n", res)
}
// $ go run main.go
// 输出如下
/**
    md5(`hello world`) = 5eb63bbbe01eeed093cb22bb8f5acdc3
*/
```