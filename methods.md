# 概述

# 语法规则

# 使用规则
结构体定义，指针调用
指针定义，结构体调用

# 例子

## 结构体方法
```go
package main

import "fmt"

type person struct {
	name string
	age  int16
}

func (p person) sayName() {
	fmt.Printf("Hi, my name is %s\n", p.name)
}

func (p person) sayAge() {
	fmt.Printf("Hi, my age is %d\n", p.age)
}

func main() {
	tom := &person{
		name: "Tom",
		age:  6,
	}
	tom.sayName()
	tom.sayAge()
}
// $ go run main.go
// 输出如下 
/**
    Hi, my name is Tom
    Hi, my age is 6
*/
```

## 结构体指针方法
相比结构体方法，指针结构体方法除了将方法参数变为指针外，**在引用对应的字段时，无需加 `*` 标识符**，
这一点和普通指针变量引用时有所区别，需要注意。
```go
package main

import "fmt"

type person struct {
	name string
	age  int16
}

func (p *person) sayName() { // 结构体为指针类型
	fmt.Printf("Hi, my name is %s\n", p.name)
}

func (p *person) sayAge() { // 结构体为指针类型
	fmt.Printf("Hi, my age is %d\n", p.age)
}

func main() {
	tom := &person{
		name: "Tom",
		age:  6,
	}
	tom.sayName()
	tom.sayAge()
}
// $ go run main.go
// 输出如下 
/**
    Hi, my name is Tom
    Hi, my age is 6
*/
```