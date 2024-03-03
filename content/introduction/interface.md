# 概述

**Go 接口是隐式实现。** 对于一个数据类型，无需声明它实现了哪些接口，只需要实现接口必需的方法即可。
当然了，存在一个小问题就是: 我们可能无意间实现了某个接口:) ，所以 `命名` 是多么重要的一件事情。

# 语法规则

```shell
type 接口名称 interface {
	方法1名称(参数列表 ...) 返回值列表...
	方法2名称(参数列表 ...) 返回值列表...
	方法3名称(参数列表 ...) 返回值列表...
	...
	...
}
```

# 例子

## 求矩形和圆的周长, 面积

```go
package main

import (
	"fmt"
	"math"
)

// 声明一个图形接口
type geometry interface {
	area() float64
	perimeter() float64
}

// 多个字段类型相同时，可以并列声明
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// rectangle 隐式实现了 geometry 接口的 area 方法
func (r *rectangle) area() float64 {
	return r.width * r.height
}

// rectangle 隐式实现了 geometry 接口的 perimeter 方法
func (r *rectangle) perimeter() float64 {
	return (r.width + r.height) * 2
}

// circle 隐式实现了 geometry 接口的 area 方法
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// circle 隐式实现了 geometry 接口的 perimeter 方法
func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := &rectangle{
		width:  10,
		height: 5,
	}
	fmt.Printf("Rectangle area = %.2f, perimeter = %.2f \n", r.area(), r.perimeter())

	c := &circle{
		radius: 10,
	}
	fmt.Printf("Circle area = %.2f, perimeter = %.2f \n", c.area(), c.perimeter())
}

// $ go run main.go
// 输出如下 
/**
  Rectangle area = 50.00, perimeter = 30.00
  Circle area = 314.16, perimeter = 62.83
*/
```

# 编译检测规则

编译器会对接口方法的 `接收者参数` 进行检查，具体来说:

- 接收者形参为普通变量类型
    - 实参为普通变量类型，编译正常
    - 实参为指针变量类型，编译正常
- 接收者形参为指针变量类型
    - **实参为普通变量类型，编译报错**，因为编译器不会无缘无故创建一个新的指针，退一步说，即使编译器可以创建指针指向变量，但是因为 Go 语言传递参数方式为 `值传递`，此时新指针指向的变量，并不是调用该方法的变量，而是经过复制传递的参数变量。 
    - 实参为指针变量类型，编译正常


# reference

1. [Go 圣经](https://book.douban.com/subject/27044219)
