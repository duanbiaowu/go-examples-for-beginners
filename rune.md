# 概述
Go 中表示字符的关键字为 `rune`, 也就是 `int32`。

# 语法规则
由单引号 `'` 括起来，只能包含一个字符

# 字符串长度
关于字符串不同编码对长度的计算方式，感兴趣的读者可以参考扩展阅读。

## 例子
```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := 'a'
	fmt.Printf("s type = %T, len = %d\n", s, utf8.RuneLen(s))

	s2 := '我'
	fmt.Printf("s2 type = %T, len = %d\n", s, utf8.RuneLen(s2))
}
// $ go run main.go
// 输出如下
/**
    s type = int32, len = 1
    s2 type = int32, len = 3
*/
```

# 扩展阅读
1. [十分钟搞清字符集和字符编码](http://cenalulu.github.io/linux/character-encoding/)