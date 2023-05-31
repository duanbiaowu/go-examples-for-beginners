---
date: 2023-01-01
title: Go 陷阱之 map 常见问题
modify: 2023-01-01
---

# 未初始化时导致的报错

当 `map` 声明后但是未初始时，可以获取元素 (虽然获取不到)，但是无法修改元素 (会导致报错)。

## map 未初始化时导致的报错

```go
package main

import "fmt"

func main() {
	var m map[string]int

	// 获取元素不报错
	if v, ok := m["zero"]; ok {
		fmt.Printf("v = %v\n", v)
	} else {
		fmt.Println("element not exist")
	}

	m["zero"] = 0 // 修改元素报错
}

// $ go run main.go
// 输出如下 
/**
  element not exist
  panic: assignment to entry in nil map

  ...
  ...

  exit status 2

*/
```

> 最佳实践: 直接使用 `make` 声明并初始化 `map` 。

# 元素作为指针和非指针时的差异

当 `map` 的元素为结构体指针时，可以修改单个字段，但是当 `map` 的元素为结构体时，无法修改单个字段，**因为 `map` 中的元素 (为结构体时) 是不可寻址的。**

## 元素为结构体指针时，可以修改单个字段

```go
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*person)
	m["tom"] = &person{
		name: "Tom",
		age:  6,
	}

	m["tom"].age = 60

	fmt.Printf("%+v\n", *(m["tom"]))
}

// $ go run main.go
// 输出如下 
/**
  {name:Tom age:60}
*/
```

## map 元素为结构体时，无法修改单个字段

### 错误的做法

```go
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	m := make(map[string]person)
	m["tom"] = person{
		name: "Tom",
		age:  6,
	}

	m["tom"].age = 60

	fmt.Printf("%+v\n", m)
}

// $ go run main.go
// 输出如下
/**
  ./main.go:17:2: cannot assign to struct field m["tom"].age in map
*/
```

### 正确的做法

通过将 `map` 元素赋值给一个临时变量，在临时变量上面做修改，然后再对 `map` 元素重新赋值，就可以修改 `map` 元素了。

```go
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	m := make(map[string]person)
	m["tom"] = person{
		name: "Tom",
		age:  6,
	}

	tom := m["tom"]
	tom.age = 60
	m["tom"] = tom

	fmt.Printf("%+v\n", m["tom"])
}

// $ go run main.go
// 输出如下 
/**
  {name:Tom age:60}
*/
```