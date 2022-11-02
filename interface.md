# 概述

# 语法规则

# 使用规则

# 例子

## 求矩形和圆的周长, 面积
```go
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64 // 多个字段类型相同时，可以并列声明
}

type circle struct {
	radius float64
}

func (r *rectangle) area() float64 {
	return r.width * r.height
}

func (r *rectangle) perimeter() float64 {
	return (r.width + r.height) * 2
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

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
