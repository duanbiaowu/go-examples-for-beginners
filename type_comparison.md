# 概述

`比较运算符` 用来比较两个操作数并返回一个 `bool` 值，常见的比较运算符: 
```shell
==    等于
!=    不等于
<     小于
<=    小于等于
>     大于
>=    大于等于
```

在任何比较中，第一个操作数必须可以赋值给第二个操作数的类型，反过来也一样。

# 不可比较类型

Go 中有 3 中数据类型不能比较，分别是 `slice`, `map`, `func`，如果要比较这 3 种类型，
使用 `reflect.DeepEqual` 函数。

## 为什么 slice 不能比较

**(个人猜测，待验证)**

- 切片是引用类型，比较地址没有意义
- 多个切片引用同一数组时，修改时会相互影响，无法保证 `key` 的一致性
- 切片除了 `len` 属性外，还有 `cap` 属性，比较的维度没办法精确衡量

基于上述原因，官方没有支持 `slice` 类型的比较。

## 为什么 map 不能比较

**(个人猜测，待验证)**

- `map` 遍历是随机的，无法保证 `key` 的一致性 

## 为什么 func 不能比较

**(个人猜测，待验证)**

- 函数的可变参数机制，无法保证 `key` 的一致性
- 函数参数可以为 `slice`, `map`, 这两种类型不可比较

# 可比较类型的具体规则

- `布尔值` 可比较。
- `整型` 可比较。
- `浮点型` 可比较。如果两个浮点型值一样 (由 `IEEE-754` 标准定义)，则两者相等。
- `复数型` 可比较。如果两个复数型值的 `real()` 方法 和 `imag()` 方法都相等，则两者相等。
- `字符串` 可比较。
- `指针` 可比较。如果两个指针指向相同的 `地址` 或者两者都为 `nil`，则两者相等，**但是指向不同的零大小变量的指针可能不相等**。
- `通道` 可比较。如果两个通道是由同一个 `make` 创建的 (引用的是同一个 channel 指针)，或者两者都为 `nil`, 则两者相等。
- `接口` 可比较。`interface` 的内部实现包含了 2 个字段，类型 `T` 和 值 `V`。 如果两个 `接口` 具有相同的动态类型和动态值，或者两者都为 `nil`, 则两者相等。
- `结构体` 可比较 (如果两个结构体的所有字段都是可比较的)。如果两个结构体对应的非空白字段相等，则两者相等。
- `数组` 可比较 (如果两个数组的所有元素都是可比较的)。如果两个数组的所有对应元素相等，则两者相等。

# 例子

## 指针的比较

### 指向相同的地址的指针

```go
package main

import "fmt"

func main() {
	n := 1024
	p := &n
	p2 := &n
	fmt.Printf("p == p2: %t\n", p == p2)
}
// $ go run main.go
// 输出如下
/**
    p == p2: true
*/
```

### 指向 nil 的指针

```go
package main

import "fmt"

func main() {
	var p *string
	var p2 *string
	fmt.Printf("p == p2: %t\n", p == p2)
}
// $ go run main.go
// 输出如下
/**
    p == p2: true
*/
```

## 通道的比较

### 同一个 make() 创建的通道

```go
package main

import "fmt"

func main() {
	ch := make(chan bool)
	ch2 := make(chan bool)

	p := &ch
	p2 := &ch2
	fmt.Printf("p == p2: %t\n", p == p2)

	p3 := &ch
	fmt.Printf("p == p3: %t\n", p == p3)
}
// $ go run main.go
// 输出如下
/**
    p == p2: false
    p == p3: true
*/
```

### 通道为 nil

```go
package main

import "fmt"

func main() {
	var p *chan bool
	var p2 *chan bool

	fmt.Printf("p == p2: %t\n", p == p2)
}
// $ go run main.go
// 输出如下
/**
    p == p2: true
*/
```

## 结构体的比较

**比较的前提: 如果两个结构体的所有字段都是可比较的**。

### 结构体对应的非空白字段相等

```go
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	tom := person{
		name: "Tom",
		age:  6,
	}
	jerry := person{
		name: "Jerry",
		age:  8,
	}
	fmt.Printf("tom == jerry: %t\n", tom == jerry)

	nobody := person{}
	nobody2 := person{}

	fmt.Printf("nobody == nobody2: %t\n", nobody == nobody2)
}
// $ go run main.go
// 输出如下
/**
    tom == jerry: false
    nobody == nobody2: true
*/
```

### 结构体为 nil

```go
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var nobody person
	var nobody2 person

	fmt.Printf("nobody == nobody2: %t\n", nobody == nobody2)
}
// $ go run main.go
// 输出如下
/**
    nobody == nobody2: true
*/
```

## 接口的比较

### 具有相同的动态类型和动态值

```go
package main

import "fmt"

type person struct {
	name string
}

func main() {
	var tom1, tom2 interface{}

	tom1 = &person{"Tom"}
	tom2 = &person{"Tom"}

	var tom3, tom4 interface{}
	tom3 = person{"Tom"}
	tom4 = person{"Tom"}

	fmt.Printf("tom1 == tom2: %t\n", tom1 == tom2) // false
	fmt.Printf("tom3 == tom4: %t\n", tom3 == tom4) // true
}
// $ go run main.go
// 输出如下
/**
    tom1 == tom2: false
    tom3 == tom4: true
*/
```

上面的示例代码中，tom1 和 tom2 对应的类型是 *person，值是 person 结构体的地址，但是两个地址不同，因此两者不相等,
tom3 和 tom4 对应的类型是 person，值是 person 结构体且各字段相等，因此两者相等。

### 接口为 nil

```go
package main

import "fmt"

func main() {
	var tom1, tom2 interface{}
	fmt.Printf("tom1 == tom2: %t\n", tom1 == tom2) // true
}
// $ go run main.go
// 输出如下
/**
    tom1 == tom2: true
*/
```

# 小结

本小节介绍了 Go 的比较运算符以及各种数据类型的比较规则。Go 中大多数数据类型都是可以比较的，
除了 `slice`, `map`, `func` 这 3 种，对于不能比较的原因，笔者给出了一些猜测，感兴趣的读者可以自行验证。
限于时间和篇幅，没有给出所有数据类型的代码示例，读者可以编写代码验证具体的类型比较规则。

# reference
1. https://go.dev/ref/spec#Comparison_operators