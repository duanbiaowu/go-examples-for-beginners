# 概述

**Go 里面没有 `implements` 关键字来判断一个结构体 (对象) 是否实现了某个接口，** 要实现类似的功能，
需要用到 `类型断言`。

## 类型断言

```go
// 判断 v 是否实现了接口 i
// 返回值 ok 为 true 表示已实现，ok 为 false 表示未实现
if _, ok := v.(i); ok {
// do something...
} 
```

# 例子

## 判断是否实现接口

```go
package main

import "fmt"

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
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

func main() {
	var r interface{}

	r = &rectangle{
		width:  10,
		height: 5,
	}
	if v, ok := r.(geometry); ok {
		fmt.Printf("r implements interface geometry, area = %.2f, perimeter = %.2f \n", v.area(), v.perimeter())
	}

	var c interface{}
	c = &circle{
		radius: 10,
	}
	if _, ok := c.(geometry); !ok {
		fmt.Println("c does not implement interface geometry")
	}
}

// $ go run main.go
// 输出如下 
/**
  r implements interface geometry, area = 50.00, perimeter = 30.00
  c does not implement interface geometry
*/
```