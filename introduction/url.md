# 概述

调用 `net/url` 包即可。

# 例子

## 构造 URL

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "go.dev"
	u.Path = "/learn/doc"

	values := u.Query()
	values.Add("hello", "world")

	u.RawQuery = values.Encode()

	fmt.Printf("URL = %s\n", u.String())
}

// $ go run main.go
// 输出如下
/**
  URL = https://go.dev/learn/doc?hello=world
*/
```

## 解析 URL

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://golang.org"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
}

// $ go run main.go
// 输出如下
/**
  https
  golang.org
*/
```

## 解析 URL (带参数)

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://go.dev/learn/doc?hello=world"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)

	fmt.Printf("Param hello = %s\n", u.Query().Get("hello"))
}

// $ go run main.go
// 输出如下
/**
  https
  go.dev
  /learn/doc
  Param hello = world
*/
```