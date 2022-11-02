# 概述

# 语法规则

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
	fmt.Printf("json.Marshal(tom) = %s\n", tomJson)	// 从输出字符串中可以看到，并没有 addr 属性
}
// $ go run main.go
// 输出如下 
/**
    json.Marshal(tom) = {"name":"Tom","age":6} 
*/
```

## JSON 字符串转为结构体
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