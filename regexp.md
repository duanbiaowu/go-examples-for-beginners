# 概述
调用 `regexp` 包即可。

# 例子

## 是否匹配
```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, err := regexp.MatchString("h[a-z]+.*d$", "hello world")
	if err != nil {
		panic(err)
	}
	fmt.Println(match)

	match, err = regexp.MatchString("h[a-z]+.*d$", "ello world")
	if err != nil {
		panic(err)
	}
	fmt.Println(match)
}
// $ go run main.go
// 输出如下
/**
    true
    false
*/
```

## 匹配所有子字符串
```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	c, err := regexp.Compile("h[a-z]")
	if err != nil {
		panic(err)
	}

	res := c.FindAllString("hello world", -1)
	fmt.Printf("res = %v\n", res)

	res2 := c.FindAllString("hello world hi ha h1", -1)
	fmt.Printf("res2 = %v\n", res2)
}
// $ go run main.go
// 输出如下
/**
    res = [he]
    res2 = [he hi ha]
*/
```

## 替换所有子字符串
```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	c, err := regexp.Compile("h[a-z]")
	if err != nil {
		panic(err)
	}

	res := c.ReplaceAll([]byte("hello world"), []byte("?"))
	fmt.Printf("res = %s\n", res)

	res2 := c.ReplaceAll([]byte("hello world hi ha h1"), []byte("?"))
	fmt.Printf("res2 = %s\n", res2)
}
// $ go run main.go
// 输出如下
/**
    res = ?llo world
    res2 = ?llo world ? ? h1
*/
```

## 匹配中文
```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, err := regexp.MatchString("\\x{4e00}-\\x{9fa5}", "hello world")
	if err != nil {
		panic(err)
	}
	fmt.Println(match)

	match, err = regexp.MatchString("\\p{Han}+", "hello 世界")
	if err != nil {
		panic(err)
	}
	fmt.Println(match)
}
// $ go run main.go
// 输出如下
/**
    false
    true
*/
```