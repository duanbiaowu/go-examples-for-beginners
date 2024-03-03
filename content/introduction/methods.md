# 概述

方法的声明和普通函数的声明类似，只是在函数名字前面多了一个 `接收者参数` (接收者参数将方法绑定到其对应的数据类型上)。
方法可以绑定到任何数据类型上，但是大多数情况下，绑定的都是 [结构体](struct.md)。

# 语法规则

```shell
func (接收者参数) 方法名称(参数列表 ...) 返回值列表... {
    // do something
}
```

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


# reference

1. [Go 圣经](https://book.douban.com/subject/27044219/)
