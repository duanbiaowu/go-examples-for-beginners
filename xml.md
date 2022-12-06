# 概述

`encoding/xml` 包含了 XML 相关处理方法。

# 例子

## 结构体转为 XML 字符串

调用 `xml.Marshal()` 方法完成。

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
	addr string `xml:"addr"` // 该属性转 XML 时会被忽略
}

func main() {
	tom := person{ // 使用字面量初始化
		Name: "Tom",
		Age:  6,
		addr: "???",
	}

	tomXml, err := xml.Marshal(tom)
	if err != nil {
		panic(err)
	}
	fmt.Printf("xml.Marshal(tom) = %s\n", tomXml) // 从输出字符串中可以看到，并没有 addr 属性
}

// $ go run main.go
// 输出如下 
/**
  xml.Marshal(tom) = <person><name>Tom</name><age>6</age></person>
*/
```

## XML 字符串转为结构体

调用 `xml.Unmarshal()` 方法完成。

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
	addr string `xml:"addr"` // 该属性转 XML 时会被忽略
}

func main() {
	// 注意: XML 字符串格式一定要正确，否则会报错
	tomXml := `
<person>
    <name>Tom</name>
    <age>6</age>
    <addr>???</addr>
</person>
`
	var tom person
	err := xml.Unmarshal([]byte(tomXml), &tom)
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

## 输出格式化 XML 字符串

调用 `xml.MarshalIndent()` 方法完成。

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name  string   `xml:"name"`
	Age   int      `xml:"age"`
	Hobby []string `xml:"hobby"`
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

	// 前缀符为空字符串，缩进符为 4 个空格
	formatted, err := xml.MarshalIndent(tom, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("xml.MarshalIndent(tom) = \n%s\n", formatted)
}

// $ go run main.go
// 输出如下 
/**
  xml.MarshalIndent(tom) =
  <person>
      <name>Tom</name>
      <age>6</age>
      <hobby>reading</hobby>
      <hobby>coding</hobby>
      <hobby>movie</hobby>
  </person>
*/
```

## 属性值(版本号)

通过 `attr` 关键字完成。

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Version string `xml:"version,attr"` // attr 关键字将字段标记为属性
	Name    string `xml:"name"`
	Age     int    `xml:"age"`
}

func main() {
	tom := person{
		Version: "1.0",
		Name:    "Tom",
		Age:     6,
	}

	formatted, err := xml.MarshalIndent(tom, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", formatted)
}

// $ go run main.go
// 输出如下 
/**
  <person version="1.0">
      <name>Tom</name>
      <age>6</age>
  </person>
*/
```

## 忽略零值

通过 `omitempty` 关键字完成。

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Version string  `xml:"version,attr"`
	Name    string  `xml:"name"`
	Age     int     `xml:"age"`
	Money   float64 `xml:"money,omitempty"` // omitempty 关键字将字段标记为忽略零值
}

func main() {
	tomNoMoney := person{
		Version: "1.0",
		Name:    "Tom",
		Age:     6,
		Money:   0,
	}

	formatted, err := xml.MarshalIndent(tomNoMoney, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n\n", formatted) // 从输出字符串中可以看到，并没有 money 属性

	tomHasMoney := person{
		Version: "1.0",
		Name:    "Tom",
		Age:     6,
		Money:   100,
	}

	formatted, err = xml.MarshalIndent(tomHasMoney, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", formatted) // 从输出字符串中可以看到，有 money 属性
}

// $ go run main.go
// 输出如下 
/**
  <person version="1.0">
      <name>Tom</name>
      <age>6</age>
  </person>

  <person version="1.0">
      <name>Tom</name>
      <age>6</age>
      <money>100</money>
  </person>
*/
```
