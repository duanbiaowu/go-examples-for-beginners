# 概述

调用 `encoding/base64` 包即可。

# 例子

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "hello world"
	sEncode := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Printf("encode(`hello world`) = %s\n", sEncode)

	sDecode, err := base64.StdEncoding.DecodeString(sEncode)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("decode(`%s`) = %s\n", sEncode, sDecode)
	}
}

// $ go run main.go
// 输出如下
/**
  encode(`hello world`) = aGVsbG8gd29ybGQ=
  decode(`aGVsbG8gd29ybGQ=`) = hello world
*/
```