# 概述
**Go 接口是隐式实现。** 对于一个数据类型，无需声明它实现了哪些接口，只需要实现接口必需的方法即可。
当然了，存在一个小问题就是: 我们可能无意间实现了某个接口:) ，所以 `命名` 是多么重要的一件事情。

# 语法规则
参考 [方法](methods.md) 语法即可。

# 例子

## 求矩形和圆的周长, 面积
```go
package main

import (
	"fmt"
	"math"
)

type geometry interface { // 声明一个图形接口
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64 // 多个字段类型相同时，可以并列声明
}

type circle struct {
	radius float64
}

func (r *rectangle) area() float64 { // rectangle 隐式实现了 geometry 接口的 area 方法
	return r.width * r.height
}

func (r *rectangle) perimeter() float64 { // rectangle 隐式实现了 geometry 接口的 perimeter 方法
	return (r.width + r.height) * 2
}

func (c *circle) area() float64 { // circle 隐式实现了 geometry 接口的 area 方法
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 { // circle 隐式实现了 geometry 接口的 perimeter 方法
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

# reference
1. [Go 圣经](https://book.douban.com/subject/27044219/)
