# 概述

# 语法规则
recover 必须和 defer 配套使用, defer 和 panic 的顺序非常重要。

## 公开/私有
结构性可见性
字段可见性

# 例子

## 声明及初始化
```go
package main

import "fmt"

type person struct {
	name string
	age  int16
}

func main() {
	tom := person{ // 使用字面量初始化
		name: "Tom",
		age:  6,
	}

	fmt.Printf("Tom's name is = %s, age is %d\n", tom.name, tom.age)

	jerry := new(person) // 使用 new 关键字创建
	jerry.name = "Jerry"
	jerry.age = 8

	fmt.Printf("Jerry's name is = %s, age is %d\n", jerry.name, jerry.age)
}

// $ go run main.go
// 输出如下 
/**
    Tom's name is = Tom, age is 6
    Jerry's name is = Jerry, age is 8
*/
```

## 结构体指针
和指针变量一样，在结构体中，也可以通过 `结构体` 指针直接修改结构体字段的值。
```go
package main

import "fmt"

type person struct {
	name string
	age  int16
}

func main() {
	tom := person{
		name: "Tom",
		age:  6,
	}

	fmt.Printf("Tom's name is = %s, age is %d\n", tom.name, tom.age)

	var tomPtr *person
	tomPtr = &tom
	tomPtr.name = "Jerry"
	tomPtr.age = 8

	fmt.Printf("Tom's name is = %s, age is %d\n", tom.name, tom.age)
}
// $ go run main.go
// 输出如下 
/**
    Tom's name is = Tom, age is 6
    Tom's name is = Jerry, age is 8
*/
```