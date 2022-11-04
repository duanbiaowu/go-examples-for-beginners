# 概述
Go 中 `字符串` 语义和其他编程语言中的字符串中一样，有一点不同的地方在于: **Go 中字符串值无法改变**，可以理解为：
一旦完成定义之后，字符串就是一个 [常量](const.md)。

# 解释型字符串
使用双引号 `"` 括起来，其中的转义符将被替换，转义符包括:
* `\n`: 换行符 
* `\r`: 回车符 
* `\t`: tab 键 
* `\u`: 或 \U：Unicode 字符 
* `\\`: 反斜杠自身
* ...
* ...

## 示例
```go
package main

func main() {
	s := "hello\nworld"
	println(s)
}
// $ go run main.go
// 输出如下
/**
    hello
    world
*/
```

# 非解释型字符串
使用反引号 **`** 括起来，其中的转义符会被原样输出。

## 示例
```go
package main

func main() {
	s := `hello world \n`
	println(s)
}
// $ go run main.go
// 输出如下
/**
    hello world \n
*/
```

# 字符串长度
关于字符串不同编码对长度的计算方式，感兴趣的读者可以参考扩展阅读。

## ASCII
```go
package main

func main() {
	s := "hello world"
	println(len(s))
}
// $ go run main.go
// 输出如下
/**
    11
*/
```

### 中文计算方式存在的问题
中文会按照 3 个字节计算，这在一些场景下 (统计时英文和中文都算一个字)，可能并不适合。
```go
package main

func main() {
	s := "hello 世界"
	println(len(s))
}
// $ go run main.go
// 输出如下
/**
    12
*/
```

## UTF-8
```go
package main

import "unicode/utf8"

func main() {
	s := "hello world"
	println(utf8.RuneCountInString(s))

	s2 := "hello 世界"
	println(utf8.RuneCountInString(s2)) // 每个中文文字算作一个字符
}
// $ go run main.go
// 输出如下
/**
    11
    8
*/
```

# 子字符串
## 语法规则
```shell
s[i:j]
```
s 代表字符串变量名称， i 和 j 的取值范围是 0 到 `len(s)` (字符串长度)。如果两个参数都不设置，则取全部字符。

## 例子
```go
package main

import "fmt"

func main() {
	s := "hello world"

	s2 := s[3:5] // 取第 3 个到第 5 个字符
	fmt.Printf("s2 = %v, type = %T\n", s2, s2)

	s3 := s[:] // 取全部字符
	fmt.Printf("s3 = %v, type = %T\n", s3, s3)
}
// $ go run main.go
// 输出如下
/**
    s2 = lo, type = string
    s3 = hello world, type = string
*/
```

# 字符串遍历
**为了处理不同编码字符串，请使用 `range` 。**

## 例子
注意区分不同编码对遍历结果的影响。

```go
package main

import "fmt"

func main() {
	str := "hello 世界"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c  ", str[i])
	}
	fmt.Println()

	for _, s := range str {
		fmt.Printf("%c  ", s)
	}
	fmt.Println()
}
// $ go run main.go
// 输出如下
/**
    h  e  l  l  o     ä  ¸     ç      
    h  e  l  l  o     世  界
*/
```

# 字符串拼接
## 使用 `+`
```go
package main

func main() {
	s := "hello world"
	s += " !"
	println(s)
}
// $ go run main.go
// 输出如下
/**
    hello world !
*/
```

## 使用 fmt.Sprintf() 函数
```go
package main

import "fmt"

func main() {
	s := "hello world"
	s2 := fmt.Sprintf("%s %s", s, " !")
	println(s2)
}
// $ go run main.go
// 输出如下
/**
    hello world !
*/
```

# 扩展阅读
1. [十分钟搞清字符集和字符编码](http://cenalulu.github.io/linux/character-encoding/)