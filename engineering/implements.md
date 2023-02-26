# 概述

Go 中检测一个类型是否实现了某个接口，通常分为两类机制: `编译期间` 和 `运行期间`。

# 编译期间

顾名思义，编译期间检测就是通过静态分析来确定类型是否实现了某个接口，如果检测不通过，则编译过程报错并退出。
通常将这种方式称为 `接口完整性检查`。

## 接口完整性检查

### 类型未实现接口

```go
package main

type Person interface {
	Name() string
	Age() int
}

type Martian struct {
}

func main() {
	// 强制类型 Martian 必须实现接口 Person 的所有方法
	var _ Person = (*Martian)(nil)

	// 1. 声明一个 _ 变量 (不使用)
	// 2. 把一个 nil 转换为 (*Martian)，然后再转换为 Person
	// 3. 如果 Martian 没有实现 Person 的全部方法，则转换失败，编译器报错
}
```

运行代码

```shell
$ go run main.go
# 输出如下 
cannot use (*Martian)(nil) (value of type *Martian) as type Person in variable declaration:
        *Martian does not implement Person (missing Age method)
```

从输出的结果中可以看到，`Martian` 并没有实现 `Person` 接口，所以报错了。下面我们为 `Martian` 实现 `Person` 接口。

### 类型实现了接口

```go
package main

import "fmt"

type Person interface {
	Name() string
	Age() int
}

type Martian struct {
}

// 实现 Name 方法
func (m *Martian) Name() string {
	return "martian"
}

// 实现 Age 方法
func (m *Martian) Age() int {
	return -1
}

func main() {
	// 此时 Martian 已实现了 Person 的全部方法
	var _ Person = (*Martian)(nil)

	m := &Martian{}
	fmt.Printf("name = %s, age = %d\n", m.Name(), m.Age())
}
```

运行代码

```shell
$ go run main.go
# 输出如下 
name = martian, age = -1
```

从输出的结果中可以看到，运行成功，`Martian`  已经实现了 `Person` 接口的全部方法。

# 运行期间

运行期间的检测方式主要有 `类型断言` 和 `反射`。

## 类型断言

### 类型未实现接口

```go
package main

import "fmt"

type Person interface {
	Name() string
	Age() int
}

type Martian struct {
}

func main() {
	// 变量必须声明为 interface 类型
	var m interface{}
	m = &Martian{}
	if v, ok := m.(Person); ok {
		fmt.Printf("name = %s, age = %d\n", v.Name(), v.Age())
	} else {
		fmt.Println("Martian does not implements Person")
	}
}
```

运行代码

```shell
$ go run main.go
# 输出如下 
Martian does not implements Person
```

下面我们为 `Martian` 实现 `Person` 接口。

### 类型实现了接口

```go
package main

import "fmt"

type Person interface {
	Name() string
	Age() int
}

type Martian struct {
}

func (m *Martian) Name() string {
	return "martian"
}

func (m *Martian) Age() int {
	return -1
}

func main() {
	// 变量必须声明为 interface 类型
	var m interface{}
	m = &Martian{}
	if v, ok := m.(Person); ok {
		fmt.Printf("name = %s, age = %d\n", v.Name(), v.Age())
	}
}
```

运行代码

```shell
$ go run main.go
# 输出如下 
name = martian, age = -1
```

从输出的结果中可以看到，运行成功，`Martian`  已经实现了 `Person` 接口的全部方法。

## 反射

通过 `reflect` 包提供的 API 来判断类型是否实现了某个接口。

### 类型未实现接口

```go
package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	Name() string
	Age() int
}

type Martian struct {
}

func main() {
	// 获取 Person 类型
	p := reflect.TypeOf((*Person)(nil)).Elem()

	// 获取 Martian 结构体指针类型
	m := reflect.TypeOf(&Martian{})

	// 判断 Martian 结构体类型是否实现了 Person 接口
	fmt.Println(m.Implements(p))
}
```

运行代码

```shell
$ go run main.go
# 输出如下 
false
```

下面我们为 `Martian` 实现 `Person` 接口。

### 类型实现了接口

```go
package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	Name() string
	Age() int
}

type Martian struct {
}

func (m *Martian) Name() string {
	return "martian"
}

func (m *Martian) Age() int {
	return -1
}

func main() {
	// 获取 Person 类型
	p := reflect.TypeOf((*Person)(nil)).Elem()

	// 获取 Martian 结构体指针类型
	m := reflect.TypeOf(&Martian{})

	// 判断 Martian 结构体类型是否实现了 Person 接口
	fmt.Println(m.Implements(p))
}
```

运行代码

```shell
$ go run main.go
# 输出如下 
true
```

# 小结

Go 的接口实现检测机制分为 `编译期间` 和 `运行期间`，其中编译期间的检测方式是 `接口完整性检查`，
而运行期间的检测方式有两种: `类型断言` 和 `反射`，一般情况尽量使用类型断言，这样可以避免反射带来的性能损耗。
文中提到的几种检测方式语法都比较新 (nan) 奇 (kan) ，读者可以参考代码的注释部分帮助理解。