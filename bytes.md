# 概述
建议先阅读 [字符串](string.md), [切片](slice.md)。

由于字符串不可变，如果每次以 `重新赋值` 的方式改变字符串，效率会非常低，这时应该使用 `[]byte` 类型，[]byte 元素可以被修改。

因为 `byte` 类型是 `uint8` 类型的别名，所以 `[]byte` 也就是 `[]uint8`。

# 语法规则

## 字符串转化为 []byte
```go
package main

import "fmt"

func main() {
	s := "hello world"
	b := []byte(s)
	fmt.Printf("b type = %T, val = %s\n", b, b)
}
// $ go run main.go
// 输出如下
/**
    b type = []uint8, val = hello world
*/
```

## []byte 转换为字符串
```go
package main

import "fmt"

func main() {
	b := []byte{'h', 'e', 'l', 'l', '0', ' ', 'w', 'o', 'r', 'l', 'd'}
	s := string(b)
	fmt.Printf("s type = %T, val = %s\n", s, s)
}
// $ go run main.go
// 输出如下
/**
    s type = string, val = hell0 world
*/
```

# 长度计算
关于字符串不同编码对长度的计算方式，感兴趣的读者可以参考扩展阅读。

## ASCII
```go
package main

import "fmt"

func main() {
	b := []byte{'h', 'i'}
	fmt.Printf("b length = %d\n", len(b))
}
// $ go run main.go
// 输出如下
/**
    b length = 2
*/
```

## 中文算作 3 个字符
```go
package main

import "fmt"

func main() {
	b := []byte("我是")
	fmt.Printf("b length = %d\n", len(b))
}
// $ go run main.go
// 输出如下
/**
    b length = 6
*/
```

## 中文算作 1 个字符
```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("我是")
	fmt.Printf("b length = %d\n", utf8.RuneCount(b))
}
// $ go run main.go
// 输出如下
/**
    b length = 2
*/
```

# 扩展阅读
1. [十分钟搞清字符集和字符编码](http://cenalulu.github.io/linux/character-encoding/)