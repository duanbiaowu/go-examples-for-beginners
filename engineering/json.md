# 概述

在 [Go 快速入门指南 - JSON]() 讲解了 `JSON` 的常用方法，但是除此之外，`JSON` 还有一些鲜为人知的使用技巧，
可以简洁地组合和忽略结构体字段，避免了重新定义结构体和内嵌结构体等较为笨拙的方式，这在 `接口输出` 和 `第三发接口对接` 业务场景中非常有帮助。
这篇做一个补充，两篇文章涉及到的 `JSON` 知识点，应该足够大部分开发场景的使用了。

# 例子

## 临时忽略某个字段

比如在接口中输出用户信息时，希望过滤掉密码字段。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	u := &User{
		UserName: "root",
		Email:    "root@gmail.com",
		Password: "123456",
	}

	data, err := json.Marshal(struct {
		*User
		// 使用一个内嵌的字段覆盖掉原字段
		Password string `json:"password,omitempty"`
	}{
		User: u,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
}

// $ go run main.go
// 输出如下 
/**
  {"userName":"root","email":"root@gmail.com"}
*/
```

## 临时添加字段

比如在接口中输出用户信息时，希望添加一个 Token 字段。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	u := &User{
		UserName: "root",
		Email:    "root@gmail.com",
		Password: "123456",
	}

	data, err := json.Marshal(struct {
		*User
		// 使用一个内嵌的字段覆盖掉原字段
		Password string `json:"password,omitempty"`
		// 新增一个字段
		Token    string `json:"token"`
	}{
		User:  u,
		Token: "askdhfh2oyy43423#14$$asdssxxx11",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
}

// $ go run main.go
// 输出如下 
/**
  {"userName":"root","email":"root@gmail.com","token":"askdhfh2oyy43423#14$$asdssxxx11"}
*/
```

## 字符串和数字转换

接口对接时，可能会存在双方字段名称一样，但是类型不一样的的情况。比如同一个字段，A 方用 `int` 类型，
B 方用 `string` 类型，下面的例子演示如何解决这个 `数据类型冲突问题` 。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// 字段是 int 类型， JSON 输出 string 类型
	Age      int    `json:"age,string"`
}

func main() {
	u := &User{
		UserName: "root",
		Email:    "root@gmail.com",
		Password: "123456",
		Age:      100,
	}

	data, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
}

// $ go run main.go
// 输出如下 
/**
  {"userName":"root","email":"root@gmail.com","password":"123456","age":"100"}
*/
```
