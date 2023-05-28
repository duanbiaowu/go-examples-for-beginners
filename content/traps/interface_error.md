# 概述

`interface{}` 类型可以表示任意数据类型，直觉上来看，当然也可以表示 `*interface` 类型。
那么两者之间可以直接转换吗？我们通过两个小例子来验证一下。

# 例子

## interface{} 不能直接转换为 *interface{}

```go
package main

// 参数为 interface
func foo(x interface{}) {
}

// 参数为 *interface
func bar(x *interface{}) {
}

func main() {
	s := "" // s 类型为字符串
	p := &s // p 类型为字符串指针

	foo(s) // ok, interface{} 可以表示字符串
	bar(s) // error, *interface 无法表示字符串

	foo(p) // ok, interface{} 可以表示字符串指针
	bar(p) // error, *interface 无法表示字符串指针
}
```

```shell
$ go run main.go
# 输出如下 
cannot use s (variable of type string) 
  as type *interface{} in argument to bar:
    string does not implement *interface{} 
    (type *interface{} is pointer to interface, not interface)
...
...
```

## *interface{} 可以直接转换为 interface{}

```go
package main

import "fmt"

// 参数为 interface
func foo(x interface{}) {
	fmt.Printf("x is %T\n", x)
}

// 参数为 *interface
func bar(x *interface{}) {
	fmt.Printf("x is %T\n", x)
}

func main() {
	var s interface{} // s 类型为 interface{} 
	s = ""
	p := &s // p 类型为 *interface{}

	foo(p)  // ok, *interface{} 可以直接转换为 interface{}
	bar(p)  // ok, bar 函数的参数类型和 p 类型一样
}
```

```shell
$ go run main.go
# 输出如下
x is *interface {}
x is *interface {}
```

# 小结

`*interface{}` 类型可以直接转换为 `interface{}` 类型，但是 `interface{}` 类型无法直接转换为 `*interface{}` 类型。