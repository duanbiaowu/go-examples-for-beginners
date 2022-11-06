# 概述
Go 支持将多个结构体通过嵌套的方式，组成一个大的结构体，降低了单个结构体复杂度，同时提高了结构体之间组合的灵活性。

# 例子
为了省略篇幅，本小节只使用 `字面量` 方式初始化，`new()` 的初始化方式请参照 [结构体](struct.md) 小节。

## 每个结构体单独初始化，最后组装
```go
package main

import (
	"fmt"
)

type person struct {
	name string
	age  int16
	hobby
	profession
	address
}

type hobby struct {
	values []string
}

type profession struct {
	desc string
}

type address struct {
	tel  string
	area string
}

func main() {
	// 这里使用单个字符作为变量名称，仅仅是为了演示，实际开发中要遵守命名规范
	h := hobby{
		values: []string{"读书", "羽毛球", "电影"},
	}

	p := profession{
		desc: "学生",
	}

	a := address{
		tel:  "123-456789",
		area: "XX 小区 1 栋 2 单元 304",
	}

	li := person{
		name:       "小李",
		age:        12,
		hobby:      h,
		profession: p,
		address:    a,
	}

	fmt.Printf(" 姓名: %s\n 年龄: %d\n 职业: %s\n 爱好: %v\n 电话: %s\n 住址: %s\n",
		li.name, li.age, li.profession.desc, li.hobby.values, li.address.tel, li.area)

	// 其中，嵌套的字段名可以省略
	fmt.Println("\n 省略掉嵌套的字段名，打印结果一样\n")

	fmt.Printf(" 姓名: %s\n 年龄: %d\n 职业: %s\n 爱好: %v\n 电话: %s\n 住址: %s\n",
		li.name, li.age, li.desc, li.values, li.tel, li.area)
}

// $ go run main.go
// 输出如下 
/**
    姓名: 小李
    年龄: 12
    职业: 学生
    爱好: [读书 羽毛球 电影]
    电话: 123-456789
    住址: XX 小区 1 栋 2 单元 304
    
    省略掉嵌套的字段名，打印结果一样
    
    姓名: 小李
    年龄: 12
    职业: 学生
    爱好: [读书 羽毛球 电影]
    电话: 123-456789
    住址: XX 小区 1 栋 2 单元 304
*/
```

## 整个结构体初始化
```go
package main

import (
	"fmt"
)

type person struct {
	name string
	age  int16
	hobby
	profession
	address
}

type hobby struct {
	values []string
}

type profession struct {
	desc string
}

type address struct {
	tel  string
	area string
}

func main() {
	li := person{
		name: "小李",
		age:  12,
		hobby: hobby{
			values: []string{"读书", "羽毛球", "电影"},
		},
		profession: profession{
			desc: "学生",
		},
		address: address{
			tel:  "123-456789",
			area: "XX 小区 1 栋 2 单元 304",
		},
	}

	fmt.Printf(" 姓名: %s\n 年龄: %d\n 职业: %s\n 爱好: %v\n 电话: %s\n 住址: %s\n",
		li.name, li.age, li.desc, li.values, li.tel, li.area)
}
// $ go run main.go
// 输出如下 
/**
    姓名: 小李
    年龄: 12
    职业: 学生
    爱好: [读书 羽毛球 电影]
    电话: 123-456789
    住址: XX 小区 1 栋 2 单元 304
*/
```

## 打印结构体
```go
package main

import (
	"fmt"
)

type person struct {
	name string
	age  int16
	hobby
	profession
	address
}

type hobby struct {
	values []string
}

type profession struct {
	desc string
}

type address struct {
	tel  string
	area string
}

func main() {
	li := person{
		name: "小李",
		age:  12,
		hobby: hobby{
			values: []string{"读书", "羽毛球", "电影"},
		},
		profession: profession{
			desc: "学生",
		},
		address: address{
			tel:  "123-456789",
			area: "XX 小区 1 栋 2 单元 304",
		},
	}

	fmt.Println("默认打印输出如下")
	fmt.Println(li)

	fmt.Println("打印时加上字段名")
	fmt.Printf("%+v\n", li)

	fmt.Println("打印时加上嵌套字段名")
	fmt.Printf("%#v\n", li)
}
// $ go run main.go
// 输出如下 
/**
    默认打印输出如下
    {小李 12 {[读书 羽毛球 电影]} {学生} {123-456789 XX 小区 1 栋 2 单元 304}}
    
    打印时加上字段名
    {name:小李 age:12 hobby:{values:[读书 羽毛球 电影]} profession:{desc:学生} address:{tel:123-456789 area:XX 小区 1 栋 2 单元 304}}
    
    打印时加上嵌套字段名
    main.person{name:"小李", age:12, hobby:main.hobby{values:[]string{"读书", "羽毛球", "电影"}}, profession:main.profession{desc:"学生"}, address:main.address{tel:"123-456789", area:"XX 小区 1 栋 2 单元 304"}}
*/
```

# 扩展阅读
1. [为什么有“组合优于继承”的说法 - 知乎](https://www.zhihu.com/question/21862257)