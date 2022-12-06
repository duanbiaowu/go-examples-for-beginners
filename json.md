# 概述

`encoding/json` 包含了 JSON 相关处理方法。

# 例子

## 结构体转为 JSON 字符串

调用 `json.Marshal()` 方法完成。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	addr string `json:"addr"` // 该属性转 JSON 时会被忽略
}

func main() {
	tom := person{ // 使用字面量初始化
		Name: "Tom",
		Age:  6,
		addr: "???",
	}

	tomJson, err := json.Marshal(tom)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json.Marshal(tom) = %s\n", tomJson) // 从输出字符串中可以看到，并没有 addr 属性
}

// $ go run main.go
// 输出如下 
/**
  json.Marshal(tom) = {"name":"Tom","age":6}
*/
```

## JSON 字符串转为结构体

调用 `json.Unmarshal()` 方法完成。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	addr string `json:"addr"` // 该属性转 json 时会被忽略
}

func main() {
	// 注意: JSON 字符串格式一定要正确，否则会报错
	tomJson := `
{
  "name": "Tom",
  "age": 6,
  "addr": "???"
}
`
	var tom person
	err := json.Unmarshal([]byte(tomJson), &tom)
	if err != nil {
		panic(err)
	}

	// 从输出字符串中可以看到，并没有为 addr 属性赋值
	fmt.Printf("Tom's name is %s, age is %d, addr is %s\n", tom.Name, tom.Age, tom.addr)
}

// $ go run main.go
// 输出如下 
/**
  Tom's name is Tom, age is 6, addr is
*/
```

## 输出格式化 JSON 字符串

调用 `json.MarshalIndent()` 方法完成。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby"`
}

func main() {
	tom := person{ // 使用字面量初始化
		Name: "Tom",
		Age:  6,
		Hobby: []string{
			"reading",
			"coding",
			"movie",
		},
	}

	// 前缀符为空字符串，缩进符为两个空格
	formatted, err := json.MarshalIndent(tom, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("json.MarshalIndent(tom) = \n%s\n", formatted)
}

// $ go run main.go
// 输出如下 
/**
  json.MarshalIndent(tom) =
  {
    "name": "Tom",
    "age": 6,
    "hobby": [
      "reading",
      "coding",
      "movie"
    ]
  }
*/
```

## 忽略零值

通过 `omitempty` 关键字完成。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby,omitempty"` // omitempty 关键字将字段标记为忽略零值
}

func main() {
	tom := person{ // 使用字面量初始化
		Name: "Tom",
		Age:  6,
	}

	// 前缀符为空字符串，缩进符为两个空格
	formatted, err := json.MarshalIndent(tom, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("json.MarshalIndent(tom) = \n%s\n", formatted)
}

// $ go run main.go
// 输出如下 
/**
  json.MarshalIndent(tom) =
  {
    "name": "Tom",
    "age": 6
  }
*/
```

## 保留零值

某些场景下，需要在输出 JSON 字符串时自动忽略零值，但是在将 JSON 字符串转为目标结构体时，需要保留零值。
可以通过将字段设置为 [指针](pointer.md) 类型即可。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	HasMoney *bool  `json:"hasMoney,omitempty"`
}

func main() {
	tomJson := `
{
  "name": "Tom",
  "age": 6,
  "hasMoney": false
}
`
	var tom person
	err := json.Unmarshal([]byte(tomJson), &tom)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Tom's name is %s, age is %d, hasMoney is %t\n", tom.Name, tom.Age, *tom.HasMoney)
}

// $ go run main.go
// 输出如下 
/**
  Tom's name is Tom, age is 6, hasMoney is false
*/
```

## 忽略公开值

某些场景下，需要将结构体字段设置为公开可导出，但是又不希望 JSON 序列化时输出该字段，
可以使用 `-` 符号标识。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"-"`
}

func main() {
	tom := person{ // 使用字面量初始化
		Name: "Tom",
		Age:  6,
		Hobby: []string{
			"reading",
			"coding",
			"movie",
		},
	}

	// 前缀符为空字符串，缩进符为两个空格
	formatted, err := json.MarshalIndent(tom, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("json.MarshalIndent(tom) = \n%s\n", formatted)
}

// $ go run main.go
// 输出如下 
/**
  json.MarshalIndent(tom) =
  {
    "name": "Tom",
    "age": 6
  }
*/
```